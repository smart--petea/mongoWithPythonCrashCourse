package command

import (
    "tutorial/infrastructure"
    "tutorial/helper"
    "tutorial/print"
)

func AddASnake(env *infrastructure.Environment, args []string) {
    var input CreateSnakeInput

    input.Name := helper.InputString("What is your snake's name? "
    if len(input.Name) == 0 {
        print.Error("canelled")
        return
    }

    input = helper.InputFloat("How long is your snake (in meters)? ")
    input.Species = helper.InputString("Species? ")
    input.IsVenomous = helper.InputBool("Is your snake venomous [y]es, [n]o? ", "y", "n")

    snake := input.ToEntity()
    err = env.Dataservice.Save(snake)
    if err != nil {
        print.Error("Something wrong %s", err.Error())
        return
    }

    print.Success("Created %s with id %s", snake.Name, snake.ID.ToString())
}
