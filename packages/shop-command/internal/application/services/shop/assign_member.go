package shop

import (
	"context"

	"shop-command-module/internal/application/commands/assign_member"
	"shop-command-module/internal/application/port"
	"shop-command-module/internal/domain/member"
)

func (s *shopService) AssignMember(ctx context.Context, cmd assign_member.Command) (*assign_member.Result, error) {
	if err := s.authorize(ctx, cmd.ShopID, cmd.User.ID, cmd.Action); err != nil {
		return nil, err
	}

	newMembers := make([]*member.Member, 0, len(cmd.MemberRoles))
	memberRolesToClear := make([]*member.MemberRole, 0, len(cmd.MemberRoles))
	outboxParams := make([]port.OutboxParam, 0, len(cmd.MemberRoles))

	for _, memberRole := range cmd.MemberRoles {
		mb := member.NewMember(member.NewMemberParams{
			ShopID:      cmd.ShopID,
			MemberID:    memberRole.MemberID,
			MemberName:  memberRole.Name,
			AddedBy:     cmd.User.ID,
			NameAddedBy: cmd.User.Name,
			RoleIDs:     memberRole.RoleIDs,
		})
		newMembers = append(newMembers, mb)

		if events := mb.FlushEvents(); len(events) > 0 {
			outboxParams = append(outboxParams, port.OutboxParam{
				AggregateID:   mb.MemberID.RawID(),
				AggregateType: mb.Type(),
				Events:        events,
			})
		}

		for _, roleID := range memberRole.RoleIDs {
			memberRolesToClear = append(memberRolesToClear, &member.MemberRole{
				ShopID:    cmd.ShopID,
				MemberID:  memberRole.MemberID,
				RoleID:    roleID,
				UpdatedBy: cmd.User.ID,
			})
		}
	}

	if err := s.txManager.WithTx(ctx, func(ctx context.Context) error {
		if err := s.memberRepo.ClearMemberRolesBatch(ctx, memberRolesToClear); err != nil {
			return err
		}

		if err := s.memberRepo.SaveMembers(ctx, newMembers); err != nil {
			return err
		}

		if len(outboxParams) > 0 {
			return s.outboxService.PublishBatch(ctx, outboxParams)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &assign_member.Result{
		Success: true,
	}, nil
}
