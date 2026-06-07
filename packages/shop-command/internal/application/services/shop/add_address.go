package shop

import (
	"context"

	"shop-command-module/internal/application/commands/add_shop_address"
	"shop-command-module/internal/application/port"
	"shop-command-module/internal/domain/member"
	"shop-command-module/internal/domain/shared"
	"shop-command-module/internal/domain/shop"
)

func (s *shopService) AddAddress(ctx context.Context, cmd add_shop_address.Command) (*add_shop_address.Result, error) {
	if err := s.authorize(ctx, cmd.ShopID, cmd.UserID, member.ActionShopManageAddress); err != nil {
		return nil, s.wrapError(err)
	}

	currentShop, err := s.shopRepo.FindByID(ctx, cmd.ShopID)
	if err != nil {
		return nil, s.wrapError(err)
	}
	if currentShop == nil {
		return nil, s.wrapError(shop.ErrShopNotFound)
	}

	addressType := shared.ValidateEnum[shop.AddressTypeEnum](cmd.Type)
	if addressType == nil {
		return nil, s.wrapError(shop.ErrAddressTypeInvalid)
	}

	newAddress, err := currentShop.AddAddress(shop.NewShopAddressParams{
		UserID: cmd.UserID,
		ShopID: currentShop.ID,

		CountryID:   cmd.Country.ID,
		CountryName: cmd.Country.Name,

		ProvinceID:   cmd.Province.ID,
		ProvinceName: cmd.Province.Name,

		WardID:   cmd.Ward.ID,
		WardName: cmd.Ward.Name,

		AddressLine: cmd.AddressLine,
		ContactName: cmd.ContactName,
		PhoneNumber: cmd.PhoneNumber,
		Type:        *addressType,
	})
	if err != nil {
		return nil, s.wrapError(err)
	}

	var outboxParams []port.OutboxParam
	if events := currentShop.FlushEvents(); len(events) > 0 {
		outboxParams = append(outboxParams, port.OutboxParam{
			AggregateID:   currentShop.ID.RawID(),
			AggregateType: currentShop.Type(),
			Events:        events,
		})
	}

	if err := s.txManager.WithTx(ctx, func(ctx context.Context) error {
		if err := s.shopRepo.CreateAddress(ctx, newAddress); err != nil {
			return err
		}

		if currentShop.IsActive() {
			if err := s.shopRepo.UpdateStatus(ctx, currentShop); err != nil {
				return err
			}
		}

		if len(outboxParams) > 0 {
			return s.outboxService.PublishBatch(ctx, outboxParams)
		}

		return nil
	}); err != nil {
		return nil, s.wrapError(err)
	}

	return &add_shop_address.Result{
		ShopAddressID: newAddress.ID.String(),
	}, nil
}
