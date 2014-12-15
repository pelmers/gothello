package main

import (
	"flag"
	"fmt"
	g "github.com/pelmers/gothello"
	"runtime"
)

// Struct that encapsulates the final score of black and white.
type Score struct {
	black, white int
}

// Width of the console for progress bar
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

/// Play a game to the end.
/// If randomize is true, the first 10 turns are randomized.
func playFullGame(game *g.Game, randomize, show bool, black, white g.Controller) {
	if randomize {
		game.SetControllers(g.NewRandomAI(show), g.NewRandomAI(show))
	}
	for game.PlayTurn() {
		if randomize {
			// stop using the random controller after the first 10 turns
			if game.Score(g.BLACK)+game.Score(g.WHITE) >= 10 {
				game.SetControllers(black, white)
			}
		}
	}
}

/// Perform num simulations with bc and wc controllers.
/// If randomize, each game's first 10 turns are played randomly.
/// Return values (black wins, white wins, draws), where their sum = num
func performSimulations(num int, randomize bool, bc, wc g.Controller) (int, int, int) {
	wins_b, wins_w, draws := 0, 0, 0
	score_channel := make(chan Score, num)
	for j := 0; j < num; j++ {
		game := g.InitGame(bc, wc)
		// play the full game and send the score back
		go func() {
			playFullGame(game, randomize, false, bc, wc)
			score_channel <- Score{game.Score(g.BLACK), game.Score(g.WHITE)}
		}()
	}
	percent, prev_percent := 0, 0
	for j := 0; j < num; j++ {
		score := <-score_channel
		if score.black > score.white {
			wins_b++
		} else if score.black == score.white {
			draws++
		} else {
			wins_w++
		}
		percent = (wins_b + wins_w + draws) * 100 / num
		if percent > prev_percent {
			progress_bar(WIDTH, percent)
			prev_percent = percent
		}
	}
	return wins_b, wins_w, draws
}

func main() {
	black := flag.String("black", "human", "Black controller: human, shallow, search, random")
	white := flag.String("white", "human", "White controller: human, shallow, search, random")
	show := flag.Bool("s", false, "Show AI moves")
	randomize := flag.Bool("r", false, "Randomize first 5 moves of each side")
	simulate := flag.Int("sim", 0, "Number of AI vs. AI games to simulate")
	flag.Parse()
	bc, wc := ParseControllers(black, white, show)
	if *simulate <= 0 {
		// just play out the game regularly
		game := g.InitGame(bc, wc)
		playFullGame(game, *randomize, *show, bc, wc)
		fmt.Printf("%s%s", game, "\nGame over.\n")
	} else {
		runtime.GOMAXPROCS(runtime.NumCPU())
		fmt.Printf("Setting GOMAXPROCS to %d\n", runtime.NumCPU())
		wins_b, wins_w, draws := performSimulations(*simulate, *randomize, bc, wc)
		fmt.Printf("\nBlack: %.2f%%\nWhite: %.2f%%\nDraw: %.2f%%\n",
			float64(wins_b)/float64(*simulate)*100.0,
			float64(wins_w)/float64(*simulate)*100.0,
			float64(draws)/float64(*simulate)*100.0)
	}
}
