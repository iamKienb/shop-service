package member

import (
	"user-command-module/internal/domain/shared"
)

type MemberRole struct {
	ShopID    shared.ShopID
	MemberID  shared.UserID
	RoleID    shared.RoleID
	UpdatedBy shared.UserID
}

type MemberPermission struct {
	ShopID  shared.ShopID
	RoleIDs []shared.RoleID
}
