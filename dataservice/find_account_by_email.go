package dataservice

import (
    "tutorial/model"
    "context"

     "go.mongodb.org/mongo-driver/bson"
)

func (service *Service) FindAccountByEmail(email string) (*model.Owner, error) {
    var owner model.Owner

    filter := bson.D{{"email", email}}
    err := service.Collection("owners").FindOne(context.Background(), filter).Decode(&owner)
    if err != nil {
        return nil, err
    }

    return &owner, nil
}
