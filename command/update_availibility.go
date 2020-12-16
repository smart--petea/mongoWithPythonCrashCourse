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
    cageNumber := helper.InputInt("Enter cage number: ")

    cages, err := env.Dataservice.FindCagesForUser(env.State.ActiveAccount) 
    if err != nil {
        print.Error(err.Error())
        return 
    }

    selectedCage := cages[cageNumber - 1]
    print.Success("Selected cage %s", selectedCage.Name)

    checkInDateStr := helper.InputString("Enter available date [yyyy-mm-dd]: ")
    checkInDate, err := time.Parse("2006-01-02", checkInDateStr)
    if err != nil {
        print.Error(err.Error())
        return
    }

    days := helper.InputInt("How many days is this block of time? ")

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
