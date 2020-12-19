package command

import (
    "tutorial/infrastructure"
    "tutorial/helper"
    "tutorial/print"
    "tutorial/model"
)

func AddASnake(env *infrastructure.Environment, args []string) {
    var (
        input model.CreateSnakeInput
        err error
    )

    input.Name, err = helper.InputString("What is your snake's name? ")
    if err != nil {
        print.Error(err.Error())
        return
    }
    if len(input.Name) == 0 {
        print.Error("canelled")
        return
    }

    input.Length, err = helper.InputFloat("How long is your snake (in meters)? ")
    if err != nil {
        print.Error(err.Error())
        return
    }

    input.Species, err = helper.InputString("Species? ")
    if err != nil {
        print.Error(err.Error())
        return
    }

    input.IsVenomous, err = helper.InputBool("Is your snake venomous [y]es, [n]o? ", "y", "n")
    if err != nil {
        print.Error(err.Error())
        return
    }

    err = env.Validate.Struct(input)
    if err != nil {
        print.Error("Something wrong %s", err.Error())
        return
    }

    snake := input.ToEntity()
    err = env.Dataservice.Save(snake)
    if err != nil {
        print.Error("Something wrong %s", err.Error())
        return
    }

    env.State.ActiveAccount.AppendSnake(snake)
    err = env.Dataservice.Update(env.State.ActiveAccount)
    if err != nil {
        print.Error("Something wrong %s", err.Error())
        return
    }

    print.Success("Created %s with id %s", snake.Name, snake.ID.String())
}
