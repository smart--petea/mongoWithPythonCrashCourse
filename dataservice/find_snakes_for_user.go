package dataservice

import (
    "tutorial/model"

    "context"
    "go.mongodb.org/mongo-driver/bson"
)

func (service *Service) FindSnakesForUser(owner *model.Owner) ([]model.Snake, error) {
    filter := bson.M{"_id": bson.M{ "$in": owner.SnakeIDs}}
    cursor, err := service.Collection("snakes").Find(context.Background(), filter)
    if err != nil {
        return nil, err
    }

    var snakes []model.Snake
    err = cursor.All(context.Background(), &snakes)
    if err != nil {
        return nil, err
    }

    return snakes, nil
}
