package command

import (
    "tutorial/infrastructure"
    "tutorial/print"
    "tutorial/helper"
    "tutorial/model"
)

func RegisterCage(env *infrastructure.Environment, args []string) {
    var (
        input model.CreateCageInput
        err error
    )
    input.SquareMeters, err = helper.InputFloat("How many square meters is the cage? ")
    if err != nil {
        print.Error(err.Error())
        return
    }

    input.IsCarpeted, err = helper.InputBool("Is it carpeted [y, n]?", "y", "n")
    if err != nil {
        print.Error(err.Error())
        return
    }

    //input.HasToys = helper.InputBool("Is it carpeted [y, n]?", "y", "n")
    input.AllowDangerousSnakes, err = helper.InputBool("Can you host venomous snakes [y, n]?", "y", "n")
    if err != nil {
        print.Error(err.Error())
        return
    }

    input.Name, err = helper.InputString("Give your cage a name?")
    if err != nil {
        print.Error(err.Error())
        return
    }

    input.Price, err = helper.InputFloat("How much are you charging?")
    if err != nil {
        print.Error(err.Error())
        return
    }

    err = env.Validate.Struct(input)
    if err != nil {
        print.Error("Something wrong %s", err.Error())
        return
    }

    cage := input.ToEntity()
    err = env.Dataservice.Save(cage)
    if err != nil {
        print.Error("Something wrong %s", err.Error())
        return
    }

    env.State.ActiveAccount.AppendCage(cage)
    err = env.Dataservice.Update(env.State.ActiveAccount)
    if err != nil {
        print.Error("Something wrong %s", err.Error())
        return
    }

    print.Success("Registered new cage with id %s", cage.ID.String())
}
