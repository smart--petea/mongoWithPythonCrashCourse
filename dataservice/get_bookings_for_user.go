package dataservice

import (
    "tutorial/model"
    "fmt"

    "context"
    "go.mongodb.org/mongo-driver/bson"
)

func (service *Service) GetBookingsForUser(user *model.Owner) ([]*model.Booking, error) {

    filter := bson.M{
        "bookings.guest_owner_id": user.ID,
    }

    cursor, err := service.Collection("cages").Find(context.Background(), filter)
    if err != nil {
        return nil, err
    }

    var cages []model.Cage
    err = cursor.All(context.Background(), &cages)
    if err != nil {
        return nil, err
    }
    fmt.Printf("leng(cages) = %d\n", len(cages))

    var bookings []*model.Booking = []*model.Booking{}
    for _, c := range cages {
        fmt.Printf("len(bookings) = %d\n", len(c.Bookings))
        for i := range c.Bookings {
            fmt.Println(c.Bookings[i].GuestOwnerID, user.ID)
            if c.Bookings[i].GuestOwnerID == user.ID {
                bookings = append(bookings, &c.Bookings[i])
            }
        }
    }

    return bookings, nil
}
