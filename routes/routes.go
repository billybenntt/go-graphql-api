package routes

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes defines and registers all API routes for the server.
func RegisterRoutes(server *gin.Engine) {

	// GraphQL Routes
	server.POST("/api", graphqlHandler())

	// Playground WebUI
	server.GET("/play", playgroundHandler())

}
