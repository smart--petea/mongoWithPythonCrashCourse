package print

import (
    "fmt"
)

func Success(format string, args... interface{}) {
    fmt.Print("\033[32m") //green
    fmt.Printf(format, args...)
    fmt.Print("\033[37m") //white

    fmt.Println()
    fmt.Println()
}
