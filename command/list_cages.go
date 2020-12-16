package command

import (
    "tutorial/infrastructure"
    "tutorial/print"
)

func ListCages(env *infrastructure.Environment, args []string) {
    cages, err := env.Dataservice.FindCagesForUser(env.State.ActiveAccount) 
    if err != nil {
        print.Error(err.Error())
        return 
    }

    print.Success("You have %d cages", len(cages))

    for i, cage := range cages {
        print.Success(" %d. %s is %f meters.", (i + 1), cage.Name, cage.SquareMeters)

        for _, b := range cage.Bookings {
            bookedLabel := "YES"
            if b.BookedDate.IsZero() {
                bookedLabel = "no"
            }
            print.Success(
                "\t\t* Booking: %s, %d days, booked? %s",
                b.CheckInDate.Format("2006-01-02"),
                int(b.CheckOutDate.Sub(b.CheckInDate).Hours() / 24),
                bookedLabel,
            )
        }
    }
}
