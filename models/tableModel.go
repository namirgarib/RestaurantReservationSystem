package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Table struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	NumOfGuests *int               `json:"num_of_guests" validate:"required"`
	TableNumber *int               `json:"table_number" validate:"required"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
	TableID     string             `json:"table_id" validate:"required"`
}
