package module

import (
	"shop-query-module/internal/application/queries/get_shop_detail"
	"shop-query-module/internal/application/queries/list_shop_addresses"
	"shop-query-module/internal/application/queries/list_shop_members"
	"shop-query-module/internal/application/queries/search_shops"
	"shop-query-module/internal/application/service"
)

type ApplicationModule struct {
	GetShopDetailExecutor     get_shop_detail.Executor
	SearchShopsExecutor       search_shops.Executor
	ListShopAddressesExecutor list_shop_addresses.Executor
	ListShopMembersExecutor   list_shop_members.Executor
}

func NewApplicationModule(infra *InfraModule) *ApplicationModule {
	shopQueryService := service.NewQueryService(infra.ESService)
	return &ApplicationModule{
		GetShopDetailExecutor:     get_shop_detail.NewHandler(shopQueryService),
		SearchShopsExecutor:       search_shops.NewHandler(shopQueryService),
		ListShopAddressesExecutor: list_shop_addresses.NewHandler(shopQueryService),
		ListShopMembersExecutor:   list_shop_members.NewHandler(shopQueryService),
	}
}
