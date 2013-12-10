package gothello

import _ "fmt"

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

// Board struct maintains the state of the game, handling operations related
// to the board itself (making flips, checking legality, etc.)
type Board struct {
    white, black Bitboard
    cp, unplayed int // current player and number of turns skipped
    bc, wc       Controller
}

// Initialize a Board with given Controllers.
func InitBoard(black, white Controller) *Board {
    return &Board{D4 | E5, D5 | E4, BLACK, 0, black, white}
}

// Return the current player.
func (b *Board) CurPlayer() int {
    return b.cp
}

// Change the current player to the other player.
func (b *Board) NextPlayer() {
    b.cp ^= WHITE
}

// Return the current side's controller.
func (b *Board) curController() Controller {
    if b.cp == BLACK {
        return b.bc
    }
    return b.wc
}

// Return the two side's boards, current side first.
func (b *Board) getPlayerBoards() (Bitboard, Bitboard) {
    if b.CurPlayer() == BLACK {
        return b.black, b.white
    }
    return b.white, b.black
}

// Return the number of pieces controlled by given side.
func (b *Board) GetScore(side int) int {
    board, _ := b.getPlayerBoards()
    return PopCount(board)
}

// Flip the pieces in the flip mask to the current player's side.
func (b *Board) doflip(move, flip Bitboard) {
    if b.CurPlayer() == BLACK {
        b.black |= flip | move
        b.white ^= flip
    } else {
        b.white |= flip | move
        b.black ^= flip
    }
}

// Make the move on the board and flip opponent pieces.
// If move == 0, increment the unplayed turn counter.
func (b *Board) MakeMove(move Bitboard) {
    if move == Bitboard(0) {
        b.unplayed++
    } else {
        b.unplayed = 0
        mine, his := b.getPlayerBoards()
        b.doflip(move, flipDir(move, mine, his, rshifter, rowmasker)|
            flipDir(move, mine, his, lshifter, rowmasker)|
            flipDir(move, mine, his, ushifter, colmasker)|
            flipDir(move, mine, his, dshifter, colmasker)|
            flipDir(move, mine, his, a45shifter, d45masker)|
            flipDir(move, mine, his, a225shifter, d45masker)|
            flipDir(move, mine, his, a135shifter, d135masker)|
            flipDir(move, mine, his, a315shifter, d135masker))
    }
}

// Return whether move is legal.
func (b *Board) IsLegalMove(move Bitboard) bool {
    mine, his := b.getPlayerBoards()
    return flipDir(move, mine, his, rshifter, rowmasker) != 0 ||
        flipDir(move, mine, his, lshifter, rowmasker) != 0 ||
        flipDir(move, mine, his, ushifter, colmasker) != 0 ||
        flipDir(move, mine, his, dshifter, colmasker) != 0 ||
        flipDir(move, mine, his, a45shifter, d45masker) != 0 ||
        flipDir(move, mine, his, a225shifter, d45masker) != 0 ||
        flipDir(move, mine, his, a135shifter, d135masker) != 0 ||
        flipDir(move, mine, his, a315shifter, d135masker) != 0
}

// Return whether the game has ended.
func (b *Board) IsEnd() bool {
    // if we've gone two turns without playing or board is full, it's over
    return b.unplayed == 2 || b.white|b.black == ^Bitboard(0)
}

// Call upon the current player's controller to make a move.
// Return whether the game continues after the move.
func (b *Board) PlayTurn() bool {
    b.MakeMove(b.curController().GetMove(b))
    b.NextPlayer()
    return b.IsEnd()
}

// Return string representation of the board with labeled rows and columns.
func (b *Board) String() string {
    repr := "\t1   2   3   4   5   6   7   8\n"
    for r, mask := range rows {
        repr += string(65+r) + "\t"
        row_white := (b.white & mask) >> uint(8*r)
        row_black := (b.black & mask) >> uint(8*r)
        for c := uint(0); c < 8; c++ {
            if row_white&(1<<c) != 0 {
                repr += "W   "
            } else if row_black&(1<<c) != 0 {
                repr += "B   "
            } else {
                repr += "    "
            }
        }
        repr += "\n------------------------------------------\n"
    }
    return repr[:len(repr)-43]
}
