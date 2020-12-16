package dataservice

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type EntityInt interface {
    Collection() string
    GetID() primitive.ObjectID 
}
