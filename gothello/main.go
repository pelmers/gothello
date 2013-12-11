package main

import (
    g "github.com/pelmers/gothello"
    "fmt"
)

func main() {
    //game := g.InitBoard(g.NewHumanController(), g.NewHumanController())
    game := g.InitBoard(g.NewRandomAI(true), g.NewRandomAI(true))
    for game.PlayTurn() {
    }
    fmt.Printf("%s\nWhite: %d\nBlack: %d\n",game,
    game.GetScore(g.BLACK), game.GetScore(g.WHITE))
}
