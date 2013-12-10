package main

import g "gothello"

func main() {
    game := g.InitBoard(g.NewHumanController(), g.NewHumanController())
    for game.PlayTurn() {
    }
}
