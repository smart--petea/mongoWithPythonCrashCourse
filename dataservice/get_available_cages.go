package dataservice

import (
    "tutorial/model"

    "context"
    "go.mongodb.org/mongo-driver/bson"
    "time"
)

func (service *Service) GetAvailableCages(checkinDate, checkoutDate time.Time, snake *model.Snake) ([]model.Cage, error) {
    minSize := snake.Length / 4

    filter := bson.M{
        "square_meters": bson.M{ "$gte": minSize},
    }

    if snake.IsVenomous {
        filter["allow_dangerous_snakes"] = true
    }

    //todo $unwind for checkin and checkout
    //todo order by price, -square_meters

    cursor, err := service.Collection("cages").Find(context.Background(), filter)
    if err != nil {
        return nil, err
    }

    var cages []model.Cage
    err = cursor.All(context.Background(), &cages)
    if err != nil {
        return nil, err
    }

    var finalCages []model.Cage
    for _, c := range cages {
        for b := range c.Bookings {
            if b.CheckInDate <= checkInDate && b.CheckOutDate >= checkOutDate && b.GuestSnakeId == nil {
                finalCages = append(finalCages, cage)
            }
        }
    }

    return finalCages, nil
}
