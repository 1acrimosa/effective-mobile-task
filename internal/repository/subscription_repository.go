package repository

import (
	"context"
	"database/sql"
	"effective-mobile-task/internal/model"
)

type SubscriptionRepository struct {
	db *sql.DB
}

func NewSubscriptionRepository(db *sql.DB) *SubscriptionRepository {
	return &SubscriptionRepository{db: db}
}

func (r *SubscriptionRepository) GetAll() ([]model.Subscription, error) {
	rows, err := r.db.QueryContext(
		context.Background(),
		`SELECT id, service_name, price, user_id, start_date, end_date FROM subscriptions`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subs []model.Subscription

	for rows.Next() {
		var s model.Subscription
		err := rows.Scan(
			&s.ID,
			&s.ServiceName,
			&s.Price,
			&s.UserID,
			&s.StartDate,
			&s.EndDate,
		)
		if err != nil {
			return nil, err
		}
		subs = append(subs, s)
	}

	return subs, nil
}
