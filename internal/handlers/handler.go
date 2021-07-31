package handlers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-boilerplate/internal/services"
)

type Handler struct {
	services *services.Service
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{
		services: service,
	}
}

func (h Handler) InitRouts() *gin.Engine {
	router := gin.New()

	api := router.Group("/")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/", h.welcome)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
