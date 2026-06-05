package shop

import (
	"context"

	domain_shared "shop-command-module/internal/domain/shared"
	"shop-command-module/internal/domain/shop"
)

const (
	shopAggregateType = "SHOP"
)

func (s *shopService) authorize(ctx context.Context, shopID domain_shared.ShopID, userID domain_shared.UserID, action string) error {
	memberPermission, err := s.memberRepo.GetUserRoles(ctx, shopID, userID)
	if err != nil {
		return err
	}

	if memberPermission == nil {
		return shop.ErrShopNotFound
	}

	return s.memAuthor.Authorize(action, memberPermission.RoleIDs)
}
