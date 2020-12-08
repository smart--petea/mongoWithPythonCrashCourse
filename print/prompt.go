package print

import (
    "fmt"

    "tutorial/model"
)

func Prompt(owner *model.Owner) {
    fmt.Print("\033[33m") //yellow

    if owner == nil {
        fmt.Print("> ")
    } else {
        fmt.Printf("%s> ", owner.Name)
    }

    fmt.Print("\033[37m") //white
}
