package print

import (
    "fmt"
)

func Print(format string, args ...interface{}) {
    fmt.Printf(format, args...)
}
