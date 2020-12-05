package main

import (
    "bufio"
    "log"
    "os"
    "strings"

    "tutorial/command"
)

func main() {
    commandReader := bufio.NewReader(os.Stdin)

    for {
        strFullCmd, err := commandReader.ReadString('\n')
        if err != nil {
            log.Fatalln(err)
        }

        arrFullCmd := strings.Split(strFullCmd, " ")
        cmd := strings.TrimSpace(arrFullCmd[0])

        switch cmd {
        case "c":
            command.CreateAccount(arrFullCmd)
        case "a":
            command.CreateAccount(arrFullCmd)
        case "l":
            command.LogIntoAccount(arrFullCmd)
        case "y":
            command.ListCages(arrFullCmd)
        case "r":
            command.RegisterCage(arrFullCmd)
        case "u":
            command.UpdateAvailibility(arrFullCmd)
        case "v":
            command.ViewBookings(arrFullCmd)
        case "m":
            command.ChangeMode(arrFullCmd)
        case "x", "bye", "exit", "exit()":
            command.ExitApp(arrFullCmd)
        case "?":
            command.ShowCommands(arrFullCmd)
        case "":
            command.None(arrFullCmd)
        default:
            command.UnknownCommand(arrFullCmd)
        }
    }
}
