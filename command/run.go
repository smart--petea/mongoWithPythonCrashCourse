package command

import (
    "log"
    "strings"
    "fmt"
    "errors"

    "tutorial/infrastructure"
    "tutorial/print"
    "tutorial/helper"
)

func Run(env *infrastructure.Environment) {
    print.LineOfStars()

    fmt.Println()
    fmt.Println("Welcome to Snake BnB!")
    fmt.Println("Why are you here?")
    fmt.Println()
    fmt.Println("[g] Book a cage for your snake")
    fmt.Println("[h] Offer extra cage space")
    fmt.Println()
    fmt.Println("Are you a [g]uest or [h]ost?")

    cmd, args, err := readCommand()
    if err != nil {
        log.Fatalln(err)
    }

    who := ""
    switch cmd {
        case "g": 
            who = "guest"
        case "h":
            who = "host"
        default:
            log.Fatalln(errors.New("wrong command. Please, start once again."))
    }
    print.HeadLine("Welcome " + who)
    fmt.Println()

    print.Info()

    for {
        print.Prompt(env.State.ActiveAccount)
        cmd, args, err = readCommand()
        if err != nil {
            log.Fatalln(err)
        }

        switch cmd {
        case "c", "C", "a", "A":
            CommandType(CreateAccount).
            Header("REGISTER")(env, args)
        case "l", "L":
            CommandType(LogIntoAccount).
            Header("LOGIN")(env, args)
        case "y", "Y":
            CommandType(ListCages).
            Header("Your cages").
            ShouldBeLoggedIn()(env, args)
        case "r", "R":
            CommandType(RegisterCage).
            Header("REGISTER CAGE").
            ShouldBeLoggedIn()(env, args)
        case "u", "U":
            CommandType(UpdateAvailibility).
            Header("Add available date").
            ShouldBeLoggedIn()(env, args)
        case "v", "V":
            ViewBookings(env, args)
        case "m", "M":
            ChangeMode(env, args)
        case "s", "S":
            CommandType(AddASnake).
            Header("Add a snake").
            ShouldBeLoggedIn()(env, args)
        case "x", "bye", "exit", "exit()":
            ExitApp(env, args)
        case "?":
            ShowCommands(env, args)
        case "":
            None(env, args)
        default:
            UnknownCommand(env, args)
        }
    }
}

func readCommand() (string, []string, error) {
        strCmdWithArgs := helper.InputString("")

        cmdWithArgs := strings.Split(strCmdWithArgs, " ")
        cmd := strings.TrimSpace(cmdWithArgs[0])

        return string(cmd), cmdWithArgs[1:], nil
} 
