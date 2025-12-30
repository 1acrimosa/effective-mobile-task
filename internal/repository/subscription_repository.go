package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID        uuid.UUID
	Service   string
	Price     int
	UserID    uuid.UUID
	StartDate time.Time
	EndDate   *time.Time
}

type SubscriptionRepository struct {
	db *sql.DB
}

func NewSubscriptionRepository(db *sql.DB) *SubscriptionRepository {
	return &SubscriptionRepository{db: db}
}

func (r *SubscriptionRepository) Create(ctx context.Context, s Subscription) error {
	query := `
		INSERT INTO subscriptions (id, service_name, price, user_id, start_date, end_date)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		s.ID,
		s.Service,
		s.Price,
		s.UserID,
		s.StartDate,
		s.EndDate,
	)

	return err
}
