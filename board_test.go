package gothello

import "testing"

func TestBoard(t *testing.T) {
    b := InitBoard(NewNullControl(), NewNullControl())
    b.String()
    if !b.PlayTurn() {
        t.Errorf("Incorrectly reporting game has ended on first move.")
    }
    // populate white with all of the empty squares
    b.white |= ^b.black
    if !b.IsEnd() {
        t.Errorf("Incorrectly reporting game continues with full board.")
    }
}

func testmove(black, white, move, bresult, wresult Bitboard, cp int) bool {
    b := &Board{white, black, cp, 0, NewNullControl(), NewNullControl()}
    b.MakeMove(move)
    return b.black == bresult && b.white == wresult
}

func TestMakeMove(t *testing.T) {
    // flip right
    if !testmove(
        A3|A5|A7|A6,
        A8,
        A4,
        A3,
        A4|A5|A6|A7|A8,
        WHITE) {
        t.Errorf("Flip right: A4 fails to flip A5-7 when A8 is owned.")
    }
    // flip left
    if !testmove(
        B1|B8,
        B2|B4|B5|B7,
        B3,
        B1|B2|B3|B8,
        B4|B5|B7,
        BLACK) {
        t.Errorf("Flip left: B3 fails to flip B2 when B1 is owned.")
    }
    // flip left and right
    if !testmove(
        H1|H6,
        H2|H3|H5,
        H4,
        H1|H2|H3|H4|H5|H6,
        Bitboard(0),
        BLACK) {
        t.Errorf("Flip L and R: H4 fails to flip H2, H3, H5 with H1, H6 owned.")
    }
    // no flip
    if !testmove(
        E1|E3|E4|E5|E6,
        D2|F2|G2,
        E2,
        E1|E3|E4|E5|E6,
        D2|F2|G2|E2,
        WHITE) {
        t.Errorf("No flip: E2 should flip nothing when none of E is owned.")
    }
    //flip up
    if !testmove(
        B2|G2|H2,
        C2|D2|H8,
        E2,
        B2|C2|D2|H2|E2|G2,
        H8,
        BLACK) {
        t.Errorf("Flip up: E2 should flip C2 and D2 when B2 owned.")
    }
    //flip down
    if !testmove(
        F4|G4,
        H4,
        E4,
        Bitboard(0),
        F4|E4|G4|H4,
        WHITE) {
        t.Errorf("Flip down: E4 should flip F4 and G4 when H4 owned.")
    }
    //flip up and down
    if !testmove(
        A8|F8,
        B8|D8|E8,
        C8,
        C8|B8|D8|E8|A8|F8,
        Bitboard(0),
        BLACK) {
        t.Errorf("Flip U & D: C8 should flip B8, D8, E8 with A8, F8 owned.")
    }
    //no flip
    if !testmove(
        cols[0]^G1,
        rows[0]^A1,
        G1,
        cols[0]^G1,
        rows[0]^A1|G1,
        WHITE) {
        t.Errorf("No flip: G1 should flip nothing with none of col 1 owned.")
    }
    //flip 45
    if !testmove(
        B4,
        D2|C3,
        E1,
        B4|D2|C3|E1,
        Bitboard(0),
        BLACK) {
        t.Errorf("Flip 45: E1 should flip D2 and C3 with B4 owned.")
    }
    //flip 180+45=225
    if !testmove(
        D6|E5,
        F4,
        C7,
        Bitboard(0),
        C7|D6|E5|F4,
        WHITE) {
        t.Errorf("Flip 225: C7 should flip D6 and E5 with F4 owned.")
    }
    //flip 45 and 225
    if !testmove(
        G2|B7,
        A1|F3|D5|C6,
        E4,
        E4|F3|D5|C6|G2|B7,
        A1,
        BLACK) {
        t.Errorf("Flip 45+225: E4 should flip F3, D5, and C6 with G2 and B7 owned.")
    }
    //no flip
    if !testmove(
        diag45[4],
        diag45[6],
        E2,
        diag45[4],
        diag45[6]|E2,
        WHITE) {
        t.Errorf("No Flip: E2 should flip nothing when none of its diagonal owned.")
    }
    //flip 135
    if !testmove(
        A2,
        C4|B3,
        D5,
        A2|C4|B3|D5,
        Bitboard(0),
        BLACK) {
        t.Errorf("Flip 135: D5 should flip C4 and B3 with A2 owned")
    }
    //flip 180+135=315
    if !testmove(
        G3,
        H4,
        F2,
        Bitboard(0),
        F2|H4|G3,
        WHITE) {
        t.Errorf("Flip 315: F2 should flip G3 with H4 owned.")
    }
    //flip 135 and 315
    if !testmove(
        A4|E8,
        B5|D7,
        C6,
        diag135[4],
        Bitboard(0),
        BLACK) {
        t.Errorf("Flip 135 and 315: C6 should flip B5 and D7 with A4 and E8 owned.")
    }
    //no flip
    if !testmove(
        diag135[9],
        diag135[11],
        E2,
        diag135[9],
        diag135[11]|E2,
        WHITE) {
        t.Errorf("No flip: E2 should not flip anything when none of its diagonal owned.")
    }
    //test all directions
    if !testmove(
        A1|C1|E1|E3|E5|C5|A5|A3,
        B2|B3|B4|C4|D4|D3|D2|C2,
        C3,
        C3|A1|C1|E1|E3|E5|C5|A5|A3|B2|B3|B4|C4|D4|D3|D2|C2,
        Bitboard(0),
        BLACK) {
        t.Errorf("Flip all: C3 should have flipped all its neighbors.")
    }
}

func TestGetScore(t *testing.T) {
    board := InitBoard(NewNullControl(), NewNullControl())
    if board.GetScore(board.CurPlayer()) != 2 {
        t.Errorf("Black's score should be 2 at start of game.")
    }
    board.MakeMove(D3)
    board.NextPlayer()
    if board.GetScore(board.CurPlayer()) != 1 {
        t.Errorf("White's score should be 1 after Black does D3.")
    }
}

func TestGetLegalMoves(t *testing.T) {
    board := InitBoard(NewNullControl(), NewNullControl())
    if board.GetLegalMoves() & C4 == 0 {
        t.Errorf("Black to C4 is legal at the start of the game.")
    }
    board.NextPlayer()
    if board.GetLegalMoves() & C4 != 0 {
        t.Errorf("White to C4 is illegal at the start of the game.")
    }
}

func BenchmarkMakeMove(b *testing.B) {
    // let's see how fast this shiz is
    board := &Board{cols[1], cols[3], BLACK,
        0, NewNullControl(), NewNullControl()}
    board2 := &Board{B1 | E1 | H1 | A8 | E8 | H7, diag45[7] |
        diag135[8] | rows[4] | cols[3] ^ (B1 | E1 | H1 | A8 | E8 | H7), BLACK,
        0, NewNullControl(), NewNullControl()}
    board3 := &Board{B3 | B8 | E6 | D4, B4 | B5 | B7 | C5 | C6 | D6, BLACK,
        0, NewNullControl(), NewNullControl()}
    for i := 0; i < b.N; i++ {
        board.MakeMove(E3)
        board2.MakeMove(E4)
        board3.MakeMove(B6)
        board2.black = B1 | E1 | H1 | A8 | E8 | H7
        board2.white = diag45[7] | diag135[8] | rows[4] | cols[3] ^ (B1 | E1 | H1 | A8 | E8 | H7)
        board3.black = B3 | B8 | E6 | D4
        board3.white = B4 | B5 | B7 | C5 | C6 | D6
    }
}
