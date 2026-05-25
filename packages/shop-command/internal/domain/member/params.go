package member

import "user-command-module/internal/domain/shared"

type NewMemberParams struct {
	ShopID shared.ShopID

	MemberID   shared.UserID
	MemberName string

	AddedBy     shared.UserID
	NameAddedBy string
	RoleIDs     []shared.RoleID
}
