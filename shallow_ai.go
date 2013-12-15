package gothello

import "fmt"

type ShallowAI struct {
    display bool
}

func NewShallowAI(disp bool) *ShallowAI {
    return &ShallowAI{disp}
}

func (ai *ShallowAI) Move(g *Game) Bitboard {
    legal := g.LegalMoves()
    oldmine, oldhis := g.Boards()
    move := Bitboard(0)
    topscore := -6400
    // make every legal move, evaluate it, and pick the best one
    for m, i := A1, 0; i < 64; m, i = m<<1, i+1 {
        if legal&m != 0 {
            g.MakeMove(m)
            newmine, _ := g.Boards()
            score := Evaluate(newmine)
            if score > topscore {
                topscore = score
                move = m
            }
            g.SetBoards(oldmine, oldhis)
        }
    }
    if ai.display {
        fmt.Println(g)
        fmt.Printf("My move is %s.\n", move)
    }
    return move
}
