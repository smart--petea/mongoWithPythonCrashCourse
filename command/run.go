package command

import (
    "bufio"
    "log"
    "os"
    "strings"

    "tutorial/infrastructure"
)

func Run(env *infrastructure.Environment) {
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
            CreateAccount(env, arrFullCmd)
        case "a":
            CreateAccount(env, arrFullCmd)
        case "l":
            LogIntoAccount(env, arrFullCmd)
        case "y":
            ListCages(env, arrFullCmd)
        case "r":
            RegisterCage(env, arrFullCmd)
        case "u":
            UpdateAvailibility(env, arrFullCmd)
        case "v":
            ViewBookings(env, arrFullCmd)
        case "m":
            ChangeMode(env, arrFullCmd)
        case "x", "bye", "exit", "exit()":
            ExitApp(env, arrFullCmd)
        case "?":
            ShowCommands(env, arrFullCmd)
        case "":
            None(env, arrFullCmd)
        default:
            UnknownCommand(env, arrFullCmd)
        }
    }
}
