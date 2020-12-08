package dataservice

import (
     "go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
    Client *mongo.Client
}

func New(client *mongo.Client) *Service {
    return &Service{
        Client: client,
    }
}
