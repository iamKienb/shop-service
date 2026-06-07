package port

import (
	"context"
	"shop-command-module/internal/domain/shared"
	"time"
)

type ShopCache interface {
	IsSlugKnown(ctx context.Context, slug string) (bool, error)
	RememberSlug(ctx context.Context, slug string) error
	IsIdemKeyTaken(ctx context.Context, userID shared.UserID) (bool, error)
	SetIdemKey(ctx context.Context, userID shared.UserID, ttl time.Duration) error
}
