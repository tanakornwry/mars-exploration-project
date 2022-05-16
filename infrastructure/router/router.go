package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanakornwry/mars-exploration-project/registry"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	c := registry.NewRegistry()

	r.GET("/mars", c.NewAppController().Greeting)

	mainRoute := r.Group("/mars")
	{
		mainRoute.GET("/explore", c.NewAppController().Explore)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "SERVICE_NOT_FOUND", "message": "Service not available"})
	})

	return r
}
