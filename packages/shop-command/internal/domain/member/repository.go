package member

import (
	"context"
	"shop-command-module/internal/domain/shared"
)

type QueryRepository interface {
	GetUserRoles(ctx context.Context, shopID shared.ShopID, userID shared.UserID) (*MemberPermission, error)
}

type CommandRepository interface {
	SaveMembers(ctx context.Context, members []*Member) error
	ClearMemberRolesBatch(ctx context.Context, memberRoles []*MemberRole) error
}

type Repository interface {
	QueryRepository
	CommandRepository
}
