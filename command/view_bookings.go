package command

import (
    "tutorial/infrastructure"
    "tutorial/print"
    "tutorial/model"

    "go.mongodb.org/mongo-driver/bson/primitive"
)

func ViewBookings(env *infrastructure.Environment, args []string) {
    snakes, err := env.Dataservice.FindSnakesForUser(env.State.ActiveAccount) 
    if err != nil {
        print.Error(err.Error())
        return 
    }

    var snakesMap map[primitive.ObjectID]*model.Snake = map[primitive.ObjectID]*model.Snake{}
    for i := range snakes {
        snakesMap[snakes[i].ID] = &snakes[i]
    }

    cages, err := env.Dataservice.GetBookedCagesForUser(env.State.ActiveAccount) 
    if err != nil {
        print.Error(err.Error())
        return 
    }

    var countBookings int
    for _, c := range cages {
        for i := range c.Bookings {
            if c.Bookings[i].GuestOwnerID == env.State.ActiveAccount.ID {
                countBookings = countBookings + 1
            }
        }
    }

    print.Success("You have %d bookings.\n", countBookings)
    for _, c := range cages {
        for i := range c.Bookings {
            b := c.Bookings[i]
            if b.GuestOwnerID != env.State.ActiveAccount.ID {
                continue
            }

            print.Success(" * Snake: %s is booked at %s from %s for %f days",
                    snakesMap[b.GuestSnakeID].Name,
                    c.Name,
                    b.CheckInDate.Format("2006-01-02"),
                    b.CheckOutDate.Sub(b.CheckInDate).Hours() / 24,
            )
        }
    }
}
