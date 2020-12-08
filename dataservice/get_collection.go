package dataservice

import (
    "go.mongodb.org/mongo-driver/mongo"
)

func (service *Service) Collection(collectionName string) *mongo.Collection {
    return service.Client.Database("snake_bnb").Collection(collectionName)
}
