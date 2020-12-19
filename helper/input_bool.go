package helper

import (
    "fmt"
    "strings"
)

func InputBool(prompt, yes, no string) (bool, error) {
    fmt.Print(prompt)

    text, err := input()
    if err != nil {
        return false, err
    }

    if strings.ToUpper(text) == strings.ToUpper(text) {
        return true, nil
    }

    return false, nil
}
