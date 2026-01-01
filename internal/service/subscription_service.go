package service

import (
	"context"

	"github.com/1acrimosa/effective-mobile-task/internal/model"
	"github.com/1acrimosa/effective-mobile-task/internal/repository"
	"github.com/google/uuid"
)

type SubscriptionService struct {
	repo repository.SubscriptionRepository
}

func NewSubscriptionService(repo repository.SubscriptionRepository) *SubscriptionService {
	return &SubscriptionService{repo: repo}
}

func (s *SubscriptionService) Create(ctx context.Context, sub *model.Subscription) error {
	sub.ID = uuid.New()
	return s.repo.Create(ctx, sub)
}
