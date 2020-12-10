package helper

import (
    "fmt"
)

func InputFloat(prompt string) float32 {
    var val float32
    fmt.Print(prompt)
    fmt.Scanf("%f", &val)

    return val
}
