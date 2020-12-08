package command

import (
    "tutorial/helper"
    "tutorial/infrastructure"
    "tutorial/model"

    "log"
)

func CreateAccount(env *infrastructure.Environment, args []string) {
    var input model.CreateOwnerInput

    input.Name = helper.ReadFromInput("What is your name: ")
    input.Email = helper.ReadFromInput("What is your email: ")

    err := env.Validate.Struct(input)
    if err != nil {
        log.Fatal(err)
    }

    owner, err := env.Dataservice.CreateAccount(&input)
    if err != nil {
        log.Fatal(err)
    }

    env.State.ActiveAccount = owner
}
