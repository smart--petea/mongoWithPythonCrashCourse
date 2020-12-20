package model

import (
    "fmt"
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateSnakeInput struct {
    Species string `validate:"required"`
    Length float64 `validate:"required"`
    Name string `validate:"required"`
    IsVenomous bool
}

func (input *CreateSnakeInput) ToEntity() *Snake {
    return &Snake{
        RegisteredDate: time.Now(),
        Name: input.Name,
        Length: input.Length,
        IsVenomous: input.IsVenomous,
        Species: input.Species,
    } 
}

type Snake struct {
    ID             primitive.ObjectID `json:"id"              bson:"_id,omitempty"`
    RegisteredDate time.Time          `json:"registered_date" bson:"registered_date"`
    Species        string             `json:"species"         bson:"species"`
    Length         float64            `json:"length"          bson:"length"`
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

func (s *Snake) ToLoopStringLine(idx int) string {
    venomous := "not"
    if s.IsVenomous {
        venomous = ""
    }

    return fmt.Sprintf(
        " %d. %s is a %s that is %fm  and is %s venomous.",
        (idx + 1),
        s.Name,
        s.Species,
        s.Length,
        venomous,
    )
}
