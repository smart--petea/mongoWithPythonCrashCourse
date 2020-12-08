package dataservice

import (
    "tutorial/model"

    "time"
    "context"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func (service *Service) CreateAccount(input *model.CreateOwnerInput) (*model.Owner, error) {
    owner := model.Owner{
        RegisteredDate: time.Now(),
        Name: input.Name,
        Email: input.Email,
    } 

    res, err := service.Collection("owners").InsertOne(context.Background(), owner)
    if err != nil {
        return nil, err
    }

    owner.ID = res.InsertedID.(primitive.ObjectID)

    return &owner, nil
}
