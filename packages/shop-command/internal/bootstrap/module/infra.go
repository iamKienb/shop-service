package module

import (
	"context"
	"fmt"

	"user-command-module/internal/application/port"
	"user-command-module/internal/bootstrap/config"
	"user-command-module/internal/domain/member"
	"user-command-module/internal/domain/shop"
	"user-command-module/internal/infra/cache"
	memberPg "user-command-module/internal/infra/postgres/member"
	outboxPg "user-command-module/internal/infra/postgres/outbox"
	shopPg "user-command-module/internal/infra/postgres/shop"

	pgx "github.com/iamKienb/go-core/postgres"
	redisx "github.com/iamKienb/go-core/redis"
)

type InfraModule struct {
	PGService    pgx.PGXService
	RedisService redisx.RedisXService

	OutboxRepo port.OutboxRepository
	ShopRepo   shop.Repository
	MemberRepo member.Repository
	ShopCache  port.ShopCache
	TxManager  port.TxManager
}

func NewInfraModule(ctx context.Context, cfg *config.ShopCommandConfig) (*InfraModule, error) {
	pgService, err := pgx.New(cfg.Postgres)
	if err != nil {
		return nil, fmt.Errorf("postgres: %w", err)
	}

	redisService, err := redisx.New(cfg.Redis)
	if err != nil {
		return nil, fmt.Errorf("redis: %w", err)
	}

	return &InfraModule{
		PGService:    pgService,
		RedisService: redisService,

		OutboxRepo: outboxPg.NewRepository(pgService),
		ShopRepo:   shopPg.NewRepository(pgService),
		MemberRepo: memberPg.NewRepository(pgService),
		ShopCache:  cache.NewShopCache(redisService),
		TxManager:  pgx.NewTxManager(pgService.GetPool()),
	}, nil
}
