package print

import (
    "fmt"
)

func Error(args ...string) {
    b := make([]interface{}, len(args[1:]), len(args[1:]))
    for i := range args[1:] {
        b[i] = args[i + 1]
    }

    fmt.Print("\033[31m") //red
    fmt.Printf(args[0], b...)
    fmt.Print("\033[37m") //white

    fmt.Println()
    fmt.Println()
}
