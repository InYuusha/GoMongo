package apiv1

import (
	"github.com/gin-gonic/gin"
	"github.com/InYuusha/goMongo/api/v1/user"
)


func ping(c*gin.Context){
	c.JSON(200, gin.H{
		"message" : "Pong",
	})
}

func ApplyRoutes(r*gin.RouterGroup){
	v1:= r.Group("/v1")
	{
		// all routes
		user.ApplyRoutes(v1)
	}
	
}