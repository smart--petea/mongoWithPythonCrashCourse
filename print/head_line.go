package print

import (
    "fmt"
    "strings"
)

func HeadLine(title string) {
    stars := strings.Repeat("*", 25)
    fmt.Printf( stars + " " + title + " " + stars + "\n")
}
