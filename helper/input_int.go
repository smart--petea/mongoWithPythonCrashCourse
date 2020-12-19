package helper

import (
    "fmt"
    "strconv"
)

func InputInt(prompt string) (int, error) {
    fmt.Print(prompt)

    text, err := input()
    if err != nil {
        return 0, err
    }

    return strconv.Atoi(text)
}
