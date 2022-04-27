package main

import (
	"os"

	"github.com/InYuusha/goMongo/api"
	"github.com/InYuusha/goMongo/api/v1/database"
	"github.com/gin-gonic/gin"
)

func main() {

	//uri := os.Getenv("MONGO_URI")
	database.ConnectDB()
	port := os.Getenv("PORT")
	app := gin.Default()
	api.ApplyRoutes(app)
	app.Run(":" + port)
}
