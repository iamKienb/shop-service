package shop

import (
	"context"

	"shop-command-module/internal/application/commands/create_shop"
	"shop-command-module/internal/application/port"
	"shop-command-module/internal/application/shared"

	"shop-command-module/internal/domain/member"
	domain_shared "shop-command-module/internal/domain/shared"
	"shop-command-module/internal/domain/shop"
)

func (s *shopService) CreateShop(ctx context.Context, cmd create_shop.Command) (*create_shop.Result, error) {
	if err := s.checkSlugAvailable(ctx, cmd.Slug); err != nil {
		return nil, s.wrapError(err)
	}

	if err := s.checkIdempotency(ctx, cmd.User.ID); err != nil {
		return nil, s.wrapError(err)
	}

	newShop := shop.NewShop(shop.NewShopParams{
		UserID:      cmd.User.ID,
		Name:        cmd.Name,
		Slug:        cmd.Slug,
		Description: descriptionOf(cmd.Profile),
		LogoUrl:     logoURLOf(cmd.Profile),
		BannerUrl:   bannerURLOf(cmd.Profile),
	})

	newMember := member.NewMember(member.NewMemberParams{
		ShopID:      newShop.ID,
		MemberID:    cmd.User.ID,
		MemberName:  cmd.User.Name,
		AddedBy:     cmd.User.ID,
		NameAddedBy: cmd.User.Name,
		RoleIDs:     []domain_shared.RoleID{member.RoleOwnerID},
	})

	var outboxParams []port.OutboxParam

	if shopEvents := newShop.FlushEvents(); len(shopEvents) > 0 {
		outboxParams = append(outboxParams, port.OutboxParam{
			AggregateID:   newShop.ID.RawID(),
			AggregateType: newShop.Type(),
			Events:        shopEvents,
		})
	}

	if memberEvents := newMember.FlushEvents(); len(memberEvents) > 0 {
		outboxParams = append(outboxParams, port.OutboxParam{
			AggregateID:   newMember.MemberID.RawID(),
			AggregateType: newMember.Type(),
			Events:        memberEvents,
		})
	}

	if err := s.txManager.WithTx(ctx, func(ctx context.Context) error {
		if err := s.shopRepo.CreateShop(ctx, newShop); err != nil {
			return err
		}

		membersToSave := []*member.Member{newMember}
		if err := s.memberRepo.SaveMembers(ctx, membersToSave); err != nil {
			return err
		}

		if len(outboxParams) > 0 {
			return s.outboxService.PublishBatch(ctx, outboxParams)
		}

		return nil
	}); err != nil {
		return nil, s.wrapError(err)
	}

	bgCtx := context.WithoutCancel(ctx)
	go func() {
		_ = s.shopCache.SetIdemKey(bgCtx, cmd.User.ID, shared.IdemKeyTTL)
		_ = s.shopCache.RememberSlug(bgCtx, cmd.Slug)
	}()

	return &create_shop.Result{
		ShopID: newShop.ID.String(),
	}, nil
}

func descriptionOf(profile *create_shop.Profile) *string {
	if profile == nil {
		return nil
	}

	return profile.Description
}

func logoURLOf(profile *create_shop.Profile) *string {
	if profile == nil {
		return nil
	}

	return profile.LogoUrl
}

func bannerURLOf(profile *create_shop.Profile) *string {
	if profile == nil {
		return nil
	}

	return profile.BannerUrl
}

func (s *shopService) checkIdempotency(ctx context.Context, userID domain_shared.UserID) error {
	exists, err := s.shopCache.IsIdemKeyTaken(ctx, userID)
	if err != nil {
		return err
	}

	if exists {
		return shop.ErrShopConflict
	}

	return nil
}

func (s *shopService) checkSlugAvailable(ctx context.Context, slug string) error {
	if isKnown, err := s.shopCache.IsSlugKnown(ctx, slug); err == nil && isKnown {
		isDuplicateSlug, err := s.shopRepo.CheckSlugExists(ctx, slug)
		if err != nil {
			return err
		}
		if isDuplicateSlug {
			return shop.ErrShopSlugTaken
		}

		return nil
	}

	isDuplicateSlug, err := s.shopRepo.CheckSlugExists(ctx, slug)
	if err != nil {
		return err
	}
	if isDuplicateSlug {
		return shop.ErrShopSlugTaken
	}

	return nil
}
