package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Item struct {
	ID               primitive.ObjectID              `json:"id" bson:"_id"`
	Name             string                          `json:"name" bson:"name"`
	Price            float64                         `json:"price" bson:"price"`
	BoughtCount      int                             `json:"bought_count" bson:"bought_count"`
	Description      string                          `json:"description" bson:"description"`
	Category         string                          `json:"category" bson:"category"`
	Images           []string                        `json:"images" bson:"images"`
	Attributes       []string                        `json:"attributes" bson:"attributes"`
	Stock			 int                              `json:"stock" bson:"stock"`
	CreatedAt        primitive.DateTime              `json:"created_at" bson:"created_at"`
	UpdatedAt        primitive.DateTime              `json:"updated_at" bson:"updated_at"`
	
	
}