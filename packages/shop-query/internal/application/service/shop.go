package service

import (
	"context"
	"strconv"

	"shop-query-module/internal/application/queries/get_shop_detail"
	"shop-query-module/internal/application/queries/list_shop_addresses"
	"shop-query-module/internal/application/queries/list_shop_members"
	"shop-query-module/internal/application/queries/search_shops"
	"shop-query-module/internal/application/service/models"
	"shop-shared-module/alias"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/iamKienb/go-core/app_error"
	esx "github.com/iamKienb/go-core/elasticsearch"
)

type QueryService interface {
	GetShopDetail(context.Context, get_shop_detail.Query) (*get_shop_detail.Result, error)
	SearchShops(context.Context, search_shops.Query) (*search_shops.Result, error)
	ListShopAddresses(context.Context, list_shop_addresses.Query) (*list_shop_addresses.Result, error)
	ListShopMembers(context.Context, list_shop_members.Query) (*list_shop_members.Result, error)
}

type queryService struct {
	esClient *elasticsearch.TypedClient
	index    string
}

const (
	defaultPageSize   = 20
	maxPageSize       = 100
	sortAsc           = "asc"
	sortDesc          = "desc"
	errMsgShopMissing = "shop was not found"
)

func NewQueryService(esService esx.ESXService) QueryService {
	return &queryService{
		esClient: esService.GetClient(),
		index:    alias.ShopAlias,
	}
}

func (s *queryService) GetShopDetail(ctx context.Context, query get_shop_detail.Query) (*get_shop_detail.Result, error) {
	shop, err := s.findShop(ctx, query.ShopID)
	if err != nil {
		return nil, err
	}

	return &get_shop_detail.Result{Shop: shop}, nil
}

func (s *queryService) SearchShops(ctx context.Context, query search_shops.Query) (*search_shops.Result, error) {
	page := normalizePage(query.Page)
	builder := NewQueryBuilder().
		WithPagination(pageOffset(page.Token), page.Size).
		MustMultiMatch(query.Keyword, []string{"name", "slug", "profile.description"}).
		WithSort("created_at", sortDesc).
		WithSort("id.keyword", sortAsc)

	if query.Status != "" {
		builder.FilterTerm("status.keyword", query.Status)
	}

	result, err := SearchDocuments[models.Shop](ctx, s.esClient, s.index, builder.Build())
	if err != nil {
		return nil, err
	}

	items := shopsFromHits(result.Hits)
	return &search_shops.Result{
		Items:         items,
		Total:         result.Total,
		NextPageToken: nextPageToken(page, len(items), result.Total),
	}, nil
}

func (s *queryService) ListShopAddresses(ctx context.Context, query list_shop_addresses.Query) (*list_shop_addresses.Result, error) {
	shop, err := s.findShop(ctx, query.ShopID)
	if err != nil {
		return nil, err
	}
	return &list_shop_addresses.Result{Addresses: shop.Addresses}, nil
}

func (s *queryService) ListShopMembers(ctx context.Context, query list_shop_members.Query) (*list_shop_members.Result, error) {
	shop, err := s.findShop(ctx, query.ShopID)
	if err != nil {
		return nil, err
	}
	return &list_shop_members.Result{Members: shop.Members}, nil
}

func (s *queryService) findShop(ctx context.Context, shopID string) (*models.Shop, error) {
	searchQuery := NewQueryBuilder().
		WithPagination(0, 1).
		FilterTerm("id", shopID).
		Build()

	result, err := SearchDocuments[models.Shop](ctx, s.esClient, s.index, searchQuery)
	if err != nil {
		return nil, err
	}
	if len(result.Hits) == 0 {
		return nil, app_error.NotFound(errMsgShopMissing, nil)
	}
	return &result.Hits[0].Source, nil
}

func shopsFromHits(hits []SearchHit[models.Shop]) []models.Shop {
	items := make([]models.Shop, 0, len(hits))
	for _, hit := range hits {
		items = append(items, hit.Source)
	}
	return items
}

func normalizePage(page models.Page) models.Page {
	if page.Size <= 0 || page.Size > maxPageSize {
		page.Size = defaultPageSize
	}
	return page
}

func pageOffset(token string) int {
	if token == "" {
		return 0
	}
	offset, err := strconv.Atoi(token)
	if err != nil || offset < 0 {
		return 0
	}
	return offset
}

func nextPageToken(page models.Page, resultCount int, total int64) string {
	nextOffset := pageOffset(page.Token) + resultCount
	if int64(nextOffset) >= total || resultCount == 0 {
		return ""
	}
	return strconv.Itoa(nextOffset)
}
