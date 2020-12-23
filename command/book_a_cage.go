package command

import (
    "tutorial/infrastructure"
    "tutorial/print"
    "tutorial/helper"
    "tutorial/model"

    "time"
)

func BookACage(env *infrastructure.Environment, args []string) {
    snakes, err := env.Dataservice.FindSnakesForUser(env.State.ActiveAccount) 
    if err != nil {
        print.Error(err.Error())
        return 
    }

    if len(snakes) == 0 {
        print.Error("You must [a]dd a snake before you can book a cage.")
        return
    }

    print.Print("Let's start by finding available cages.\n")
    checkinText, err := helper.InputString("Check-in date [yyyy-mm-dd]? ")
    if err != nil {
        print.Error(err.Error())
        return
    }

    checkinDate, err := time.Parse("2006-01-02", checkinText)
    if err != nil {
        print.Error(err.Error())
        return
    }

    checkoutText, err := helper.InputString("Check-out date [yyyy-mm-dd]? ")
    if err != nil {
        print.Error(err.Error())
        return
    }

    checkoutDate, err := time.Parse("2006-01-02", checkoutText)
    if err != nil {
        print.Error(err.Error())
        return
    }

    if checkinDate.Before(checkoutDate) == false {
        print.Error("Check in must be before check out")
        return
    }

    for idx, snake := range snakes {
            print.Success(snake.ToLoopStringLine(idx) + "\n")
    }

    chosenSnakeIdx, err := helper.InputInt("Which snake do you want to book (number)?")
    if err != nil {
        print.Error(err.Error())
        return
    }

    snake := snakes[chosenSnakeIdx - 1]
    cages, err := env.Dataservice.GetAvailableCages(checkinDate, checkoutDate, &snake)
    if err != nil {
        print.Error(err.Error())
        return
    }

    if len(cages) == 0 {
        print.Error("Sorry, no cages are available for that date.")
        return
    }

    print.Print("There are %d cages available in that time.\n", len(cages))
    for idx, c := range cages {
        carpeted := "no"
        if c.IsCarpeted {
            carpeted = "yes"
        }

        print.Print(
            " %d. %s %fm carpeted: %s\n",
            idx + 1,
            c.Name,
            c.SquareMeters,
            carpeted,
        )
    }

    chosenCageIdx, err := helper.InputInt("Which cage do you want to book (number)")
    if err != nil {
        print.Error(err.Error())
        return
    }

    cage := cages[chosenCageIdx - 1]
    var booking *model.Booking
    for i := range cage.Bookings {
        b := &cage.Bookings[i]
        if b.CheckInDate.Before(checkinDate) && b.CheckOutDate.After(checkoutDate) && b.GuestSnakeID.IsZero() {
            booking = b
            break
        }
    }

    booking.GuestOwnerID = env.State.ActiveAccount.ID
    booking.GuestSnakeID = snake.ID
    booking.BookedDate = time.Now()
    err = env.Dataservice.Update(&cage)
    if err != nil {
        print.Error(err.Error())
        return
    }

    print.Success(
        "Successfully booked cage %s for snake %s at $%f/night", 
        cage.Name,
        snake.Name,
        cage.Price,
    )
}
