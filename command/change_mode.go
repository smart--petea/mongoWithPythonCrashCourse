package command

import (
    "tutorial/infrastructure"
    "tutorial/helper"
    "tutorial/print"
)

func ChangeMode(env *infrastructure.Environment, args []string) {
    mode, err := helper.InputString("Are you a [g]uest or [h]ost?")
    if err != nil {
        print.Error(err.Error())
        return
    }

    err = infrastructure.ValidateMode(mode)
    if err != nil {
        print.Error(err.Error())

    }

    env.State.Mode = infrastructure.Mode(mode)

    print.Info(env.State.Mode)
}
