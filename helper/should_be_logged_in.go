package helper

import (
    "tutorial/infrastructure"
    "tutorial/print"
)

func ShouldBeLoggedIn(safeCommand infrastructure.CommandType) infrastructure.CommandType {
    return func(env *infrastructure.Environment, args []string) {
        if env.State.ActiveAccount == nil {
            print.Error("You must login first in order to perform the operation")
            return
        }

        safeCommand(env, args)
    }
}
