package router

import (
	"genuinebasilnt/newsletter-go/api/handlers"
	"genuinebasilnt/newsletter-go/internal/env"
	"genuinebasilnt/newsletter-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Router(env *env.Env) *gin.Engine {
	router := gin.New()
	router.Use(middleware.RequestLogger(env.Logger))
	router.Use(gin.Recovery())

	router.GET("/health_check", handlers.HealthCheck)
	router.POST("/subscriptions", handlers.NewSubscriberHandler(env).Subscribe())

	return router
}
