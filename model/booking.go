package model

import (
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Booking struct {
    GuestOwnerID primitive.ObjectID `json:"guest_owner_id" bson:"guest_owner_id"`
    GuestSnakeID primitive.ObjectID `json:"guest_snake_id" bson:"guest_snake_id"`
    BookedDate time.Time `json:"booked_date,omitempty" bson:"booked_date,omitempty"`
    CheckInDate time.Time `json:"check_in_date" bson:"check_in_date" `//validate:"required"
    CheckOutDate time.Time `json:"check_out_date" bson:"check_out_date"`//validate:"required"
    Review string `json:"review,omitempty" bson:"review,omitempty"`
    Rating int `json:"rating" bson:"rating"`
}
