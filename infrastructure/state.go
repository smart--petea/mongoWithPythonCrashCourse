package infrastructure

import (
    "tutorial/model"
)

type State struct {
    ActiveAccount *model.Owner
    Mode Mode
}

func NewState() *State {
    return &State{}
}
