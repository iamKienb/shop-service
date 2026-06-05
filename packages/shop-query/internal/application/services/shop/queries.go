package shop

import (
	"context"

	"shop-query-module/internal/application/port"
	"shop-query-module/internal/application/queries/get_shop_detail"
	"shop-query-module/internal/application/queries/list_shop_addresses"
	"shop-query-module/internal/application/queries/list_shop_members"
	"shop-query-module/internal/application/queries/search_shops"
	"shop-shared-module/alias"

	"github.com/iamKienb/go-core/app_error"
)

const errMsgShopNotFound = "shop was not found"

func (s *queryService) GetShopDetail(ctx context.Context, query get_shop_detail.Query) (*get_shop_detail.Result, error) {
	shop, err := s.findShop(ctx, query.ShopID)
	if err != nil {
		return nil, err
	}
	return &get_shop_detail.Result{Shop: shop}, nil
}

func (s *queryService) SearchShops(ctx context.Context, query search_shops.Query) (*search_shops.Result, error) {
	filters := make([]Filter, 0, 1)
	if query.Status != "" {
		filters = append(filters, Term("status.keyword", query.Status))
	}
	must := make([]Filter, 0, 1)
	if query.Keyword != "" {
		must = append(must, MultiMatch(query.Keyword, []string{"name", "slug", "profile.description"}))
	}
	result, err := SearchDocuments[port.Shop](ctx, s.search, SearchSpec{
		Index:   alias.ShopAlias,
		Page:    normalizePage(query.Page),
		Filters: filters,
		Must:    must,
		Sorts: []Sort{
			{Field: "created_at", Direction: SortDesc},
			{Field: "id.keyword", Direction: SortAsc},
		},
	})
	if err != nil {
		return nil, err
	}
	return &search_shops.Result{Items: result.Items, Total: result.Total, NextPageToken: result.NextPageToken}, nil
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

func (s *queryService) findShop(ctx context.Context, shopID string) (*port.Shop, error) {
	result, err := SearchDocuments[port.Shop](ctx, s.search, SearchSpec{
		Index:   alias.ShopAlias,
		Page:    port.Page{Size: 1},
		Filters: []Filter{Term("id.keyword", shopID)},
	})
	if err != nil {
		return nil, err
	}
	if len(result.Items) == 0 {
		return nil, app_error.NotFound(errMsgShopNotFound, nil)
	}
	return &result.Items[0], nil
}

func normalizePage(page port.Page) port.Page {
	if page.Size <= 0 || page.Size > 100 {
		page.Size = 20
	}
	return page
}
