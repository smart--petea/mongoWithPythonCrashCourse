package print

import (
    "fmt"
    "tutorial/infrastructure"
)

func Info(mode infrastructure.Mode) {
    fmt.Println()
    fmt.Println()

    switch infrastructure.Mode(mode) {
    case infrastructure.MODE_HOST:
        HeadLine("HOST MODE")

        fmt.Println("What action would you like to take:")
        fmt.Println("[C]reate an account")
        fmt.Println("[L]ogin to your account")
        fmt.Println("List [y]our cages")
        fmt.Println("[R]egister a cage")
        fmt.Println("[U]pdate cage availablity")
        fmt.Println("[V]iew your bookings")
        fmt.Println("Change [M]ode (guest or host)")
        fmt.Println("e[X]it app")
        fmt.Println("[?] Help (this info)")

    case infrastructure.MODE_GUEST:
        HeadLine("GUEST MODE")

        fmt.Println("What action would you like to take:")
        fmt.Println("[C]reate an account")
        fmt.Println("[L]ogin to your account")
        fmt.Println("[B]ook a cage")
        fmt.Println("[A]dd a snake")
        fmt.Println("View [y]our snakes")
        fmt.Println("[V]iew your bookings")
        fmt.Println("Change [M]ode (guest or host)")
        fmt.Println("e[X]it app")
        fmt.Println("[?] Help (this info)")

    default:
        fmt.Printf("Wrong mode %s", mode)
    }
    

    fmt.Println()
    fmt.Println()
}
