package api

import (
	"github.com/gin-gonic/gin"
	"github.com/InYuusha/goMongo/api/v1"
)

func ApplyRoutes(r *gin.Engine) {

	api := r.Group("/api")
	{
		apiv1.ApplyRoutes(api)
	}
}
