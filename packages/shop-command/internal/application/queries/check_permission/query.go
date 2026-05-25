package check_permission

import (
	"context"
	"user-command-module/internal/domain/shared"
)

type Query struct {
	ShopID shared.ShopID
	UserID shared.UserID
	Action string
}

type Result struct {
	IsAllowed bool
	Message   string
}

type Executor interface {
	Execute(ctx context.Context, qry Query) (*Result, error)
}
