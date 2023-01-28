package postgres

import (
	"context"
	"errors"
	"fmt"
	"github/architecture/config"
	"github/architecture/internal/entity"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresDB struct {
	*pgxpool.Pool
	Builder *builder
}

func New(ctx context.Context, cfg *config.Config) (*PostgresDB, error) {
	builder := NewBuilder()
	db := &PostgresDB{Builder: builder}

	err := db.connect(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (p *PostgresDB) connect(ctx context.Context, cfg *config.Config) error {
	url := p.generateUrl(cfg)

	pgConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return err
	}

	pool, err := pgxpool.ConnectConfig(ctx, pgConfig)
	if err != nil {
		return err
	}

	p.Pool = pool
	return err
}

func (p *PostgresDB) generateUrl(cfg *config.Config) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Database,
		cfg.Postgres.Sslmode,
	)
}

func (p *PostgresDB) PgError(err error) error {
	var pgErr *pgconn.PgError

	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		}
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return entity.ErrNotFound
	}
	return err
}
