package shop

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"shop-command-module/internal/domain/shared"
	domain_shop "shop-command-module/internal/domain/shop"

	"github.com/iamKienb/go-core/postgres/conv"
	"github.com/jackc/pgx/v5"
)

func (r *shopRepository) CheckSlugExists(ctx context.Context, slug string) (bool, error) {
	_, err := r.getQuerier(ctx).CountBySlug(ctx, slug)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, fmt.Errorf("infra:postgres: count by slug: %w", err)
	}

	return true, nil
}

func (r *shopRepository) FindByID(ctx context.Context, shopID shared.ShopID) (*domain_shop.Shop, error) {
	q := r.getQuerier(ctx)

	shopRow, err := q.FindShopByID(ctx, conv.UUID(shopID))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("infra:postgres: find shop by id: %w", err)
	}

	profileRow, err := q.FindShopProfileByID(ctx, conv.UUID(shopID))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("infra:postgres: missing shop profile: %w", err)
		}
		return nil, fmt.Errorf("infra:postgres: find shop profile: %w", err)
	}

	addressRows, err := q.ListShopAddressesByShopID(ctx, conv.UUID(shopID))
	if err != nil {
		return nil, fmt.Errorf("infra:postgres: list shop addresses: %w", err)
	}

	return toDomainShop(shopRow, profileRow, addressRows), nil
}
