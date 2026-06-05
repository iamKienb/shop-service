package member

import (
	"shop-command-module/internal/domain/shared"
	"time"
)

const (
	RoleOwnerID     shared.RoleID = 1
	RoleManagerID   shared.RoleID = 2
	RoleCashierID   shared.RoleID = 3
	RoleWarehouseID shared.RoleID = 4
	RoleMarketingID shared.RoleID = 5
)

var RoleMetadata = map[shared.RoleID]struct {
	Code string
	Name string
}{
	RoleOwnerID:     {Code: "OWNER", Name: "Store Owner"},
	RoleManagerID:   {Code: "MANAGER", Name: "Store Manager"},
	RoleCashierID:   {Code: "CASHIER", Name: "Cashier"},
	RoleWarehouseID: {Code: "WAREHOUSE", Name: "Warehouse Staff"},
	RoleMarketingID: {Code: "MARKETING", Name: "Marketing Staff"},
}

type Member struct {
	ShopID   shared.ShopID
	MemberID shared.UserID
	JoinedAt time.Time
	AddedBy  shared.UserID

	MemberRoles []MemberRole
	shared.EventEntity
}

func NewMember(params NewMemberParams) *Member {
	now := time.Now().UTC()

	member := &Member{
		ShopID:   params.ShopID,
		MemberID: params.MemberID,
		JoinedAt: now,
		AddedBy:  params.AddedBy,
	}

	var memberRoles []MemberRole
	var eventRoles []RoleEvent

	for _, roleID := range params.RoleIDs {
		memberRoles = append(memberRoles, MemberRole{
			ShopID:    params.ShopID,
			MemberID:  params.MemberID,
			RoleID:    roleID,
			UpdatedBy: params.AddedBy,
		})

		if meta, exists := RoleMetadata[roleID]; exists {
			eventRoles = append(eventRoles, RoleEvent{
				ID:   roleID,
				Code: meta.Code,
				Name: meta.Name,
			})
		}
	}

	member.MemberRoles = memberRoles

	member.AddEvent(MemberAddedEvent{
		ShopID:     params.ShopID,
		MemberID:   params.MemberID,
		MemberName: params.MemberName,

		AddedBy:     params.AddedBy,
		NameAddedBy: params.NameAddedBy,

		JoinedAt: now,
		Roles:    eventRoles,
	})

	return member
}

func (m *Member) FlushEvents() []shared.DomainEvent {
	var domainEvents []shared.DomainEvent
	domainEvents = append(domainEvents, m.Flush()...)
	m.ClearEvent()

	return domainEvents
}

func (m Member) Type() string {
	return "MEMBER"
}
