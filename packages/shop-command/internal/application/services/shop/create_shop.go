package shop

import (
	"context"

	"user-command-module/internal/application/commands/create_shop"
	"user-command-module/internal/application/port"
	"user-command-module/internal/application/shared"

	"user-command-module/internal/domain/member"
	domain_shared "user-command-module/internal/domain/shared"
	"user-command-module/internal/domain/shop"
)

func (s *shopService) CreateShop(ctx context.Context, cmd create_shop.Command) (*create_shop.Result, error) {
	if err := s.checkSlugAvailable(ctx, cmd.Slug); err != nil {
		return nil, s.wrapError(err)
	}

	if err := s.checkIdempotency(ctx, cmd.User.ID); err != nil {
		return nil, s.wrapError(err)
	}

	addressType := domain_shared.ValidateEnum[shop.AddressTypeEnum](cmd.Address.Type)
	if addressType == nil {
		return nil, s.wrapError(shop.ErrAddressTypeInvalid)
	}

	newShop := shop.NewShop(shop.NewShopParams{
		UserID:      cmd.User.ID,
		Name:        cmd.Name,
		Slug:        cmd.Slug,
		Description: cmd.Profile.Description,
		LogoUrl:     cmd.Profile.LogoUrl,
		BannerUrl:   cmd.Profile.BannerUrl,
		Address: shop.NewShopAddressParams{
			UserID: cmd.User.ID,

			CountryID:   cmd.Address.Country.ID,
			CountryName: cmd.Address.Country.Name,

			CityID:   cmd.Address.City.ID,
			CityName: cmd.Address.City.Name,

			DistrictID:   cmd.Address.District.ID,
			DistrictName: cmd.Address.District.Name,

			WardID:   cmd.Address.Ward.ID,
			WardName: cmd.Address.Ward.Name,

			AddressLine: cmd.Address.AddressLine,
			ContactName: cmd.Address.ContactName,
			PhoneNumber: cmd.Address.PhoneNumber,
			Type:        *addressType,
		},
	})

	newMember := member.NewMember(member.NewMemberParams{
		ShopID:      newShop.ID,
		MemberID:    cmd.User.ID,
		MemberName:  cmd.User.Name,
		AddedBy:     cmd.User.ID,
		NameAddedBy: cmd.User.Name,
		RoleIDs:     []domain_shared.RoleID{member.RoleOwnerID},
	})

	if err := s.txManager.WithTx(ctx, func(ctx context.Context) error {
		if err := s.shopRepo.CreateShop(ctx, newShop); err != nil {
			return err
		}

		if shopEvents := newShop.FlushEvents(); len(shopEvents) > 0 {
			if err := s.outboxService.Publish(ctx, port.OutboxParam{
				AggregateID:   newShop.ID.RawID(),
				AggregateType: newShop.Type(),
				Events:        shopEvents,
			}); err != nil {
				return err
			}
		}

		membersToSave := []*member.Member{newMember}
		if err := s.memberRepo.SaveMembers(ctx, membersToSave); err != nil {
			return err
		}

		if memberEvents := newMember.FlushEvents(); len(memberEvents) > 0 {
			if err := s.outboxService.Publish(ctx, port.OutboxParam{
				AggregateID:   newMember.MemberID.RawID(),
				AggregateType: newMember.Type(),
				Events:        memberEvents,
			}); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return nil, s.wrapError(err)
	}

	bgCtx := context.WithoutCancel(ctx)
	go func() {
		_ = s.shopCache.SetIdemKey(bgCtx, cmd.User.ID, shared.IdemKeyTTL)
		_ = s.shopCache.AddSlugToBloomFilter(bgCtx, cmd.Slug)
	}()

	return &create_shop.Result{
		ShopID: newShop.ID.String(),
	}, nil
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
	exists, err := s.shopCache.GetSlugFromBloomFilter(ctx, slug)
	if err != nil {
		return err
	}

	if exists == 0 {
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
