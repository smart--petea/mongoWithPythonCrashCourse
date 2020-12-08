package infrastructure

import (
    "tutorial/model"
)

type State struct {
    ActiveAccount *model.Owner
}

func NewState() *State {
    return &State{}
}
