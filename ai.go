package gothello

import (
    "fmt"
    "math/rand"
    "time"
)

type RandomAI struct {
    display bool
}

func NewRandomAI(disp bool) *RandomAI {
    rand.Seed(time.Now().UnixNano())
    return &RandomAI{disp}
}

func (ai *RandomAI) GetMove(b *Board) Bitboard {
    moves := b.GetLegalMoves()
    allowed := make([]uint, 0, 64)
    for i := uint(0); i < 64; i++ {
        if (Bitboard(1) << i) & moves != 0 {
            allowed = append(allowed, i)
        }
    }
    var choice Bitboard
    if len(allowed) == 0 {
        choice = Bitboard(0)
    } else {
        choice = Bitboard(1) << allowed[rand.Intn(len(allowed))]
    }
    if ai.display {
        fmt.Printf("%s to play.\n%s\nBlack: %d\nWhite: %d\n",
        b.CurPlayerName(), b, b.GetScore(BLACK), b.GetScore(WHITE))
        fmt.Printf("My move is %s.\n", choice)
    }
    return choice
}
