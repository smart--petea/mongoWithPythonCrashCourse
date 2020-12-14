package command

import (
    "tutorial/infrastructure"
    "tutorial/print"
)

func ListCages(env *infrastructure.Environment, args []string) {
    print.HeadLine("Your cages")

    if env.State.ActiveAccount == nil {
        print.Error("You must login first to see the user's list of cages")
        return
    }

    cages, err := env.Dataservice.FindCagesForUser(env.State.ActiveAccount) 
    if err != nil {
        print.Error(err.Error())
        return 
    }

    print.Success("You have %d cages", len(cages))

    for i := range cages {
        print.Success(" * %s is %f meters", cages[i].Name, cages[i].SquareMeters)
    }
}
