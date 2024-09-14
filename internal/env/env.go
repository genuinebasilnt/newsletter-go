package env

import (
	"context"
	"genuinebasilnt/newsletter-go/internal/config"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type Env struct {
	Pool   *pgxpool.Pool
	Logger zerolog.Logger
}

func SetupEnv() (*Env, error) {
	configuration, err := config.GetConfiguration("config")
	if err != nil {
		return nil, err
	}

	connectionString := configuration.DatabaseSettings.ConnectionString()
	conn, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		return nil, err
	}

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	return &Env{
		Pool:   conn,
		Logger: logger,
	}, nil
}
