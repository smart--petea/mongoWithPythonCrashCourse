package model

import (
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Owner struct {
    ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    RegisteredDate time.Time `json:"registered_date" bson:"registered_date"`
    Name string `json:"name" bson:"name"`
    Email string `json:"email" bson:"email"`
}

type CreateOwnerInput struct {
    Name string `validate:"required"`
    Email string `validate:"required,email"`
}
