package gothello

import "fmt"

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a,b int) int {
    if a < b {
        return a
    }
    return b
}

type SearchAI struct {
    display bool
}

func NewSearchAI(disp bool) *SearchAI {
    return &SearchAI{disp}
}

func (ai *SearchAI) maximize(g *Game, depth, alpha, beta int) int {
    mine, his := g.Boards()
    if g.IsEnd() || depth == 0 {
        return Evaluate(mine) - Evaluate(his)
    }
    legal := g.LegalMoves()
    for m, i := A1, 0; i < 64; m, i = m <<1, i + 1 {
        if legal&m != 0 {
            g.MakeMove(m)
            g.NextPlayer()
            alpha = max(ai.minimize(g, depth-1, alpha, beta), alpha)
            g.NextPlayer()
            g.SetBoards(mine, his)
            if beta <= alpha {
                break
            }
        }
    }
    return alpha
}

func (ai *SearchAI) minimize(g *Game, depth, alpha, beta int) int {
    mine, his := g.Boards()
    if g.IsEnd() || depth == 0 {
        return Evaluate(mine) - Evaluate(his)
    }
    legal := g.LegalMoves()
    for m, i := A1, 0; i < 64; m, i = m <<1, i + 1 {
        if legal&m != 0 {
            g.MakeMove(m)
            g.NextPlayer()
            beta = min(ai.maximize(g, depth-1, alpha, beta), beta)
            g.NextPlayer()
            g.SetBoards(mine, his)
            if beta <= alpha {
                break
            }
        }
    }
    return beta
}

func (ai *SearchAI) Move(g *Game) Bitboard {
    bestmove := Bitboard(0)
    bestscore := -999999
    mine, his := g.Boards()
    legal := g.LegalMoves()
    for m, i := A1, 0; i < 64; m, i = m << 1, i+1 {
        if legal&m != 0 {
            g.MakeMove(m)
            g.NextPlayer()
            var score int
            if PopCount(mine|his) > 56 {
                score = ai.minimize(g, 9, bestscore, 999999)
            } else {
                score = ai.minimize(g, 3, bestscore, 999999)
            }
            if score > bestscore {
                bestscore = score
                bestmove = m
            }
            g.NextPlayer()
            g.SetBoards(mine, his)
        }
    }
    if ai.display {
        fmt.Println(g)
        fmt.Printf("My move is %s.\n", bestmove)
    }
    return bestmove
}
