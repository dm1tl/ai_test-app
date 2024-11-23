package handler

import (
	"ai_test-app/internal/services"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *services.Service
}

func NewHandler(serv *services.Service) *Handler {
	return &Handler{
		service: serv,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api", h.userIdentity)
	{
		api.POST("/gentest", h.genTest)
		api.POST("/answtest", h.answTest)
	}
	return router
}
