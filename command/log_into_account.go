package command

import (
    "tutorial/infrastructure"
    "tutorial/helper"
    "tutorial/print"

    "go.mongodb.org/mongo-driver/mongo"
)

func LogIntoAccount(env *infrastructure.Environment, args []string) {
    email, err := helper.InputString("What is your email: ")
    if err != nil {
        print.Error(err.Error())
        return
    }

    owner, err := env.Dataservice.FindAccountByEmail(email)
    if err == mongo.ErrNoDocuments {
        print.Error("Could not find account with email %s", email)
        return
    }

    if err != nil {
        print.Error("Something wrong %s", err.Error())
        return
    }

    env.State.ActiveAccount = owner
    print.Success("Logged in successfully.")
}
