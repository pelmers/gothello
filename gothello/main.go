package main

import (
    "flag"
    "fmt"
    g "github.com/pelmers/gothello"
)

func main() {
    black := flag.String("black", "human", "Black controller: human, shallow, search, random")
    white := flag.String("white", "human", "White controller: human, shallow, search, random")
    show := flag.Bool("s", true, "show AI moves")
    flag.Parse()
    var bc, wc g.Controller
    switch *black {
    case "shallow":
        bc = g.NewShallowAI(*show)
    case "random":
        bc = g.NewRandomAI(*show)
    case "search":
        bc = g.NewSearchAI(*show)
    default:
        bc = g.NewHumanController()
    }
    switch *white {
    case "shallow":
        wc = g.NewShallowAI(*show)
    case "random":
        wc = g.NewRandomAI(*show)
    case "search":
        wc = g.NewSearchAI(*show)
    default:
        wc = g.NewHumanController()
    }
    game := g.InitGame(bc, wc)
    for game.PlayTurn() {
    }
    fmt.Printf("%s%s", game, "\nGame over.\n")
}
