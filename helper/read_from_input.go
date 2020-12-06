package helper

import (
    "fmt"
)

func ReadFromInput(prompt string) string {
    var val string
    fmt.Print(prompt)
    fmt.Scanf("%s", &val)

    return val
}
