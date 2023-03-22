package user

import (
	"github.com/gin-gonic/gin"
)


func ApplyRoutes(r*gin.RouterGroup){
	userRoutes := r.Group("/user")
	{

		userRoutes.GET("/", GetAllUsers)
		userRoutes.POST("/", CreateUser)
		userRoutes.GET("/:userId", GetUserById)
		userRoutes.PATCH("/:userId", UpdateUser)
		
	}
}