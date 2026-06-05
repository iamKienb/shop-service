package shop

import (
	"context"
	"shop-command-module/internal/application/queries/check_permission"

	"github.com/iamKienb/go-core/app_error"
)

func (s *shopService) CheckPermission(ctx context.Context, qry check_permission.Query) (*check_permission.Result, error) {
	if err := s.authorize(ctx, qry.ShopID, qry.UserID, qry.Action); err != nil {
		mappedErr := app_error.From(s.wrapError(err))
		return &check_permission.Result{
			IsAllowed: false,
			Message:   mappedErr.PublicMessage(),
		}, nil
	}

	return &check_permission.Result{
		IsAllowed: true,
		Message:   "",
	}, nil
}
