package command

import (
    "bufio"
    "log"
    "os"
    "strings"
)

func Run() {
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
            CreateAccount(arrFullCmd)
        case "a":
            CreateAccount(arrFullCmd)
        case "l":
            LogIntoAccount(arrFullCmd)
        case "y":
            ListCages(arrFullCmd)
        case "r":
            RegisterCage(arrFullCmd)
        case "u":
            UpdateAvailibility(arrFullCmd)
        case "v":
            ViewBookings(arrFullCmd)
        case "m":
            ChangeMode(arrFullCmd)
        case "x", "bye", "exit", "exit()":
            ExitApp(arrFullCmd)
        case "?":
            ShowCommands(arrFullCmd)
        case "":
            None(arrFullCmd)
        default:
            UnknownCommand(arrFullCmd)
        }
    }
}
