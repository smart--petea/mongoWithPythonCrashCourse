package helper

import (
    "fmt"
    "strconv"
)

func InputFloat(prompt string) (float64, error) {
    fmt.Print(prompt)

    text, err := input()
    if err != nil {
        return 0, err
    }

    return strconv.ParseFloat(text, 32)
}
