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

    bookings, err := env.Dataservice.GetBookingsForUser(env.State.ActiveAccount) 
    if err != nil {
        print.Error(err.Error())
        return 
    }

    print.Success("You have %d bookings.\n", len(bookings))
    for _, b := range bookings {
        print.Success(" * Snake: %s is booked at  from %s for %f days",
                snakesMap[b.GuestSnakeID].Name,
                //todo cage name,
                b.CheckInDate.Format("2006-01-02"),
                b.CheckOutDate.Sub(b.CheckInDate).Hours() / 24,
        )
    }
}
