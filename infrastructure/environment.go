package infrastructure

import (
     "go.mongodb.org/mongo-driver/mongo"

     "tutorial/dataservice"
     validator "github.com/go-playground/validator/v10"
)

type Environment struct {
    Dataservice *dataservice.Service
    State *State
    Validate *validator.Validate
}

func NewEnvironment(client *mongo.Client) *Environment {
    return &Environment{
        Dataservice: dataservice.New(client),
        State: NewState(),
        Validate: validator.New(),
    }
}
