package check_permission

import "context"

type service interface {
	CheckPermission(ctx context.Context, qry Query) (*Result, error)
}

type handler struct {
	service service
}

func NewHandler(service service) Executor {
	return &handler{service: service}
}

func (h *handler) Execute(ctx context.Context, qry Query) (*Result, error) {
	return h.service.CheckPermission(ctx, qry)
}
