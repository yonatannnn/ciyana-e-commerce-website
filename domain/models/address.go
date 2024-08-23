package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Address struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	Country string 		       `json:"country" bson:"country"`
	City    string 		       `json:"city" bson:"city"`
}