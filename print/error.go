package print

import (
    "fmt"
)

func Error(format string, args ...interface{}) {
    fmt.Print("\033[31m") //red
    fmt.Printf(format, args...)
    fmt.Print("\033[37m") //white

    fmt.Println()
    fmt.Println()
}
