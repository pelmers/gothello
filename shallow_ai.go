package gothello

import "fmt"

// borrowed from Norvig's Paradigms of AI programming
var weights [64]int = [64]int{
    -120, -20, 20, 5, 5, 20, -20, 120, -20, -40, -5, -5, -5, -5, -40, -20,
    20, -5, 3, 3, 3, 3, -5, 20, 5, -5, 3, 3, 3, 3, -5, 5,
    -120, -20, 20, 5, 5, 20, -20, 120, -20, -40, -5, -5, -5, -5, -40, -20,
    20, -5, 3, 3, 3, 3, -5, 20, 5, -5, 3, 3, 3, 3, -5, 5,
}

type ShallowAI struct {
    display bool
}

func NewShallowAI(disp bool) *ShallowAI {
    return &ShallowAI{disp}
}

func (ai *ShallowAI) Evaluate(board Bitboard) int {
    score := 0
    for i, w := range weights {
        score += w * int((board>>uint(i))&1)
    }
    return score
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
            score := ai.Evaluate(newmine)
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
