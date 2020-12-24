package command

import (
    "tutorial/infrastructure"
    "tutorial/print"

    "os"
)

func ExitApp(env *infrastructure.Environment, args []string) {
    print.Print("\n")
    print.Success("bye!\n")

    os.Exit(0)
}
