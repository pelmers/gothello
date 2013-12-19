package main

import (
    "flag"
    "fmt"
    g "github.com/pelmers/gothello"
)

const WIDTH = 80

func progress_bar(width, percent int) {
    width -= 9
    var i int
    filled := width * percent / 100
    fmt.Print("\r[ ")
    for i = 0; i < filled; i++ {
        fmt.Print("#")
    }
    for ; i < width; i++ {
        fmt.Print("-")
    }
    fmt.Printf(" ] %d%%", percent)
}

func ParseControllers(b, w *string, show *bool) (g.Controller, g.Controller) {
    var bc, wc g.Controller
    switch *b {
    case "shallow":
        bc = g.NewShallowAI(*show)
    case "random":
        bc = g.NewRandomAI(*show)
    case "search":
        bc = g.NewSearchAI(*show)
    default:
        bc = g.NewHumanController()
    }
    switch *w {
    case "shallow":
        wc = g.NewShallowAI(*show)
    case "random":
        wc = g.NewRandomAI(*show)
    case "search":
        wc = g.NewSearchAI(*show)
    default:
        wc = g.NewHumanController()
    }
    return bc, wc
}

func main() {
    black := flag.String("black", "human", "Black controller: human, shallow, search, random")
    white := flag.String("white", "human", "White controller: human, shallow, search, random")
    show := flag.Bool("s", false, "Show AI moves")
    randomize := flag.Bool("r", false, "Randomize first 5 moves")
    simulate := flag.Int("sim", 0, "Number of AI vs. AI games to simulate")
    flag.Parse()
    bc, wc := ParseControllers(black, white, show)
    if *simulate <= 0 {
        game := g.InitGame(bc, wc)
        if *randomize {
            game.SetControllers(g.NewRandomAI(*show), g.NewRandomAI(*show))
        }
        for game.PlayTurn() {
            if *randomize {
                if game.Score(g.BLACK)+game.Score(g.WHITE) >= 10 {
                    game.SetControllers(bc, wc)
                    *randomize = false // stop randomizing
                }
            }
        }
        fmt.Printf("%s%s", game, "\nGame over.\n")
    } else {
        wins_b, wins_w, draws := 0, 0, 0
        percent, prev_percent := 0, 0
        for j := 0; j < *simulate; j++ {
            percent = j * 100 / (*simulate)
            if percent > prev_percent {
                progress_bar(WIDTH, percent)
                prev_percent = percent
            }
            game := g.InitGame(bc, wc)
            if *randomize {
                game.SetControllers(g.NewRandomAI(*show), g.NewRandomAI(*show))
            }
            for game.PlayTurn() {
                if *randomize {
                    if game.Score(g.BLACK)+game.Score(g.WHITE) >= 10 {
                        game.SetControllers(bc, wc)
                        *randomize = false // stop randomizing
                    }
                }
            }
            score_b := game.Score(g.BLACK)
            score_w := game.Score(g.WHITE)
            if score_b > score_w {
                wins_b++
            } else if score_b == score_w {
                draws++
            } else {
                wins_w++
            }
        }
        fmt.Printf("\nBlack: %.2f%%\nWhite: %.2f%%\nDraw: %.2f%%\n",
            float64(wins_b)/float64(*simulate)*100.0,
            float64(wins_w)/float64(*simulate)*100.0,
            float64(draws)/float64(*simulate)*100.0)
    }
}
