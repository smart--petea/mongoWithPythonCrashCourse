package command

import (
    "tutorial/infrastructure"
    "tutorial/print"
    "tutorial/helper"
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

    print.Print("Let's start by finding available cages.")
    checkinText, err := helper.InputString("Check-in date [yyyy-mm-dd]")
    if err != nil {
        print.Error(err.Error())
        return
    }

    checkinDate, err := time.Parse("2006-01-02", checkinText)
    if err != nil {
        print.Error(err.Error())
        return
    }

    checkoutText, err := helper.InputString("Check-out date [yyyy-mm-dd]")
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
            print.Print(snake.ToLoopStringLine(i))
    }

    chosenSnakeIdx, err := helper.InputInt("Which snake do you want to book (number)?")
    if err != nil {
        print.Error(err.Error())
        return
    }

    snake := snakes[chosenSnakeIdx - 1]
    cages, err := env.Dataservice.GetAvailableCages(checkinDate, checkoutDate, snake)
    if err != nil {
        print.Error(err.Error())
        return
    }

    if len(cages) == 0 {
        print.Error("Sorry, no cages are available for that date.")
        return
    }

    print.Print("There are %d cages available in that time.", len(cages))
    //todo has_toys
    for idx, c := cages {
        carpeted := "no"
        if cage.IsCarpeted {
            carpeted = "yes"
        }

        print.Print(
            " %d. %s %dm carpeted: %s",
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
    for _, b := range cage.Bookings {
        if b.CheckinDate <= checkinDate && b.CheckoutDate >= checkoutDate && b.GuestSnakeId == nil {
            booking = b
            break
        }
    }

    booking.GuestOwnerId = env.State.env.State.ActiveAccount.ID
    booking.GuestSnakeId = snake.ID
    booking.BookedDate = time.Now()
    err := env.Dataservice.Update(cage)
    if err != nil {
        print.Error(err.Error())
        return
    }

    print.Success(
        "Successfully booked %s for  at $%f/night", 
        cage.Name,
        snake.Name,
        cage.Price,
    )
}
