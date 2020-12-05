package model

import (
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Owner struct {
    RegisteredDate time.Time `json:"registered_date" bson:"registered_date"`
    Name string `json:"name" bson:"name" validate:"required"`
    Email string `json:"email" bson:"email" validate:"required,email"`
}
