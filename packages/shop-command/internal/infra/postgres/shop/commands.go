package shop

import (
	"context"
	"fmt"
	domain_shop "user-command-module/internal/domain/shop"
)

func (r *shopRepository) CreateShop(ctx context.Context, shop *domain_shop.Shop) error {
	q := r.getQuerier(ctx)

	if err := q.SaveShop(ctx, toInfraShop(shop)); err != nil {
		if r.IsDuplicateSlug(err) {
			return domain_shop.ErrShopSlugTaken
		}
		return fmt.Errorf("infra: save shop failed: %w", err)
	}

	if err := q.SaveShopProfile(ctx, toInfraProfile(&shop.Profile)); err != nil {
		return fmt.Errorf("infra: save profile failed: %w", err)
	}

	for i := range shop.Addresses {
		if err := q.SaveShopAddress(ctx, toInfraShopAddress(&shop.Addresses[i])); err != nil {
			return fmt.Errorf("infra: save shop address failed: %w", err)
		}
	}

	return nil
}

func (r *shopRepository) CreateAddress(ctx context.Context, addr *domain_shop.ShopAddress) error {
	if err := r.getQuerier(ctx).SaveShopAddress(ctx, toInfraShopAddress(addr)); err != nil {
		return fmt.Errorf("infra: save shop address failed: %w", err)
	}

	return nil
}
