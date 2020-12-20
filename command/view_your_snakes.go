package command

import (
    "tutorial/infrastructure"
    "tutorial/print"
)

func ViewYourSnakes(env *infrastructure.Environment, args []string) {
    snakes, err := env.Dataservice.FindSnakesForUser(env.State.ActiveAccount) 
    if err != nil {
        print.Error(err.Error())
        return 
    }

    print.Success("You have %d snakes", len(snakes))

    for i, snake := range snakes {
            print.Success(snake.ToLoopStringLine(i))
    }
}
