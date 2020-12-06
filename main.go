package main

import (
    "tutorial/command"
    "tutorial/infrastructure"

    "context"
    "log"
     "go.mongodb.org/mongo-driver/mongo"
     "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    client, err := mongo.Connect(
        context.Background(),
        options.Client().ApplyURI("mongodb://localhost:27017"),
    )
    if err != nil {
        log.Fatalln(err)
    }
    defer func() {
        if err := client.Disconnect(context.Background()); err != nil {
            panic(err)
        }
    }()

    environment := infrastructure.NewEnvironment(client)

    command.Run(environment)
}
