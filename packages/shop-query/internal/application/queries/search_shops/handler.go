package search_shops

import "context"

type service interface {
	SearchShops(context.Context, Query) (*Result, error)
}
type handler struct{ service service }

func NewHandler(service service) Executor { return &handler{service: service} }

func (h *handler) Execute(ctx context.Context, query Query) (*Result, error) {
	return h.service.SearchShops(ctx, query)
}
