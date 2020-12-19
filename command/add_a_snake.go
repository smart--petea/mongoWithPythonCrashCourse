package command

import (
    "tutorial/infrastructure"
    "tutorial/helper"
    "tutorial/print"
    "tutorial/model"

    "fmt"
)

func AddASnake(env *infrastructure.Environment, args []string) {
    var input model.CreateSnakeInput

    input.Name = helper.InputString("What is your snake's name? ")
    if len(input.Name) == 0 {
        print.Error("canelled")
        return
    }

    input.Length = helper.InputFloat("How long is your snake (in meters)? ")
    input.Species = helper.InputString("Species? ")
    input.IsVenomous = helper.InputBool("Is your snake venomous [y]es, [n]o? ", "y", "n")

    err := env.Validate.Struct(input)
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

    fmt.Printf("%+v\n", env.State.ActiveAccount)

    print.Success("Created %s with id %s", snake.Name, snake.ID.String())
}
