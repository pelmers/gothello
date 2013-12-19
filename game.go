package gothello

import "fmt"

// Return a Bitboard mask representing the pieces that move flips
// in the direction determined by a shifter and a masker.
func flipDir(move, mine, his Bitboard, shift shifter, mask masker) Bitboard {
    for c, flip := uint(1), Bitboard(0); c < 8; c++ {
        if s := shift(move, c) & mask(move); s&his != 0 {
            flip |= s & his
        } else {
            if s&mine != 0 {
                return flip
            }
            break
        }
    }
    return Bitboard(0)
}

// Return a Bitboard mask representing the valid plays of the current side,
// in the direction defined by wrapmask and shift.
func findValid(wrapmask, mine, his Bitboard, shift shifter) Bitboard {
    moves := Bitboard(0)
    for p := shift(mine, 1) & his; p != 0; p &= his {
        p = shift(p&wrapmask, 1)
        moves |= p
    }
    return moves & ^(mine | his)
}

// Game struct maintains the state of the game, handling operations related
// to the board itself (making flips, checking legality, etc.)
type Game struct {
    white, black Bitboard
    cp, unplayed int // current player and number of turns skipped
    bc, wc       Controller
}

// Initialize a Game with given Controllers.
func InitGame(black, white Controller) *Game {
    return &Game{D4 | E5, D5 | E4, BLACK, 0, black, white}
}

// Set the current and other player's boards.
func (g *Game) SetBoards(current, other Bitboard) {
    if g.CurPlayer() == BLACK {
        g.black, g.white = current, other
    } else {
        g.white, g.black = current, other
    }
}

// Set black and white's controllers.
func (g *Game) SetControllers(black, white Controller) {
    g.bc, g.wc = black, white
}

// Return the two side's boards, current side first.
func (g *Game) Boards() (Bitboard, Bitboard) {
    if g.CurPlayer() == BLACK {
        return g.black, g.white
    }
    return g.white, g.black
}

// Return the current player.
func (g *Game) CurPlayer() int {
    return g.cp
}

// Return a string representation of the current player.
func (g *Game) CurPlayerName() string {
    if g.cp == WHITE {
        return "White"
    }
    return "Black"
}

// Change the current player to the other player.
func (g *Game) NextPlayer() {
    g.cp ^= WHITE
}

// Return the current side's controller.
func (g *Game) curController() Controller {
    if g.cp == BLACK {
        return g.bc
    }
    return g.wc
}

// Return the number of pieces controlled by given side.
func (g *Game) Score(side int) int {
    board := g.black
    if side == WHITE {
        board = g.white
    }
    return PopCount(board)
}

// Flip the pieces in the flip mask to the current player's side.
func (g *Game) doflip(move, flip Bitboard) {
    if g.CurPlayer() == BLACK {
        g.black |= flip | move
        g.white ^= flip
    } else {
        g.white |= flip | move
        g.black ^= flip
    }
}

// Make the move on the board and flip opponent pieces.
// If move == 0, increment the unplayed turn counter.
// Note: does not change the current player.
func (g *Game) MakeMove(move Bitboard) {
    if move == Bitboard(0) {
        g.unplayed++
    } else {
        g.unplayed = 0
        mine, his := g.Boards()
        g.doflip(move, flipDir(move, mine, his, rshifter, rowmasker)|
            flipDir(move, mine, his, lshifter, rowmasker)|
            flipDir(move, mine, his, ushifter, colmasker)|
            flipDir(move, mine, his, dshifter, colmasker)|
            flipDir(move, mine, his, a45shifter, d45masker)|
            flipDir(move, mine, his, a225shifter, d45masker)|
            flipDir(move, mine, his, a135shifter, d135masker)|
            flipDir(move, mine, his, a315shifter, d135masker))
    }
}

// Return a mask of all of the legal moves for the current player.
func (g *Game) LegalMoves() Bitboard {
    mine, his := g.Boards()
    return findValid(WRAPUP, mine, his, ushifter) |
        findValid(WRAPDN, mine, his, dshifter) |
        findValid(WRAPRGT, mine, his, rshifter) |
        findValid(WRAPLFT, mine, his, lshifter) |
        findValid(WRAP45, mine, his, a45shifter) |
        findValid(WRAP225, mine, his, a225shifter) |
        findValid(WRAP135, mine, his, a135shifter) |
        findValid(WRAP315, mine, his, a315shifter)
}

// Return whether the game has ended.
func (g *Game) IsEnd() bool {
    // if we've gone two turns without playing or board is full, it's over
    return g.unplayed == 2 || g.white|g.black == ^Bitboard(0)
}

// Call upon the current player's controller to make a move.
// Return whether the game continues after the move.
func (g *Game) PlayTurn() bool {
    g.MakeMove(g.curController().Move(g))
    g.NextPlayer()
    return !g.IsEnd()
}

// Return string representation of the board with labeled rows and columns.
func (g *Game) String() string {
    repr := "\t1   2   3   4   5   6   7   8\n"
    legal := g.LegalMoves()
    for r, mask := range rows {
        repr += string(65+r) + "\t"
        row_white := (g.white & mask) >> uint(8*r)
        row_black := (g.black & mask) >> uint(8*r)
        row_legal := (legal & mask) >> uint(8*r)
        for c := uint(0); c < 8; c++ {
            if row_white&(1<<c) != 0 {
                repr += "W   "
            } else if row_black&(1<<c) != 0 {
                repr += "B   "
            } else if row_legal&(1<<c) != 0 {
                repr += "-   "
            } else {
                repr += "    "
            }
        }
        repr += "\n------------------------------------------\n"
    }
    return fmt.Sprintf("%s to play.\n%s\nBlack: %d\nWhite: %d",
        g.CurPlayerName(), repr[:len(repr)-43], g.Score(BLACK), g.Score(WHITE))
}
