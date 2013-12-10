package gothello

import (
    "fmt"
    "strings"
)

type Bitboard uint64 // 8x8 is 64 bits

const (
    BLACK = iota // 0
    WHITE        // 1
)

const (
    A1  Bitboard = 1 << iota
    A2  Bitboard = 1 << iota
    A3  Bitboard = 1 << iota
    A4  Bitboard = 1 << iota
    A5  Bitboard = 1 << iota
    A6  Bitboard = 1 << iota
    A7  Bitboard = 1 << iota
    A8  Bitboard = 1 << iota
    B1  Bitboard = 1 << iota
    B2  Bitboard = 1 << iota
    B3  Bitboard = 1 << iota
    B4  Bitboard = 1 << iota
    B5  Bitboard = 1 << iota
    B6  Bitboard = 1 << iota
    B7  Bitboard = 1 << iota
    B8  Bitboard = 1 << iota
    C1  Bitboard = 1 << iota
    C2  Bitboard = 1 << iota
    C3  Bitboard = 1 << iota
    C4  Bitboard = 1 << iota
    C5  Bitboard = 1 << iota
    C6  Bitboard = 1 << iota
    C7  Bitboard = 1 << iota
    C8  Bitboard = 1 << iota
    D1  Bitboard = 1 << iota
    D2  Bitboard = 1 << iota
    D3  Bitboard = 1 << iota
    D4  Bitboard = 1 << iota
    D5  Bitboard = 1 << iota
    D6  Bitboard = 1 << iota
    D7  Bitboard = 1 << iota
    D8  Bitboard = 1 << iota
    E1  Bitboard = 1 << iota
    E2  Bitboard = 1 << iota
    E3  Bitboard = 1 << iota
    E4  Bitboard = 1 << iota
    E5  Bitboard = 1 << iota
    E6  Bitboard = 1 << iota
    E7  Bitboard = 1 << iota
    E8  Bitboard = 1 << iota
    F1  Bitboard = 1 << iota
    F2  Bitboard = 1 << iota
    F3  Bitboard = 1 << iota
    F4  Bitboard = 1 << iota
    F5  Bitboard = 1 << iota
    F6  Bitboard = 1 << iota
    F7  Bitboard = 1 << iota
    F8  Bitboard = 1 << iota
    G1  Bitboard = 1 << iota
    G2  Bitboard = 1 << iota
    G3  Bitboard = 1 << iota
    G4  Bitboard = 1 << iota
    G5  Bitboard = 1 << iota
    G6  Bitboard = 1 << iota
    G7  Bitboard = 1 << iota
    G8  Bitboard = 1 << iota
    H1  Bitboard = 1 << iota
    H2  Bitboard = 1 << iota
    H3  Bitboard = 1 << iota
    H4  Bitboard = 1 << iota
    H5  Bitboard = 1 << iota
    H6  Bitboard = 1 << iota
    H7  Bitboard = 1 << iota
    H8  Bitboard = 1 << iota
)

var rows [8]Bitboard = [8]Bitboard{
    A1 | A2 | A3 | A4 | A5 | A6 | A7 | A8,
    B1 | B2 | B3 | B4 | B5 | B6 | B7 | B8,
    C1 | C2 | C3 | C4 | C5 | C6 | C7 | C8,
    D1 | D2 | D3 | D4 | D5 | D6 | D7 | D8,
    E1 | E2 | E3 | E4 | E5 | E6 | E7 | E8,
    F1 | F2 | F3 | F4 | F5 | F6 | F7 | F8,
    G1 | G2 | G3 | G4 | G5 | G6 | G7 | G8,
    H1 | H2 | H3 | H4 | H5 | H6 | H7 | H8,
}

var cols [8]Bitboard = [8]Bitboard{
    A1 | B1 | C1 | D1 | E1 | F1 | G1 | H1,
    A2 | B2 | C2 | D2 | E2 | F2 | G2 | H2,
    A3 | B3 | C3 | D3 | E3 | F3 | G3 | H3,
    A4 | B4 | C4 | D4 | E4 | F4 | G4 | H4,
    A5 | B5 | C5 | D5 | E5 | F5 | G5 | H5,
    A6 | B6 | C6 | D6 | E6 | F6 | G6 | H6,
    A7 | B7 | C7 | D7 | E7 | F7 | G7 | H7,
    A8 | B8 | C8 | D8 | E8 | F8 | G8 | H8,
}

