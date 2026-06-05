package search_shops

import (
	"context"

	"shop-query-module/internal/application/service/models"
)

type Query struct {
	Keyword string
	Status  string
	Page    models.Page
}
type Result = models.ShopPage

type Executor interface {
	Execute(context.Context, Query) (*Result, error)
}
