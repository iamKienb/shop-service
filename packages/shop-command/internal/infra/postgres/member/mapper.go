package member

import (
	"user-command-module/db/repository"
	domain_member "user-command-module/internal/domain/member"
	"user-command-module/internal/domain/shared"

	"github.com/iamKienb/go-core/postgres/conv"
	"github.com/jackc/pgx/v5/pgtype"
)

func toInfraClearMemberRoles(memberRoles []*domain_member.MemberRole) repository.ClearShopMemberRolesBatchParams {
	if len(memberRoles) == 0 {
		return repository.ClearShopMemberRolesBatchParams{}
	}

	memberIDs := make([]pgtype.UUID, 0, len(memberRoles))

	for _, memberRole := range memberRoles {
		if memberRole == nil {
			continue
		}

		memberIDs = append(memberIDs, conv.UUID(memberRole.MemberID))
	}

	return repository.ClearShopMemberRolesBatchParams{
		ShopID:    conv.UUID(memberRoles[0].ShopID),
		MemberIds: memberIDs,
	}
}

func toDomainMemberPermission(rows []repository.GetUserRolesInShopRow) *domain_member.MemberPermission {
	if len(rows) == 0 {
		return nil
	}

	roleIds := make([]shared.RoleID, 0, len(rows))

	for _, row := range rows {
		if !row.RoleID.Valid {
			continue
		}
		roleIds = append(roleIds, shared.RoleID(row.RoleID.Int32))
	}

	return &domain_member.MemberPermission{
		ShopID:  rows[0].ShopID.Bytes,
		RoleIDs: roleIds,
	}
}

func toInfraGetUserRole(shopID shared.ShopID, userID shared.UserID) repository.GetUserRolesInShopParams {
	return repository.GetUserRolesInShopParams{
		ID:       conv.UUID(shopID),
		MemberID: conv.UUID(userID),
	}
}

func toInfraMemberAndRoleBatch(members []*domain_member.Member) (repository.AddShopMembersBatchParams, repository.AssignShopMemberRolesBatchParams) {
	if len(members) == 0 {
		return repository.AddShopMembersBatchParams{}, repository.AssignShopMemberRolesBatchParams{}
	}

	firstMember := members[0]
	shopID := conv.UUID(firstMember.ShopID)
	addedBy := conv.UUID(firstMember.AddedBy)
	joinedAt := conv.TimeStampZ(&firstMember.JoinedAt)

	var memberIDs []pgtype.UUID
	var roleMemberIDs []pgtype.UUID
	var roleIDs []int32

	for _, member := range members {
		if member == nil {
			continue
		}

		memberID := conv.UUID(member.MemberID)
		memberIDs = append(memberIDs, memberID)

		for _, role := range member.MemberRoles {
			roleMemberIDs = append(roleMemberIDs, memberID)
			roleIDs = append(roleIDs, int32(role.RoleID))
		}
	}

	membersParams := repository.AddShopMembersBatchParams{
		ShopID:    shopID,
		AddedBy:   addedBy,
		MemberIds: memberIDs,
		JoinedAt:  joinedAt,
	}

	memberRolesParams := repository.AssignShopMemberRolesBatchParams{
		ShopID:    shopID,
		MemberIds: roleMemberIDs,
		RoleIds:   roleIDs,
		UpdatedBy: addedBy,
	}

	return membersParams, memberRolesParams
}
