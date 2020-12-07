package command

import (
    "bufio"
    "log"
    "os"
    "strings"
    "fmt"
    "errors"

    "tutorial/infrastructure"
)

func Run(env *infrastructure.Environment) {
    commandReader := bufio.NewReader(os.Stdin)

    printLineOfStars()

    fmt.Println()
    fmt.Println("Welcome to Snake BnB!")
    fmt.Println("Why are you here?")
    fmt.Println()
    fmt.Println("[g] Book a cage for your snake")
    fmt.Println("[h] Offer extra cage space")
    fmt.Println()
    fmt.Println("Are you a [g]uest or [h]ost?")

    cmd, args, err := readCommand(commandReader)
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
    fmt.Println(strings.Repeat("*", 19) + " Welcome " + who + " " + strings.Repeat("*", 19))
    fmt.Println()
    fmt.Println()

    printInfo()

    for {
        fmt.Print("> ")
        cmd, args, err = readCommand(commandReader)
        if err != nil {
            log.Fatalln(err)
        }

        switch cmd {
        case "c", "C", "a", "A":
            CreateAccount(env, args)
        case "l", "L":
            LogIntoAccount(env, args)
        case "y", "Y":
            ListCages(env, args)
        case "r", "R":
            RegisterCage(env, args)
        case "u", "U":
            UpdateAvailibility(env, args)
        case "v", "V":
            ViewBookings(env, args)
        case "m", "M":
            ChangeMode(env, args)
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

func printLineOfStars() {
    fmt.Println(strings.Repeat("*", 50))
}

func readCommand(commandReader *bufio.Reader) (string, []string, error) {
        strCmdWithArgs, err := commandReader.ReadString('\n')
        if err != nil {
           return "", nil, err
        }

        cmdWithArgs := strings.Split(strCmdWithArgs, " ")
        cmd := strings.TrimSpace(cmdWithArgs[0])

        return string(cmd), cmdWithArgs[1:], nil
} 

func printInfo() {
    fmt.Println("What action would you like to take:")
    fmt.Println("[C]reate an [a]ccount")
    fmt.Println("[L]ogin ton tyour account")
    fmt.Println("List [y]our cages")
    fmt.Println("[R]egister a cage")
    fmt.Println("[U]pdate cage availablity")
    fmt.Println("[V]iew your bookings")
    fmt.Println("Change [M]ode (guest or host)")
    fmt.Println("e[X]it app")
    fmt.Println("[?] Help (this info)")

    fmt.Println()
    fmt.Println()
}
