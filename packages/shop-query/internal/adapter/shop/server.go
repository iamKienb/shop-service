package shop

import (
	"context"

	"shop-query-module/internal/application/queries/get_shop_detail"
	"shop-query-module/internal/application/queries/list_shop_addresses"
	"shop-query-module/internal/application/queries/list_shop_members"
	"shop-query-module/internal/application/queries/search_shops"
	"shop-query-module/internal/application/service/models"

	"connectrpc.com/connect"
	api "github.com/iamKienb/api-contract/gen/shop"
	"github.com/iamKienb/api-contract/gen/shop/shopconnect"
)

type queryServer struct {
	getShopDetailExecutor     get_shop_detail.Executor
	searchShopsExecutor       search_shops.Executor
	listShopAddressesExecutor list_shop_addresses.Executor
	listShopMembersExecutor   list_shop_members.Executor
}

func NewQueryServer(
	getShopDetailExecutor get_shop_detail.Executor,
	searchShopsExecutor search_shops.Executor,
	listShopAddressesExecutor list_shop_addresses.Executor,
	listShopMembersExecutor list_shop_members.Executor,
) *queryServer {
	return &queryServer{getShopDetailExecutor: getShopDetailExecutor, searchShopsExecutor: searchShopsExecutor, listShopAddressesExecutor: listShopAddressesExecutor, listShopMembersExecutor: listShopMembersExecutor}
}

func (s *queryServer) GetShopDetail(ctx context.Context, req *connect.Request[api.GetShopDetailRequest]) (*connect.Response[api.GetShopDetailResponse], error) {
	result, err := s.getShopDetailExecutor.Execute(ctx, get_shop_detail.Query{ShopID: req.Msg.GetShopId()})
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&api.GetShopDetailResponse{Shop: ToShopView(result.Shop)}), nil
}

func (s *queryServer) SearchShops(ctx context.Context, req *connect.Request[api.SearchShopsRequest]) (*connect.Response[api.SearchShopsResponse], error) {
	result, err := s.searchShopsExecutor.Execute(ctx, search_shops.Query{
		Keyword: req.Msg.GetKeyword(),
		Status:  req.Msg.GetStatus(),
		Page:    models.Page{Size: int(req.Msg.GetPageSize()), Token: req.Msg.GetPageToken()},
	})
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&api.SearchShopsResponse{Shops: ToShopViews(result.Items), Total: result.Total, NextPageToken: result.NextPageToken}), nil
}

func (s *queryServer) ListShopAddresses(ctx context.Context, req *connect.Request[api.ListShopAddressesRequest]) (*connect.Response[api.ListShopAddressesResponse], error) {
	result, err := s.listShopAddressesExecutor.Execute(ctx, list_shop_addresses.Query{ShopID: req.Msg.GetShopId()})
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&api.ListShopAddressesResponse{Addresses: ToShopAddressViews(result.Addresses)}), nil
}

func (s *queryServer) ListShopMembers(ctx context.Context, req *connect.Request[api.ListShopMembersRequest]) (*connect.Response[api.ListShopMembersResponse], error) {
	result, err := s.listShopMembersExecutor.Execute(ctx, list_shop_members.Query{ShopID: req.Msg.GetShopId()})
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&api.ListShopMembersResponse{Members: ToShopMemberViews(result.Members)}), nil
}

var _ shopconnect.ShopQueryServiceHandler = (*queryServer)(nil)
