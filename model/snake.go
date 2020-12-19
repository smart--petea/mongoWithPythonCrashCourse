package model

import (
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateSnakeInput struct {
    Species string `validate:"required"`
    Length float32 `validate:"required"`
    Name string `validate:"required"`
    IsVenomous bool `validate:"required"`              
}

func (input *CreateSnakeInput) ToEntity() *Snake {
    return &Snake{
        RegisteredDate: time.Now(),
        Name: input.Name,
        Length: input.Length,
        IsVenomous: input.IsVenomous,
    } 
}

type Snake struct {
    ID             primitive.ObjectID `json:"id"              bson:"_id,omitempty"`
    RegisteredDate time.Time          `json:"registered_date" bson:"registered_date"`
    Species        string             `json:"species"         bson:"species"`
    Length         float32            `json:"length"          bson:"length"`
    Name           string             `json:"name"            bson:"name"`
    IsVenomous     bool               `json:"is_venomous"     bson:"is_venomous"`
}

func (s *Snake) Collection() string {
    return "snakes"
}

func (s *Snake) GetID() primitive.ObjectID {
    return s.ID
}

func (s *Snake) SetID(id primitive.ObjectID) {
    s.ID = id
}
