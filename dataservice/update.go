package dataservice

import (
    "context"
    "go.mongodb.org/mongo-driver/bson"
)

func (service *Service) Update(entity EntityInt) error {
    filter := bson.M{"_id": entity.GetID()}
    update := bson.D{{"$set", entity}}

    _, err := service.
        Collection(entity.Collection()).
        UpdateOne(context.Background(), filter, update)
    if err != nil {
        return err
    }

    return nil
}
