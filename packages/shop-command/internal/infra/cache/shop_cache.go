package cache

import (
	"context"
	"fmt"
	"shop-command-module/internal/application/port"
	"shop-command-module/internal/domain/shared"
	"time"

	redisx "github.com/iamKienb/go-core/redis"
)

const slugKey = "shop-command:slug:%s"
const idemKey = "shop-command:idempotency:create-shop:%s"

type shopCache struct {
	cache redisx.RedisXService
}

func NewShopCache(service redisx.RedisXService) port.ShopCache {
	return &shopCache{
		cache: service,
	}
}

func (c *shopCache) IsIdemKeyTaken(ctx context.Context, userID shared.UserID) (bool, error) {
	return c.cache.Exists(ctx, fmt.Sprintf(idemKey, userID))
}

func (c *shopCache) SetIdemKey(ctx context.Context, userID shared.UserID, ttl time.Duration) error {
	return c.cache.Set(ctx, fmt.Sprintf(idemKey, userID), "exists", ttl)
}

func (c *shopCache) IsSlugKnown(ctx context.Context, slug string) (bool, error) {
	return c.cache.Exists(ctx, fmt.Sprintf(slugKey, slug))
}

func (c *shopCache) RememberSlug(ctx context.Context, slug string) error {
	return c.cache.Set(ctx, fmt.Sprintf(slugKey, slug), "exists", 0)
}
