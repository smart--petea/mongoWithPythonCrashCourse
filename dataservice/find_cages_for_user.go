package dataservice

import (
    "tutorial/model"

    "context"
    "go.mongodb.org/mongo-driver/bson"
)

func (service *Service) FindCagesForUser(owner *model.Owner) ([]model.Cage, error) {
    filter := bson.M{"_id": bson.M{ "$in": owner.CageIDs}}
    cursor, err := service.Collection("cages").Find(context.Background(), filter)
    if err != nil {
        return nil, err
    }

    var cages []model.Cage
    err = cursor.All(context.Background(), &cages)
    if err != nil {
        return nil, err
    }

    return cages, nil
}
