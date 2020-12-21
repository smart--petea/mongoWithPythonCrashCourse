package dataservice

import (
    "tutorial/model"

    "context"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo/options"
    "time"
)

func (service *Service) GetAvailableCages(checkinDate, checkoutDate time.Time, snake *model.Snake) ([]model.Cage, error) {
    minSize := snake.Length / 4

    filter := bson.M{
        "square_meters": bson.M{ "$gte": minSize},
        "bookings.check_in_date": bson.M{"$lte": checkinDate},
        "bookings.check_out_date": bson.M{"$gte": checkoutDate},
    }

    findOptions := options.Find()
    findOptions.SetSort(bson.D{{"price", 1}, {"square_meters", -1}})

    if snake.IsVenomous {
        filter["allow_dangerous_snakes"] = true
    }

    cursor, err := service.Collection("cages").Find(context.Background(), filter, findOptions)
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
        for _, b := range c.Bookings {
            if b.CheckInDate.Before(checkinDate) && b.CheckOutDate.After(checkoutDate) && b.GuestSnakeId.IsZero() {
                finalCages = append(finalCages, c)
            }
        }
    }

    return finalCages, nil
}
