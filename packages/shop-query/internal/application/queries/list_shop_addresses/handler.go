package list_shop_addresses

import "context"

type service interface {
	ListShopAddresses(context.Context, Query) (*Result, error)
}
type handler struct{ service service }

func NewHandler(service service) Executor { return &handler{service: service} }
func (h *handler) Execute(ctx context.Context, query Query) (*Result, error) {
	return h.service.ListShopAddresses(ctx, query)
}
