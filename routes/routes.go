package routes

import (
	"config-service/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	configs := router.Group("/configurations")
	{
		configs.POST("/", handlers.CreateConfiguration)

		configs.GET("/:name/:version", handlers.GetConfiguration)

		configs.DELETE("/:name/:version", handlers.DeleteConfiguration)
	}

	groups := router.Group("/groups")
	{
		groups.POST("/", handlers.CreateGroup)

		groups.GET("/:name/:version", handlers.GetGroup)

		groups.DELETE("/:name/:version", handlers.DeleteGroup)

		// Dodavanje i uklanjanje konfiguracija iz grupe
		groups.POST("/:name/:version/configurations", handlers.AddConfigurationToGroup)

		groups.DELETE("/:name/:version/configurations/:configName/:configVersion", handlers.RemoveConfigurationFromGroup)
	}
}