var diag45 [15]Bitboard = [15]Bitboard{
    A1,
    B1 | A2,
    C1 | B2 | A3,
    D1 | C2 | B3 | A4,
    E1 | D2 | C3 | B4 | A5,
    F1 | E2 | D3 | C4 | B5 | A6,
    G1 | F2 | E3 | D4 | C5 | B6 | A7,
    H1 | G2 | F3 | E4 | D5 | C6 | B7 | A8,
    H2 | G3 | F4 | E5 | D6 | C7 | B8,
    H3 | G4 | F5 | E6 | D7 | C8,
    H4 | G5 | F6 | E7 | D8,
    H5 | G6 | F7 | E8,
    H6 | G7 | F8,
    H7 | G8,
    H8,
}

var diag135 [15]Bitboard = [15]Bitboard{
    A8,
    A7 | B8,
    A6 | B7 | C6,
    A5 | B6 | C7 | D8,
    A4 | B5 | C6 | D7 | E8,
    A3 | B4 | C5 | D6 | E7 | F8,
    A2 | B3 | C4 | D5 | E6 | F7 | G8,
    A1 | B2 | C3 | D4 | E5 | F6 | G7 | H8,
    B1 | C2 | D3 | E4 | F5 | G6 | H7,
    C1 | D2 | E3 | F4 | G5 | H6,
    D1 | E2 | F3 | G4 | H5,
    E1 | F2 | G3 | H4,
    F1 | G2 | H3,
    G1 | H2,
    H1,
}

// map each square to its row mask
var rowmap map[Bitboard]uint = map[Bitboard]uint{
    A1: 0, A2: 0, A3: 0, A4: 0, A5: 0, A6: 0, A7: 0, A8: 0,
    B1: 1, B2: 1, B3: 1, B4: 1, B5: 1, B6: 1, B7: 1, B8: 1,
    C1: 2, C2: 2, C3: 2, C4: 2, C5: 2, C6: 2, C7: 2, C8: 2,
    D1: 3, D2: 3, D3: 3, D4: 3, D5: 3, D6: 3, D7: 3, D8: 3,
    E1: 4, E2: 4, E3: 4, E4: 4, E5: 4, E6: 4, E7: 4, E8: 4,
    F1: 5, F2: 5, F3: 5, F4: 5, F5: 5, F6: 5, F7: 5, F8: 5,
    G1: 6, G2: 6, G3: 6, G4: 6, G5: 6, G6: 6, G7: 6, G8: 6,
    H1: 7, H2: 7, H3: 7, H4: 7, H5: 7, H6: 7, H7: 7, H8: 7,
}

// map each square to its column mask
var colmap map[Bitboard]uint = map[Bitboard]uint{
    A1: 0, A2: 1, A3: 2, A4: 3, A5: 4, A6: 5, A7: 6, A8: 7,
    B1: 0, B2: 1, B3: 2, B4: 3, B5: 4, B6: 5, B7: 6, B8: 7,
    C1: 0, C2: 1, C3: 2, C4: 3, C5: 4, C6: 5, C7: 6, C8: 7,
    D1: 0, D2: 1, D3: 2, D4: 3, D5: 4, D6: 5, D7: 6, D8: 7,
    E1: 0, E2: 1, E3: 2, E4: 3, E5: 4, E6: 5, E7: 6, E8: 7,
    F1: 0, F2: 1, F3: 2, F4: 3, F5: 4, F6: 5, F7: 6, F8: 7,
    G1: 0, G2: 1, G3: 2, G4: 3, G5: 4, G6: 5, G7: 6, G8: 7,
    H1: 0, H2: 1, H3: 2, H4: 3, H5: 4, H6: 5, H7: 6, H8: 7,
}

