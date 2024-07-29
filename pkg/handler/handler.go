package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"messaggio/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {

	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	message := router.Group("/message")
	{
		message.POST("/new", h.SaveNewMessage)
		message.GET("/stats", h.ProcessedMessagesStats)
		message.GET("/ping", h.Ping)
	}
	return router
}
