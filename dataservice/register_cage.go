package dataservice

import (
    "tutorial/model"

    "time"
)

func (service *Service) RegisterCage(owner *model.Owner, input *model.CreateCageInput) (*model.Cage, error) {
    cage := model.Cage{
        RegisteredDate: time.Now(),
        Name: input.Name,
        SquareMeters: input.SquareMeters,
        IsCarpeted: input.IsCarpeted,
        AllowDangerousSnakes: input.AllowDangerousSnakes,
        Price: input.Price,
    } 

    err := service.Save(&cage)
    if err != nil {
        return nil, err
    }

    return &cage, nil
}
