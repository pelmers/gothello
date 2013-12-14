package gothello

import (
    "fmt"
    "math/rand"
    "time"
)

type RandomAI struct {
    display bool
}

// Construct a new Random AI controller.
// The disp parameter determines whether a board will be displayed during
// this of this controller's turn.
func NewRandomAI(disp bool) *RandomAI {
    rand.Seed(time.Now().UnixNano())
    return &RandomAI{disp}
}

// Given a board, return a random legal move.
// If there are no legal moves, return an empty Bitboard.
func (ai *RandomAI) Move(g *Game) Bitboard {
    moves := g.LegalMoves() // get a mask for this side's legal moves
    // allowed is a slice of the shifts that correspond to legal moves
    // e.g. with [0101100], allowed = [2,3,5]
    allowed := make([]Bitboard, 0, 64)
    for m := A1; ; m<<=1 {
        if m&moves != 0 {
            allowed = append(allowed, m)
        }
        if m == H8 {
            break
        }
    }
    var choice Bitboard
    if len(allowed) == 0 {
        // no allowed moves, so return no move
        choice = Bitboard(0)
    } else {
        // randomly pick a shift from the allowed list
        choice = allowed[rand.Intn(len(allowed))]
    }
    if ai.display {
        fmt.Printf("%s to play.\n%s\nBlack: %d\nWhite: %d\n",
            g.CurPlayerName(), g, g.Score(BLACK), g.Score(WHITE))
        fmt.Printf("My move is %s.\n", choice)
    }
    return choice
}
