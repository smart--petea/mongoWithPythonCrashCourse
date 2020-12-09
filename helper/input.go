package helper

import (
    "fmt"
)

func Input(prompt string) string {
    var val string
    fmt.Print(prompt)
    fmt.Scanf("%s", &val)

    return val
}