var diag45map map[Bitboard]uint = map[Bitboard]uint{
    A1: 0, A2: 1, A3: 2, A4: 3, A5: 4, A6: 5, A7: 6, A8: 7,
    B1: 1, B2: 2, B3: 3, B4: 4, B5: 5, B6: 6, B7: 7, B8: 8,
    C1: 2, C2: 3, C3: 4, C4: 5, C5: 6, C6: 7, C7: 8, C8: 9,
    D1: 3, D2: 4, D3: 5, D4: 6, D5: 7, D6: 8, D7: 9, D8: 10,
    E1: 4, E2: 5, E3: 6, E4: 7, E5: 8, E6: 9, E7: 10, E8: 11,
    F1: 5, F2: 6, F3: 7, F4: 8, F5: 9, F6: 10, F7: 11, F8: 12,
    G1: 6, G2: 7, G3: 8, G4: 9, G5: 10, G6: 11, G7: 12, G8: 13,
    H1: 7, H2: 8, H3: 9, H4: 10, H5: 11, H6: 12, H7: 13, H8: 14,
}

var diag135map map[Bitboard]uint = map[Bitboard]uint{
    A1: 7, A2: 6, A3: 5, A4: 4, A5: 3, A6: 2, A7: 1, A8: 0,
    B1: 8, B2: 7, B3: 6, B4: 5, B5: 4, B6: 3, B7: 2, B8: 1,
    C1: 9, C2: 8, C3: 7, C4: 6, C5: 5, C6: 4, C7: 3, C8: 2,
    D1: 10, D2: 9, D3: 8, D4: 7, D5: 6, D6: 5, D7: 4, D8: 3,
    E1: 11, E2: 10, E3: 9, E4: 8, E5: 7, E6: 6, E7: 5, E8: 4,
    F1: 12, F2: 11, F3: 10, F4: 9, F5: 8, F6: 7, F7: 6, F8: 5,
    G1: 13, G2: 12, G3: 11, G4: 10, G5: 9, G6: 8, G7: 7, G8: 6,
    H1: 14, H2: 13, H3: 12, H4: 11, H5: 10, H6: 9, H7: 8, H8: 7,
}

type masker func(Bitboard) Bitboard
type shifter func(Bitboard, uint) Bitboard

var rowmasker masker = func(m Bitboard) Bitboard { return rows[rowmap[m]] }
var colmasker masker = func(m Bitboard) Bitboard { return cols[colmap[m]] }
var d45masker masker = func(m Bitboard) Bitboard { return diag45[diag45map[m]] }
var d135masker masker = func(m Bitboard) Bitboard { return diag135[diag135map[m]] }
var rshifter shifter = func(m Bitboard, c uint) Bitboard { return m >> c }
var lshifter shifter = func(m Bitboard, c uint) Bitboard { return m << c }
var ushifter shifter = func(m Bitboard, c uint) Bitboard { return m << (8 * c) }
var dshifter shifter = func(m Bitboard, c uint) Bitboard { return m >> (8 * c) }
var a45shifter shifter = func(m Bitboard, c uint) Bitboard { return m >> (7 * c) }
var a225shifter shifter = func(m Bitboard, c uint) Bitboard { return m << (7 * c) }
var a135shifter shifter = func(m Bitboard, c uint) Bitboard { return m >> (9 * c) }
var a315shifter shifter = func(m Bitboard, c uint) Bitboard { return m << (9 * c) }

func RC2Mask(row, col uint) Bitboard {
    return Bitboard(1 << (row*8 + col))
}

func Mask2RC(mask Bitboard) (uint, uint) {
    if row, exists := rowmap[mask]; exists {
        return row, colmap[mask]
    } else {
        return 8, 8
    }
}

func PopCount(board Bitboard) int {
    var c int
    for c = 0; board != 0; c++ {
        board &= board - 1
    }
    return c
}

func (b Bitboard) String() string {
    str := fmt.Sprintf("%b", b)
    str = strings.Repeat("0", 64-len(str)) + str
    for i := 0; i < 7; i++ {
        str = str[:i*8+i+8] + "\n" + str[i*8+i+8:]
    }
    return str
}
