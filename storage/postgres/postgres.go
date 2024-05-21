package postgres

import (
	"context"
	"exam/config"
	"exam/storage"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	_ "github.com/lib/pq"
)

type Store struct {
	Pool *pgxpool.Pool
	cfg  config.Config
}

func New(ctx context.Context, cfg config.Config) (storage.IStorage, error) {
	url := fmt.Sprintf(`host=%s port=%v user=%s password=%s database=%s sslmode=disable`,
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDatabase)

	pgxPoolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, err
	}
	pgxPoolConfig.MaxConns = 50
	pgxPoolConfig.MaxConnLifetime = time.Hour

	newPool, err := pgxpool.NewWithConfig(ctx, pgxPoolConfig)
	if err != nil {
		return nil, err
	}

	return Store{
		Pool: newPool,
		cfg:  cfg,
	}, nil
}
func (s Store) CloseDB() {
	s.Pool.Close()
}

func (s Store) CustomersStorage() storage.CustomersStorage {
	newCustomers := NewCustomers(s.Pool)

	return &newCustomers
}
