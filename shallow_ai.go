package gothello

import "fmt"

// borrowed from Norvig's Paradigms of AI programming
var weights [64]int = [64]int{
    -120,-20,20,5,5,20,-20,120,-20,-40,-5,-5,-5,-5,-40,-20,
    20,-5,3,3,3,3,-5,20,5,-5,3,3,3,3,-5,5,
    -120,-20,20,5,5,20,-20,120,-20,-40,-5,-5,-5,-5,-40,-20,
    20,-5,3,3,3,3,-5,20,5,-5,3,3,3,3,-5,5,
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
        score += w*int((board>>uint(i))&1)
    }
    return score
}

func (ai *ShallowAI) Move(g *Game) Bitboard {
    legal := g.LegalMoves()
    oldmine, oldhis := g.Boards()
    move := Bitboard(0)
    topscore := 0
    // make every legal move, evaluate it, and pick the best one
    for m := A1; ; m <<= 1 {
        if legal & m != 0 {
            g.MakeMove(m)
            newmine, _ := g.Boards()
            score := ai.Evaluate(newmine)
            if score > topscore {
                topscore = score
                move = m
            }
            g.SetBoards(oldmine, oldhis)
        }
        // kind of miss a do-while loop
        if m == H8 {
            break
        }
    }
    if ai.display {
        fmt.Printf("%s to play.\n%s\nBlack: %d\nWhite: %d\n",
            g.CurPlayerName(), g, g.Score(BLACK), g.Score(WHITE))
        fmt.Printf("My move is %s.\n", move)
    }
    return move
}
