package gothello

import _ "fmt"

type Board struct {
    white, black Bitboard
    cp           int // current player
}

func InitBoard() *Board {
    return &Board{D4 | E5, D5 | E4, BLACK}
}

func (b *Board) CurPlayer() int {
    return b.cp
}

func (b *Board) NextPlayer() {
    b.cp ^= WHITE
}

func (b *Board) GetScore(side int) int {
    var pieces Bitboard
    if side == BLACK {
        pieces = b.black
    } else {
        pieces = b.white
    }
    return PopCount(pieces)
}

func flipDir(move, mine, his Bitboard, shifter func(Bitboard, uint) Bitboard,
    masker func(Bitboard) Bitboard) Bitboard {
    for c, flip := uint(1), Bitboard(0); c < 8; c++ {
        if shift := shifter(move, c) & masker(move); shift&his != 0 {
            flip |= shift & his
        } else {
            if shift&mine != 0 {
                return flip
            }
            break
        }
    }
    return Bitboard(0)
}

func (b *Board) doflip(move, flip Bitboard) {
    if b.CurPlayer() == BLACK {
        b.black |= flip | move
        b.white ^= flip
    } else {
        b.white |= flip | move
        b.black ^= flip
    }
}

func (b *Board) MakeMove(move Bitboard) {
    var mine, his Bitboard
    if b.CurPlayer() == BLACK {
        mine = b.black
        his = b.white
    } else {
        mine = b.white
        his = b.black
    }
    b.doflip(move, flipDir(move, mine, his, rshifter, rowmasker)|
        flipDir(move, mine, his, lshifter, rowmasker)|
        flipDir(move, mine, his, ushifter, colmasker)|
        flipDir(move, mine, his, dshifter, colmasker)|
        flipDir(move, mine, his, a45shifter, d45masker)|
        flipDir(move, mine, his, a225shifter, d45masker)|
        flipDir(move, mine, his, a135shifter, d135masker)|
        flipDir(move, mine, his, a315shifter, d135masker))
}

// Return string representation of the board.
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
