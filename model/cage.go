package model

import (
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

//todo add validators
type CreateCageInput struct {
    Name string `validate:"required"`
    Price float32 `validate:"required"`
    SquareMeters float32 `validate:"required"`
    IsCarpeted bool `validate:"required"`
    AllowDangerousSnakes bool `validate:"required"`
}

type Cage struct {
    ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    RegisteredDate time.Time `json:"registered_date" bson:"registered_date" validate:"required"`
    Name string `json:"name" bson:"name" validate:"required"`
    Price float32 `json:"price" bson:"price" validate:"required"`
    SquareMeters float32 `json:"square_meters" bson:"square_meters" validate:"required"`
    IsCarpeted bool `json:"is_carpeted" bson:"is_carpeted" validate:"required"`
    AllowDangerousSnakes bool `json:"allow_dangerous_snakes" bson:"allow_dangerous_snakes" validate:"required"`
    Bookings []Booking `json:"bookings" bson:"bookings"`
}
