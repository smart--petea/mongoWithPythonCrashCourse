package helper

import (
    "fmt"
)

func InputString(prompt string) (string, error) {
    fmt.Print(prompt)

    return input()
}
