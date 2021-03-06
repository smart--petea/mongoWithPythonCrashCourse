package model

import (
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

//todo add validators
type CreateCageInput struct {
    Name string `validate:"required"`
    Price float64 `validate:"required"`
    SquareMeters float64 `validate:"required"`
    IsCarpeted bool
    AllowDangerousSnakes bool
}

func (input *CreateCageInput) ToEntity() *Cage {
    return &Cage{
        RegisteredDate: time.Now(),
        Name: input.Name,
        SquareMeters: input.SquareMeters,
        IsCarpeted: input.IsCarpeted,
        AllowDangerousSnakes: input.AllowDangerousSnakes,
        Price: input.Price,
    } 
}

type Cage struct {
    ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    RegisteredDate time.Time `json:"registered_date" bson:"registered_date"`
    Name string `json:"name" bson:"name"`
    Price float64 `json:"price" bson:"price"`
    SquareMeters float64 `json:"square_meters" bson:"square_meters,truncate"`
    IsCarpeted bool `json:"is_carpeted" bson:"is_carpeted"`
    AllowDangerousSnakes bool `json:"allow_dangerous_snakes" bson:"allow_dangerous_snakes"`
    Bookings []Booking `json:"bookings" bson:"bookings"`
}

func (c *Cage) Collection() string {
    return "cages"
}

func (c *Cage) GetID() primitive.ObjectID {
    return c.ID
}

func (c *Cage) SetID(id primitive.ObjectID) {
    c.ID = id
}
