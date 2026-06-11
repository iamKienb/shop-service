package list_shop_members

import "context"

type service interface {
	ListShopMembers(context.Context, Query) (*Result, error)
}
type handler struct {
	service service
}

func NewHandler(service service) Executor {
	return &handler{service: service}
}

func (h *handler) Execute(ctx context.Context, query Query) (*Result, error) {
	return h.service.ListShopMembers(ctx, query)
}
