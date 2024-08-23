package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Category struct {
	ID               primitive.ObjectID      `json:"id" bson:"_id"`
	Name             string                  `json:"name" bson:"name"`
	Description	  	 string                  `json:"description" bson:"description"`
	ParentID         primitive.ObjectID      `json:"parent_id" bson:"parent_id"`
	CreatedAt        time.Time      		 `json:"created_at" bson:"created_at"`
	UpdatedAt        time.Time               `json:"updated_at" bson:"updated_at"`
}