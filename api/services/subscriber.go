package services

import (
	"genuinebasilnt/newsletter-go/api/models"
	"genuinebasilnt/newsletter-go/api/repository"
)

type SubscriptionService struct {
	subscriberRepo repository.SubscriberRepository
}

func NewSubscriptionService(repo repository.SubscriberRepository) *SubscriptionService {
	return &SubscriptionService{subscriberRepo: repo}
}

func (s *SubscriptionService) Subscribe(subscriber models.Subscriber) error {
	return s.subscriberRepo.Subscribe(&subscriber)
}
