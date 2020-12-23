package dataservice

import (
    "tutorial/model"

    "context"
    "go.mongodb.org/mongo-driver/bson"
)

func (service *Service) GetBookedCagesForUser(user *model.Owner) ([]model.Cage, error) {

    filter := bson.M{
        "bookings.guest_owner_id": user.ID,
    }

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
