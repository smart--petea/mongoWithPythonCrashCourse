package helper

import (
    "fmt"
)

func InputInt(prompt string) int {
    var val int
    fmt.Print(prompt)
    fmt.Scanf("%d", &val)

    return val
}
