package infrastructure

import (
    "fmt"
)

type Mode string

const (
    MODE_GUEST Mode = "g"
    MODE_HOST Mode = "h"
)

func ValidateMode(mode string) error {
    switch Mode(mode) {
    case MODE_GUEST, MODE_HOST:
        return nil
    default:
        return fmt.Errorf("Invalid mode %s", mode)
    }
}
