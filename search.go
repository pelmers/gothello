package gothello

import (
    "fmt"
    "goheap"
)

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

type OrderedMoves struct {
    Moves  []Bitboard
    Scores []int
}

func NewOrderedMoves(moves []Bitboard, scores []int) *OrderedMoves {
    return &OrderedMoves{moves, scores}
}

func (om *OrderedMoves) Swap(i, j int) {
    om.Moves[i], om.Moves[j] = om.Moves[j], om.Moves[i]
    om.Scores[i], om.Scores[j] = om.Scores[j], om.Scores[i]
}

func (om *OrderedMoves) Len() int {
    return len(om.Moves)
}

func (om *OrderedMoves) Less(i, j int) bool {
    return om.Scores[i] < om.Scores[j]
}

func (om *OrderedMoves) Pop() Bitboard {
    ret := om.Moves[len(om.Moves)-1]
    om.Moves = om.Moves[:len(om.Moves)-1]
    om.Scores = om.Scores[:len(om.Scores)-1]
    return ret
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
    for m, i := A1, 0; i < 64; m, i = m<<1, i+1 {
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
    for m, i := A1, 0; i < 64; m, i = m<<1, i+1 {
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
    /*
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
    */
    ordered := ai.OrderMoves(g)
    //for _, move := range ordered.Moves {
    for ordered.Len() > 0 {
        goheap.MinToEnd(ordered)
        move := ordered.Pop()
        g.MakeMove(move)
        g.NextPlayer()
        var score int
        if PopCount(mine|his) > 56 {
            score = ai.minimize(g, 9, bestscore, 999999)
        } else {
            score = ai.minimize(g, 3, bestscore, 999999)
        }
        if score > bestscore {
            bestscore = score
            bestmove = move
        }
        g.NextPlayer()
        g.SetBoards(mine, his)
    }
    if ai.display {
        fmt.Println(g)
        fmt.Printf("My move is %s.\n", bestmove)
    }
    return bestmove
}

func (ai *SearchAI) OrderMoves(g *Game) *OrderedMoves {
    mine, his := g.Boards()
    legal := g.LegalMoves()
    var moves []Bitboard
    var scores []int
    for m, i := A1, 0; i < 64; m, i = m<<1, i+1 {
        if legal&m != 0 {
            g.MakeMove(m)
            newmine, newhis := g.Boards()
            moves = append(moves, m)
            scores = append(scores, Evaluate(newmine)-Evaluate(newhis))
            g.SetBoards(mine, his)
        }
    }
    ordered := NewOrderedMoves(moves, scores)
    goheap.Heapify(ordered)
    return ordered
}
