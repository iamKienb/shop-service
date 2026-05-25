package shop

import (
	"context"
	"user-command-module/internal/application/commands/assign_member"
	domain_shared "user-command-module/internal/domain/shared"
	"user-command-module/internal/domain/shop"
)

func (s *shopService) AssignMember(ctx context.Context, cmd assign_member.Command) (*assign_member.Result, error) {
	// userRoleIDs, err := s.getUserRoles(ctx, cmd.ShopID, cmd.User.ID)
	// if err != nil {
	// 	return nil, s.wrapError(err)
	// }

	// if err := s.memberRepo.Authorize(cmd.Action, userRoleIDs); err != nil {
	// 	return nil, s.wrapError(err)
	// }

	// newMembers := make([]*member.Member, 0, len(cmd.MemberRoles))

	// for _, memberRole := range cmd.MemberRoles {
	// 	member := member.NewMember(member.NewMemberParams{
	// 		ShopID:      cmd.ShopID,
	// 		MemberID:    memberRole.ID,
	// 		MemberName:  memberRole.Name,
	// 		AddedBy:     cmd.User.ID,
	// 		NameAddedBy: cmd.User.Name,
	// 		RoleIDs:     memberRole.RoleIDs,
	// 	})

	// 	newMembers = append(newMembers, member)
	// }

	// if err := s.txManager.WithTx(ctx, func(ctx context.Context) error {
	// 	if err := s.shopRepo.ClearShopMemberRolesBatch(ctx, membersAgg); err != nil {
	// 		return err
	// 	}

	// 	if err := s.shopRepo.UpsertMemberAggregate(ctx, membersAgg); err != nil {
	// 		return err
	// 	}

	// 	var domainEvents []shared.DomainEvent
	// 	for _, memberAgg := range membersAgg {
	// 		domainEvents = append(domainEvents, memberAgg.FlushEvents()...)

	// 	}
	// 	if len(domainEvents) == 0 {
	// 		return nil
	// 	}

	// 	outboxParams := port.OutboxParam{
	// 		AggregateID:   cmd.ShopID.RawID(),
	// 		AggregateType: shop.AggregateTypeShop,
	// 		Events:        domainEvents,
	// 	}

	// 	if err := s.outboxService.Publish(ctx, outboxParams); err != nil {
	// 		return err
	// 	}

	// 	return nil
	// }); err != nil {
	// 	return nil, s.wrapError(err)
	// }

	return &assign_member.Result{
		Success: true,
	}, nil
}

func (s *shopService) getUserRoles(ctx context.Context, shopID domain_shared.ShopID, userID domain_shared.UserID) ([]domain_shared.RoleID, error) {
	data, err := s.memberRepo.GetUserRoles(ctx, shopID, userID)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, shop.ErrShopNotFound
	}

	return data.RoleIDs, nil
}
