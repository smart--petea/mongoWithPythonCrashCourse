package dataservice

import (
    "tutorial/model"

    "context"
    "go.mongodb.org/mongo-driver/bson"
)

func (service *Service) UpdateCage(cage *model.Cage) error {
    filter := bson.M{"_id": cage.ID}
    update := bson.D{{"$set", cage}}

    _, err := service.Collection("cages").UpdateOne(context.Background(), filter, update)
    if err != nil {
        return err
    }

    return nil
}
