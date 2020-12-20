package command

import (
    "log"
    "strings"
    "fmt"

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

    CommandType(ChangeMode)(env, []string{})

    for {
        print.Prompt(env.State.ActiveAccount)
        cmd, args, err := readCommand()
        if err != nil {
            log.Fatalln(err)
        }

        switch cmd {
        case "c", "C":
            CommandType(CreateAccount).
            Header("REGISTER")(env, args)

        case "a", "A":
            CommandType(AddASnake).
            Header("Add a snake").
            ShouldBeLoggedIn()(env, args)

        case "b", "B":
            CommandType(BookACage).
            Header("Book a cage").
            ShouldBeLoggedIn()(env, args)

        case "l", "L":
            CommandType(LogIntoAccount).
            Header("LOGIN")(env, args)

        case "y", "Y":
            if env.State.Mode == infrastructure.MODE_HOST {
                CommandType(ListCages).
                Header("Your cages").
                ShouldBeLoggedIn()(env, args)
            }

            if env.State.Mode == infrastructure.MODE_GUEST {
                CommandType(ViewYourSnakes).
                Header("Your cages").
                ShouldBeLoggedIn()(env, args)
            }

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
            CommandType(ChangeMode)(env, args)

        case "f", "F":
            CommandType(ViewYourSnakes).
            Header("Your snakes").
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
        strCmdWithArgs, err := helper.InputString("")
        if err != nil {
            return "", []string{}, err
        }

        cmdWithArgs := strings.Split(strCmdWithArgs, " ")
        cmd := strings.TrimSpace(cmdWithArgs[0])

        return string(cmd), cmdWithArgs[1:], nil
} 
