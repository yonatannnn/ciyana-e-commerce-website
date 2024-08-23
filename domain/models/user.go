package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID               primitive.ObjectID      `json:"id" bson:"_id"`
	FirstName        string                  `json:"first_name" bson:"first_name"`
	LastName         string                  `json:"last_name" bson:"last_name"`
	Email 			 string                  `json:"email" bson:"email"`
	Password 	     string                  `json:"password" bson:"password"`
	PendingOrders    []string                `json:"pending_orders" bson:"pending_orders"`
	Orders 			 []string                `json:"orders" bson:"orders"`
	Address          Address                 `json:"address" bson:"address"`
	CreatedAt        time.Time               `json:"created_at" bson:"created_at"`
	UpdatedAt        time.Time               `json:"updated_at" bson:"updated_at"`
}