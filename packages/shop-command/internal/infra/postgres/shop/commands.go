package shop

import (
	"context"
	"fmt"
	domain_shop "shop-command-module/internal/domain/shop"
)

func (r *shopRepository) CreateShop(ctx context.Context, shop *domain_shop.Shop) error {
	q := r.getQuerier(ctx)

	if err := q.CreateShop(ctx, toInfraShop(shop)); err != nil {
		if r.IsDuplicateSlug(err) {
			return domain_shop.ErrShopSlugTaken
		}
		return fmt.Errorf("infra: save shop failed: %w", err)
	}

	if err := q.CreateShopProfile(ctx, toInfraProfile(&shop.Profile)); err != nil {
		return fmt.Errorf("infra: save profile failed: %w", err)
	}

	return nil
}

func (r *shopRepository) CreateAddress(ctx context.Context, addr *domain_shop.ShopAddress) error {
	if err := r.getQuerier(ctx).CreateShopAddress(ctx, toInfraShopAddress(addr)); err != nil {
		return fmt.Errorf("infra: save shop address failed: %w", err)
	}

	return nil
}

func (r *shopRepository) UpdateStatus(ctx context.Context, shop *domain_shop.Shop) error {
	if err := r.getQuerier(ctx).UpdateShopStatus(ctx, toInfraShopStatus(shop)); err != nil {
		return fmt.Errorf("infra: update shop status failed: %w", err)
	}

	return nil
}
