package main

import (
    "fmt"
    g "github.com/pelmers/gothello"
)

func main() {
    //game := g.InitGame(g.NewHumanController(), g.NewHumanController())
    //game := g.InitGame(g.NewShallowAI(true), g.NewHumanController())
    game := g.InitGame(g.NewShallowAI(true), g.NewRandomAI(true))
    for game.PlayTurn() {
    }
    fmt.Printf("%s\nBlack: %d\nWhite: %d\nGame over.\n", game,
        game.Score(g.BLACK), game.Score(g.WHITE))
}
