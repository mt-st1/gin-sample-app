package router

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// InitializeRouter ルーティングの初期化
func InitializeRouter() *gin.Engine {
	// Set the router as the default
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	apiGroup := router.Group("/api")
	apiRegister(apiGroup)
	return router
}

// /api
func apiRegister(group *gin.RouterGroup) {
	group.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
