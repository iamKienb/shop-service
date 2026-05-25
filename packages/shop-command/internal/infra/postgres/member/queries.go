package member

import (
	"context"
	"errors"
	"fmt"
	domain_member "user-command-module/internal/domain/member"
	"user-command-module/internal/domain/shared"

	"github.com/jackc/pgx/v5"
)

func (r *memberRepository) GetUserRoles(ctx context.Context, shopID shared.ShopID, userID shared.UserID) (*domain_member.MemberPermission, error) {
	row, err := r.getQuerier(ctx).GetUserRolesInShop(ctx, toInfraGetUserRole(shopID, userID))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("infra: get user roles in shop: %w", err)
	}

	return toDomainMemberPermission(row), nil
}
