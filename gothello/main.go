package main

import g "github.com/pelmers/gothello"

func main() {
    game := g.InitBoard(g.NewHumanController(), g.NewHumanController())
    for game.PlayTurn() {
    }
}
