package shop

import (
	"context"
	"errors"
	"shop-command-module/internal/application/queries/check_permission"
	"shop-command-module/internal/application/services/shop/i18n"
	"shop-command-module/internal/domain/member"
	"shop-command-module/internal/domain/shop"
)

func (s *shopService) CheckPermission(ctx context.Context, qry check_permission.Query) (*check_permission.Result, error) {
	if err := s.authorize(ctx, qry.ShopID, qry.UserID, qry.Action); err != nil {
		return &check_permission.Result{
			IsAllowed: false,
			Message:   permissionDeniedMessage(err),
		}, nil
	}

	return &check_permission.Result{
		IsAllowed: true,
		Message:   "",
	}, nil
}

func permissionDeniedMessage(err error) string {
	switch {
	case errors.Is(err, shop.ErrShopNotFound):
		return i18n.MsgShopNotFound
	case errors.Is(err, member.ErrActionNotDefined):
		return i18n.MsgActionInvalid
	case errors.Is(err, member.ErrShopDenied):
		return i18n.MsgShopDenied
	case errors.Is(err, member.ErrProductDenied):
		return i18n.MsgProductDenied
	case errors.Is(err, member.ErrInventoryDenied):
		return i18n.MsgInventoryDenied
	default:
		return i18n.MsgShopDenied
	}
}
