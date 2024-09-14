package handlers

import (
	"genuinebasilnt/newsletter-go/api/models"
	"genuinebasilnt/newsletter-go/api/repository"
	"genuinebasilnt/newsletter-go/api/services"
	"genuinebasilnt/newsletter-go/internal/env"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SubscriberHandler struct {
	subscriberService *services.SubscriptionService
}

func NewSubscriberHandler(env *env.Env) *SubscriberHandler {
	repo := repository.NewPostgresSubscriberRepository(env.Pool)
	service := services.NewSubscriptionService(repo)

	return &SubscriberHandler{
		subscriberService: service,
	}
}

func (h *SubscriberHandler) Subscribe() gin.HandlerFunc {
	return func(c *gin.Context) {
		var subscriber models.Subscriber
		if err := c.Bind(&subscriber); err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		if err := h.subscriberService.Subscribe(subscriber); err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		c.Status(http.StatusOK)
	}
}
