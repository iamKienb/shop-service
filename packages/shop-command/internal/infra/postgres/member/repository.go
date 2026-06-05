package member

import (
	"context"
	"errors"
	"shop-command-module/db/repository"
	domain_member "shop-command-module/internal/domain/member"

	pgx "github.com/iamKienb/go-core/postgres"
	"github.com/jackc/pgx/v5/pgconn"
)

type memberRepository struct {
	queries *repository.Queries
}

func NewRepository(service pgx.PGXService) domain_member.Repository {
	return &memberRepository{
		queries: repository.New(service.GetPool()),
	}
}
func (r *memberRepository) getQuerier(ctx context.Context) *repository.Queries {
	if tx := pgx.ExtractTx(ctx); tx != nil {
		return r.queries.WithTx(tx)
	}
	return r.queries
}

func (r *memberRepository) IsDuplicateSlug(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code == "23505" && pgErr.ConstraintName == "uq_shops_slug"
	}

	return false
}
