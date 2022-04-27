package user

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/InYuusha/goMongo/api/v1/database"
	"github.com/InYuusha/goMongo/api/v1/database/models"
	"github.com/InYuusha/goMongo/api/v1/user/dto/response"
	"github.com/InYuusha/goMongo/api/v1/user/dto/request"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User models.User

var userCollection *mongo.Collection = database.GetCollection(database.DB, "user")

func CreateUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	var user request.UserRequest

	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, response.Response{
			Status:  400,
			Message: "Invalid request body",
			Data: map[string]interface{}{
				"data": err.Error(),
			},
		})
		return
	}

	newUser := User{
		Id:          primitive.NewObjectID(),
		Name:        user.Name,
		Address:     user.Address,
		Dob:         user.Dob,
		Description: user.Description,
	}
	newUser.CreatedAt = time.Now()
	newUser.UpdatedAt = time.Now()

	result, err := userCollection.InsertOne(ctx, newUser)

	if err != nil {
		res := response.Response{
			Status:  500,
			Message: "Failed to create user",
		}
		c.JSON(500, res)
	}
	res := response.Response{
		Status:  200,
		Message: "User created successfully",
		Data: map[string]interface{}{
			"data": result,
		},
	}
	c.JSON(201, res)
}

func GetAllUsers(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var users []models.User
	defer cancel()

	results, err := userCollection.Find(ctx, bson.M{})

	if err != nil {
		res := response.Response{
			Status:  500,
			Message: "Failed to get users",
			Data: map[string]interface{}{
				"data": err.Error(),
			},
		}
		c.JSON(500, res)
	}
	defer results.Close(ctx)

	for results.Next(ctx) {
		var singleUser models.User
		if err = results.Decode(&singleUser); err != nil {
			c.JSON(http.StatusInternalServerError,
				response.Response{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    map[string]interface{}{"data": err.Error()}})
		}
		users = append(users, singleUser)
	}

	c.JSON(200, response.Response{
		Status:  200,
		Message: "User",
		Data: map[string]interface{}{
			"data": users,
		},
	})
}

func GetUserById(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Param("userId")
	var user models.User
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)

	err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
	if err != nil {
		log.Fatalf("Failed to get user %v", err)
		c.JSON(500, response.Response{
			Status:  500,
			Message: "Failed to get user by id ",
			Data: map[string]interface{}{
				"data": err.Error(),
			},
		})
		return
	}
	c.JSON(200, response.Response{
		Status:  200,
		Message: "User",
		Data: map[string]interface{}{
			"data": user,
		},
	})
}

func UpdateUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Param("userId")
	var user models.User
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)

	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, response.Response{
			Status:  400,
			Message: "Invalid request body",
			Data: map[string]interface{}{
				"data": err.Error(),
			},
		})
	}
	update := bson.M{
		"name":        user.Name,
		"address":     user.Address,
		"dob":         user.Dob,
		"description": user.Description,
		"updatedAt" : time.Now(),
	}
	
	result, err := userCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

	if err != nil {
		log.Fatalf("Failed to update user : %v", err)
		c.JSON(500, response.Response{
			Status:  500,
			Message: "Failed to update user",
			Data: map[string]interface{}{
				"data": err.Error(),
			},
		})
		return
	}
	var updatedUser models.User

	if result.MatchedCount == 1 {
		err = userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedUser)
		if err != nil {
			log.Fatalf("Failed to get updated user %v", err)
			c.JSON(500, response.Response{
				Status:  500,
				Message: "Failed at user/updateUser",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			})
			return
		}
	}

	c.JSON(200, response.Response{
		Status:  200,
		Message: "Successfully updated user",
		Data: map[string]interface{}{
			"data": updatedUser,
		},
	})
}
func DeleteUser(c*gin.Context){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Param("userId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)

	result,err:= userCollection.DeleteOne(ctx,bson.M{"id":objId})
	if err != nil {
		log.Fatalf("Failed to get updated user %v", err)
		c.JSON(500, response.Response{
			Status:  500,
			Message: "Failed at user/updateUser",
			Data: map[string]interface{}{
				"data": err.Error(),
			},
		})
		return
	}
	if result.DeletedCount < 1 {
		c.JSON(http.StatusNotFound,
			response.Response{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "User with specified ID not found!"}},
		)
		return
	}
	c.JSON(http.StatusOK,
		response.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "User successfully deleted!"}},
	)
}
