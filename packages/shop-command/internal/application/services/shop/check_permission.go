package shop

import (
	"context"
	"user-command-module/internal/application/queries/check_permission"
)

func (s *shopService) CheckPermission(ctx context.Context, qry check_permission.Query) (*check_permission.Result, error) {
	userRoleIDs, err := s.getUserRoles(ctx, qry.ShopID, qry.UserID)
	if err != nil {
		return nil, s.wrapError(err)
	}

	if err := s.memAuthor.Authorize(qry.Action, userRoleIDs); err != nil {
		return &check_permission.Result{
			IsAllowed: false,
			Message:   err.Error(),
		}, nil
	}

	return &check_permission.Result{
		IsAllowed: true,
		Message:   "",
	}, nil
}
