package router

import (
	"github.com/gin-gonic/gin"
	"learn-go/web-service/api/controllers"
)

var R *gin.Engine

func StartServer() {
	// Create router
	r := gin.Default()

	// Create router group for end points with common path prefix
	g := r.Group("/api")

	// Routes
	g.GET("/user", controllers.FindUsers)
	g.POST("/user", controllers.CreateUser)
	g.GET("/user/:id", controllers.FindUser)
	g.PATCH("/user/:id", controllers.UpdateUser)
	g.DELETE("/user/:id", controllers.DeleteUser)

	// Start server
	if err := r.Run(); err != nil {
		panic("failed to start server")
	}

	R = r
}
