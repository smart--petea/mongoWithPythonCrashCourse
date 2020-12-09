package command

import (
    "tutorial/infrastructure"
    "tutorial/print"
    //"tutorial/helper"
)

func RegisterCage(env *infrastructure.Environment, args []string) {
    print.HeadLine("REGISTER CAGE")

    if env.State.ActiveAccount == nil {
        print.Error("You must login first to register a cage.")
        return
    }

    //meters := helper.Input("How many square meters is the cage? ")
}
