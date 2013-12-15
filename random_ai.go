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
    // allowed is a slice of the legal moves
    allowed := make([]Bitboard, 0, 64)
    for m, i := A1, 0; i < 64; m, i = m<<1, i+1 {
        if m&moves != 0 {
            allowed = append(allowed, m)
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
        fmt.Println(g)
        fmt.Printf("My move is %s.\n", choice)
    }
    return choice
}
