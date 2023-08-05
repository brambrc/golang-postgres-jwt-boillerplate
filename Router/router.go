package Router

import (
	"fmt"
	"liamelior-api/Controller"

	"github.com/gin-gonic/gin"
)

func ServeApps() {
	router := gin.Default()

	authRoutes := router.Group("/auth")
	{
		AuthRoutes(authRoutes)
	}

	router.Run(":8080")
	fmt.Println("Server is running on port 8080")
}

func AuthRoutes(router *gin.RouterGroup) {
	router.POST("/register", Controller.Register)
	router.POST("/login", Controller.Login)
}
