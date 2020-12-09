package print

import (
    "fmt"
    "strings"
)

func HeadLine(title string) {
    startLen := (60 - len(title)) / 2
    stars := strings.Repeat("*", startLen)
    fmt.Printf( stars + " " + title + " " + stars + "\n")
}
