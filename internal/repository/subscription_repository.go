package repository

import (
	"context"
	"github.com/1acrimosa/effective-mobile-task/internal/model"
	"github.com/google/uuid"
)

type SubscriptionRepository interface {
	Create(ctx context.Context, sub *model.Subscription) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Subscription, error)
}
