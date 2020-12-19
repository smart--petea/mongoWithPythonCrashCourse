package command

import (
    "tutorial/infrastructure"
    "tutorial/helper"
    "tutorial/print"
    "tutorial/model"

    "time"
)

func UpdateAvailibility(env *infrastructure.Environment, args []string) {
    ListCages(env, args)
    cageNumber, err := helper.InputInt("Enter cage number: ")
    if err != nil {
        print.Error(err.Error())
        return
    }

    cages, err := env.Dataservice.FindCagesForUser(env.State.ActiveAccount) 
    if err != nil {
        print.Error(err.Error())
        return 
    }

    selectedCage := cages[cageNumber - 1]
    print.Success("Selected cage %s", selectedCage.Name)

    checkInDateStr, err := helper.InputString("Enter available date [yyyy-mm-dd]: ")
    if err != nil {
        print.Error(err.Error())
        return
    }

    checkInDate, err := time.Parse("2006-01-02", checkInDateStr)
    if err != nil {
        print.Error(err.Error())
        return
    }

    days, err := helper.InputInt("How many days is this block of time? ")
    if err != nil {
        print.Error(err.Error())
        return
    }

    booking := model.Booking{}
    booking.CheckInDate = checkInDate
    booking.CheckOutDate = checkInDate.AddDate(0, 0, days)

    selectedCage.Bookings = append(selectedCage.Bookings, booking)
    err = env.Dataservice.Update(&selectedCage)
    if err != nil {
        print.Error(err.Error())
        return
    }

    print.Success("Date added to cage %s", selectedCage.Name)
}
