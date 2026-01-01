package repository

import (
	"context"

	"github.com/1acrimosa/effective-mobile-task/internal/model"
	"github.com/google/uuid"
)

type MemorySubscriptionRepository struct{}

func NewMemorySubscriptionRepository() *MemorySubscriptionRepository {
	return &MemorySubscriptionRepository{}
}

func (r *MemorySubscriptionRepository) Create(ctx context.Context, sub *model.Subscription) error {
	return nil
}

func (r *MemorySubscriptionRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Subscription, error) {
	return nil, nil
}
