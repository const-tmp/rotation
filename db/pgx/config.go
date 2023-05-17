package pgx

import (
	"context"
	"fmt"
	"github.com/const-tmp/rotation/db"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Configure(db db.DB, config *pgxpool.Config) {
	config.BeforeConnect = factory(db)
}

func factory(db db.DB) func(context.Context, *pgx.ConnConfig) error {
	var funcs []func(context.Context, *pgx.ConnConfig) error

	if db.Host != nil {
		funcs = append(funcs, func(_ context.Context, config *pgx.ConnConfig) error {
			config.Host = db.Host.GetValue()
			return nil
		})
	}

	if db.Port != nil {
		funcs = append(funcs, func(_ context.Context, config *pgx.ConnConfig) error {
			var port uint16
			if _, err := fmt.Sscanf(db.Port.GetValue(), "%d", &port); err != nil {
				return err
			}
			config.Port = port
			return nil
		})
	}

	if db.User != nil {
		funcs = append(funcs, func(_ context.Context, config *pgx.ConnConfig) error {
			config.User = db.User.GetValue()
			return nil
		})
	}

	if db.Password != nil {
		funcs = append(funcs, func(_ context.Context, config *pgx.ConnConfig) error {
			config.Password = db.Password.GetValue()
			return nil
		})
	}

	return func(ctx context.Context, config *pgx.ConnConfig) error {
		for _, f := range funcs {
			if err := f(ctx, config); err != nil {
				return err
			}
		}
		return nil
	}
}
