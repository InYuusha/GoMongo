package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct{
	Id          primitive.ObjectID   `json:"id"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	Dob         string    `json:"dob"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt 	time.Time `json:"updated_at"`
}

