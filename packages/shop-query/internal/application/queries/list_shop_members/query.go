package list_shop_members

import (
	"context"

	"shop-query-module/internal/application/service/models"
)

type Query struct{ ShopID string }
type Result struct{ Members []models.ShopMember }
type Executor interface {
	Execute(context.Context, Query) (*Result, error)
}
