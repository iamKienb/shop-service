package shop

import (
	"context"

	"shop-query-module/internal/application/queries/get_shop_detail"
	"shop-query-module/internal/application/queries/list_shop_addresses"
	"shop-query-module/internal/application/queries/list_shop_members"
	"shop-query-module/internal/application/queries/search_shops"

	esx "github.com/iamKienb/go-core/elasticsearch"
)

type QueryService interface {
	GetShopDetail(context.Context, get_shop_detail.Query) (*get_shop_detail.Result, error)
	SearchShops(context.Context, search_shops.Query) (*search_shops.Result, error)
	ListShopAddresses(context.Context, list_shop_addresses.Query) (*list_shop_addresses.Result, error)
	ListShopMembers(context.Context, list_shop_members.Query) (*list_shop_members.Result, error)
}

type queryService struct{ search *SearchService }

func NewQueryService(esService esx.ESXService) QueryService {
	return &queryService{search: NewSearchService(esService)}
}
