package dataservice

import (
    "tutorial/model"

    "time"
    "context"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func (service *Service) RegisterCage(owner *model.Owner, input *model.CreateCageInput) (*model.Cage, error) {
    cage := model.Cage{
        RegisteredDate: time.Now(),
        Name: input.Name,
        SquareMeters: input.SquareMeters,
        IsCarpeted: input.IsCarpeted,
        AllowDangerousSnakes: input.AllowDangerousSnakes,
    } 

    res, err := service.Collection("cages").InsertOne(context.Background(), cage)
    if err != nil {
        return nil, err
    }

    cage.ID = res.InsertedID.(primitive.ObjectID)

    return &cage, nil
}
