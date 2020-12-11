package dataservice

import (
    "tutorial/model"

    "context"
    "go.mongodb.org/mongo-driver/bson"
)

func (service *Service) UpdateOwner(owner *model.Owner) error {
    filter := bson.M{"_id": owner.ID}
    update := bson.D{{"$set", owner}}

    _, err := service.Collection("owners").UpdateOne(context.Background(), filter, update)
    if err != nil {
        return err
    }

    return nil
}
