package search_shops

import (
	"context"

	"shop-query-module/internal/application/port"
)

type Query struct {
	Keyword string
	Status  string
	Page    port.Page
}
type Result = port.ShopPage

type Executor interface {
	Execute(context.Context, Query) (*Result, error)
}
