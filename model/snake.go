package model

import (
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Snake struct {
    ID             primitive.ObjectID `json:"id"              bson:"_id,omitempty"`
    RegisteredDate time.Time          `json:"registered_date" bson:"registered_date"`
    Species        string             `json:"species"         bson:"species" validate:"required"`
    Lenght         float32            `json:"length"          bson:"length"  validate:"required"`
    Name           string             `json:"name"            bson:"name"    validate:"required"`
    IsVenomous     bool               `json:"is_venomous"     bson:"is_venomous" validate:"required"`              
}
