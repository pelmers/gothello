package gothello
import "fmt"

type ShallowAI struct {
    display bool
}

func NewShallowAI(disp bool) *ShallowAI {
    return &ShallowAI{disp}
}

func (ai *ShallowAI) Evaluate(board Bitboard) int {
    // naive evaluation: just count the number of pieces owned
    return PopCount(board)
}

func (ai *ShallowAI) Move(g *Game) Bitboard {
    legal := g.LegalMoves()
    oldmine, oldhis := g.Boards()
    move := Bitboard(0)
    topscore := 0
    for m := Bitboard(1); m != Bitboard(1)<<63 ; m <<= 1 {
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
    }
    if ai.display {
        fmt.Printf("%s to play.\n%s\nBlack: %d\nWhite: %d\n",
            g.CurPlayerName(), g, g.Score(BLACK), g.Score(WHITE))
        fmt.Printf("My move is %s.\n", move)
    }
    return move
}
