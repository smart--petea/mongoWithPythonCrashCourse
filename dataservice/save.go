package dataservice

import (
    "context"

    "go.mongodb.org/mongo-driver/bson/primitive"
)

func (service *Service) Save(entity EntityInt) error {
    res, err := service.
        Collection(entity.Collection()).
        InsertOne(context.Background(), entity)
    if err != nil {
        return err
    }

    entity.SetID(res.InsertedID.(primitive.ObjectID))

    return nil
}
