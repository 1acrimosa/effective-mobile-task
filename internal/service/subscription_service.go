package service

import (
	"context"
	"time"

	"github.com/google/uuid"

	"effective-mobile-task/internal/model"
	"effective-mobile-task/internal/repository"
)

type SubscriptionService struct {
	repo *repository.SubscriptionRepository
}

func NewSubscriptionService(repo *repository.SubscriptionRepository) *SubscriptionService {
	return &SubscriptionService{repo: repo}
}

func (s *SubscriptionService) Create(
	ctx context.Context,
	sub *model.Subscription,
) error {
	sub.ID = uuid.New().String()
	sub.StartDate = time.Now()

	return s.repo.Create(ctx, sub)
}
