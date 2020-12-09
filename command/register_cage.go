package command

import (
    "tutorial/infrastructure"
    "tutorial/print"
    "tutorial/helper"
    "tutorial/model"
)

func RegisterCage(env *infrastructure.Environment, args []string) {
    print.HeadLine("REGISTER CAGE")

    if env.State.ActiveAccount == nil {
        print.Error("You must login first to register a cage.")
        return
    }

    var input model.CreateCageInput
    input.Meters = helper.InputFloat("How many square meters is the cage? ")
    input.IsCarpeted = helper.InputBool("Is it carpeted [y, n]?", "y", "n")
    //input.HasToys = helper.InputBool("Is it carpeted [y, n]?", "y", "n")
    input.AllowDangerousSnakes = helper.InputBool("Can you host venomous snakes [y, n]?", "y", "n")
    input.Name = helper.InputString("Give your cage a name")

    fmt.Println(input)
}
