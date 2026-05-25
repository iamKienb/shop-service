package member

import (
	"time"
	"user-command-module/internal/domain/shared"
)

type RoleEvent struct {
	ID   shared.RoleID
	Code string
	Name string
}

type MemberAddedEvent struct {
	ShopID shared.ShopID

	MemberID   shared.UserID
	MemberName string

	AddedBy     shared.UserID
	NameAddedBy string
	Roles       []RoleEvent
	JoinedAt    time.Time
}

func (e MemberAddedEvent) EventName() string {
	return "user-service.shop.member_added"
}

func (e MemberAddedEvent) IntegrationPayload() map[string]interface{} {
	rolesMap := make([]map[string]interface{}, 0, len(e.Roles))
	for _, r := range e.Roles {
		rolesMap = append(rolesMap, map[string]interface{}{
			"id":   int(r.ID),
			"code": r.Code,
			"name": r.Name,
		})
	}

	return map[string]interface{}{
		"shop_id":       e.ShopID.String(),
		"member_id":     e.MemberID.String(),
		"member_name":   e.MemberName,
		"added_by":      e.AddedBy.String(),
		"name_added_by": e.NameAddedBy,
		"joined_at":     e.JoinedAt,
		"roles":         rolesMap,
	}
}
