package get_shop_detail

import "context"

type service interface {
	GetShopDetail(context.Context, Query) (*Result, error)
}
type handler struct{ service service }

func NewHandler(service service) Executor { return &handler{service: service} }
func (h *handler) Execute(ctx context.Context, query Query) (*Result, error) {
	return h.service.GetShopDetail(ctx, query)
}
