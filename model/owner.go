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
    CageIDs  []primitive.ObjectID `json:"cage_ids" bson:"cage_ids"`
    SnakeIDs []primitive.ObjectID `json:"cage_ids" bson:"cage_ids"`
}

type CreateOwnerInput struct {
    Name string `validate:"required"`
    Email string `validate:"required,email"`
}

func (input *CreateOwnerInput) ToEntity() *Owner {
    return &Owner{
        RegisteredDate: time.Now(),
        Name: input.Name,
        Email: input.Email,
    } 
}

func (o *Owner) AppendSnake(s *Snake) {
    o.SnakeIDs = append(o.SnakeIDs, s.ID)
}

func (o *Owner) AppendCage(c *Cage) {
    o.CageIDs = append(o.CageIDs, c.ID)
}

func (o *Owner) Collection() string {
    return "owners"
}

func (o *Owner) GetID() primitive.ObjectID {
    return o.ID
}

func (o *Owner) SetID(id primitive.ObjectID) {
    o.ID = id
}
