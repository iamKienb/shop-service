package member

import (
	"context"
	"fmt"
	domain_member "user-command-module/internal/domain/member"
)

func (r *memberRepository) SaveMembers(ctx context.Context, members []*domain_member.Member) error {
	if len(members) == 0 {
		return nil
	}

	q := r.getQuerier(ctx)
	memberModels, roleModels := toInfraMemberAndRoleBatch(members)

	if err := q.AddShopMembersBatch(ctx, memberModels); err != nil {
		return fmt.Errorf("infra: add shop member batch failed: %w", err)
	}

	if err := q.AssignShopMemberRolesBatch(ctx, roleModels); err != nil {
		return fmt.Errorf("infra: assign shop member roles batch failed: %w", err)
	}

	return nil
}

func (r *memberRepository) ClearMemberRolesBatch(ctx context.Context, memberRoles []*domain_member.MemberRole) error {
	q := r.getQuerier(ctx)
	if err := q.ClearShopMemberRolesBatch(ctx, toInfraClearMemberRoles(memberRoles)); err != nil {
		return fmt.Errorf("infra: clear shop member batch failed: %w", err)
	}

	return nil
}
