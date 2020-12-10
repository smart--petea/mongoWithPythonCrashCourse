package helper

import (
    "fmt"
    "strings"
)

func InputBool(prompt, yes, no string) bool {
    var val string
    fmt.Print(prompt)
    fmt.Scanf("%s", &val)

    if strings.ToUpper(val) == strings.ToUpper(yes) {
        return true
    }

    return false
}
