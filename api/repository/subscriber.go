package repository

import (
	"context"
	"fmt"
	"genuinebasilnt/newsletter-go/api/models"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

type SubscriberRepository interface {
	Subscribe(subscriber *models.Subscriber) error
}

type PostgresSubscriberRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresSubscriberRepository(pool *pgxpool.Pool) *PostgresSubscriberRepository {
	return &PostgresSubscriberRepository{pool: pool}
}

func (repo *PostgresSubscriberRepository) Subscribe(subscriber *models.Subscriber) error {
	log.Info().Msgf("Adding '%s' '%s' as a new subscriber", subscriber.Name, subscriber.Email)

	query := "INSERT INTO subscriptions (id, email, name, subscribed_at) VALUES ($1, $2, $3, $4)"

	id, err := uuid.NewRandom()
	if err != nil {
		log.Error().Msg("Failed to create a unique uuid")
		return err
	}

	if _, err := repo.pool.Exec(context.Background(), query, id, subscriber.Email, subscriber.Name, time.Now().UTC()); err != nil {
		log.Error().Msg(fmt.Sprintf("Failed to execute query: %s", err))
		return err
	}

	log.Info().Msg("New subscriber details have been saved")
	return nil
}
