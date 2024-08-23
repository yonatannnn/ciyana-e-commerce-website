package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Cart struct {
	ID               primitive.ObjectID      `json:"id" bson:"_id"`
	UserID           primitive.ObjectID      `json:"user_id" bson:"user_id"`
	Items            map[string]int          `json:"items" bson:"items"`
	TotalPrice       float64                 `json:"total_price" bson:"total_price"`
}