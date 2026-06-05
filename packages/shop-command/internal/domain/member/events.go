package member

import (
	"shop-command-module/internal/domain/shared"
	"time"
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
	return "shop-service.shop.member.added"
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

// type MemberRolesAssignedEvent struct {
// 	ShopID shared.ShopID

// 	MemberID   shared.UserID
// 	MemberName string

// 	UpdatedBy     shared.UserID
// 	NameUpdatedBy string
// 	Roles         []RoleEvent
// 	UpdatedAt     time.Time
// }

// func (e MemberRolesAssignedEvent) EventName() string {
// 	return "shop-service.shop.member.roles_assigned"
// }

// func (e MemberRolesAssignedEvent) IntegrationPayload() map[string]interface{} {
// 	rolesMap := make([]map[string]interface{}, 0, len(e.Roles))
// 	for _, r := range e.Roles {
// 		rolesMap = append(rolesMap, map[string]interface{}{
// 			"id":   int(r.ID),
// 			"code": r.Code,
// 			"name": r.Name,
// 		})
// 	}

// 	return map[string]interface{}{
// 		"shop_id":         e.ShopID.String(),
// 		"member_id":       e.MemberID.String(),
// 		"member_name":     e.MemberName,
// 		"updated_by":      e.UpdatedBy.String(),
// 		"name_updated_by": e.NameUpdatedBy,
// 		"updated_at":      e.UpdatedAt,
// 		"roles":           rolesMap,
// 	}
// }

// func ToRoleEvents(roleIDs []shared.RoleID) []RoleEvent {
// 	events := make([]RoleEvent, 0, len(roleIDs))
// 	for _, roleID := range roleIDs {
// 		if meta, exists := RoleMetadata[roleID]; exists {
// 			events = append(events, RoleEvent{
// 				ID:   roleID,
// 				Code: meta.Code,
// 				Name: meta.Name,
// 			})
// 		}
// 	}

// 	return events
// }
