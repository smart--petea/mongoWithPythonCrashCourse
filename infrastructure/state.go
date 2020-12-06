package infrastructure

type State struct {
    ActiveAccount string
}

func NewState() *State {
    return &State{}
}
