package print

import (
    "fmt"
)

func Info() {
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

