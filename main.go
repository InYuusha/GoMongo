package main

import (
	"log"


	"github.com/InYuusha/goMongo/api"
	"github.com/InYuusha/goMongo/api/v1/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load environment %v", err)
	}
	//uri := os.Getenv("MONGO_URI")
	database.ConnectDB()
	//port := os.Getenv("PORT")
	app := gin.Default()
	api.ApplyRoutes(app)
	app.Run(":" + "4000")
}
