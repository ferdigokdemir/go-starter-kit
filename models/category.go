package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Category struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Name        string             `json:"name,omitempty"`
	Description string             `json:"description,omitempty"`
	Active      bool               `json:"active,omitempty"`
	Depth       int                `json:"depth,omitempty"`
	Image       primitive.ObjectID `json:"image" bson:"image"`
	CreatedAt   time.Time          `json:"createdAt,omitempty"`
	UpdatedAt   time.Time          `json:"updatedAt,omitempty"`
}
