package model

import (
	
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
	Watched bool `json:"watched,omitempty" bson:"watched,omitempty"`

}