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
    IsCarpeted bool
    AllowDangerousSnakes bool
}

type Cage struct {
    ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    RegisteredDate time.Time `json:"registered_date" bson:"registered_date"`
    Name string `json:"name" bson:"name"`
    Price float32 `json:"price" bson:"price"`
    SquareMeters float32 `json:"square_meters" bson:"square_meters,truncate"`
    IsCarpeted bool `json:"is_carpeted" bson:"is_carpeted"`
    AllowDangerousSnakes bool `json:"allow_dangerous_snakes" bson:"allow_dangerous_snakes"`
    Bookings []Booking `json:"bookings" bson:"bookings"`
}
