package helper

import (
    "bufio"
    "strings"
    "os"
)

func input() (string, error) {
    text, err := bufio.NewReader(os.Stdin).ReadString('\n')

    return strings.TrimSpace(text), err
}
