package test

import (
	"context"
	"fmt"
	"genuinebasilnt/newsletter-go/internal/config"
	"genuinebasilnt/newsletter-go/internal/env"
	"os"
	"path/filepath"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func configureDatabase(t testing.TB, config *config.Settings) *env.Env {
	t.Helper()

	connectionString := config.DatabaseSettings.ConnectionStringWithoutDB()
	pool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := pool.Exec(context.Background(), fmt.Sprintf(`CREATE DATABASE "%s";`, config.DatabaseSettings.DatabaseName)); err != nil {
		t.Fatal(err)
	}

	connectionString = config.DatabaseSettings.ConnectionString()
	pool, err = pgxpool.New(context.Background(), connectionString)
	if err != nil {
		t.Fatal(err)
	}

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	migrationsDir := filepath.Join(dir, "../migrations")

	//var embedMigrations embed.FS
	goose.SetBaseFS(nil)

	db := stdlib.OpenDBFromPool(pool)
	if err := goose.SetDialect("postgres"); err != nil {
		t.Fatal(err)
	}

	if err := goose.Up(db, migrationsDir); err != nil {
		t.Fatal(err)
	}

	return &env.Env{
		Pool: pool,
	}
}
