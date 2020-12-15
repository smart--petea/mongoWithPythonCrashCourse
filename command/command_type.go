package command

import (
    "tutorial/print"
    "tutorial/infrastructure"
)

type CommandType func(env *infrastructure.Environment, args []string)

func (cmd CommandType) ShouldBeLoggedIn() CommandType {
    return func(env *infrastructure.Environment, args []string) {
        if env.State.ActiveAccount == nil {
            print.Error("You must login first in order to perform the operation")
            return
        }

        cmd(env, args)
    }
}

func (cmd CommandType) Header(header string) CommandType {
    return func(env *infrastructure.Environment, args []string) {
        print.HeadLine(header)

        cmd(env, args)
    }
}
