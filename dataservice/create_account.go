package dataservice

import (
    "tutorial/model"

    "time"
)

func (service *Service) CreateAccount(input *model.CreateOwnerInput) (*model.Owner, error) {
    owner := model.Owner{
        RegisteredDate: time.Now(),
        Name: input.Name,
        Email: input.Email,
    } 

    err := service.Save(&owner)
    if err != nil {
        return nil, err
    }

    return &owner, nil
}
