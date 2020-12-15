package infrastructure

type CommandType func(env *Environment, args []string)
