package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID               primitive.ObjectID      `json:"id" bson:"_id"`
	Name			 string                  `json:"name" bson:"name"`
	Price            float64                 `json:"price" bson:"price"`
	CartID		     primitive.ObjectID      `json:"cart_id" bson:"cart_id"`
}