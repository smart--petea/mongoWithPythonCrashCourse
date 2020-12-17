package command

import (
    "tutorial/helper"
    "tutorial/infrastructure"
    "tutorial/model"
    "tutorial/print"

     "go.mongodb.org/mongo-driver/mongo"
     "fmt"
)

func CreateAccount(env *infrastructure.Environment, args []string) {
    var input model.CreateOwnerInput

    input.Name = helper.InputString("What is your name: ")
    input.Email = helper.InputString("What is your email: ")

    err := env.Validate.Struct(input)
    if err != nil {
        print.Error("Something wrong %s", err.Error())
        return
    }

    owner, err := env.Dataservice.FindAccountByEmail(input.Email)
    if owner != nil {
        print.Error("Account with email already exists.")
        return
    }
    if err != mongo.ErrNoDocuments {
        print.Error("Account with email %s already exists.", err.Error())
        return
    }

    owner = input.ToEntity()
    err = env.Dataservice.Save(owner)
    if err != nil {
        print.Error("Something wrong %s", err.Error())
        return
    }

    env.State.ActiveAccount = owner
    print.Success("Created a new account with id %s", owner.ID.String())
}
