package gothello

import (
    "bufio"
    "fmt"
    "os"
)

type HumanController struct {
    reader *bufio.Reader
}

func NewHumanController() *HumanController {
    return &HumanController{bufio.NewReader(os.Stdin)}
}

func (h *HumanController) GetMove(b *Board) Bitboard {
    allowed := b.GetLegalMoves()
    if allowed == Bitboard(0) {
        return Bitboard(0)
    }
    fmt.Printf("%s to play.\n%s\nBlack: %d\nWhite: %d\n", b.CurPlayerName(), b,
        b.GetScore(BLACK), b.GetScore(WHITE))
    fmt.Printf("Legal moves: %s\n", allowed)
    fmt.Print("Enter the location of your move: ")
    movestr, _ := h.reader.ReadString('\n')
    if len(movestr) >= 2 {
        r := uint(movestr[0] - 'A')
        c := uint(movestr[1] - '1')
        move := RC2Mask(r, c)
        if move&allowed != 0 {
            return move
        }
    }
    fmt.Println("Sorry, your input was not a legal move. Please retry.")
    return h.GetMove(b)
}
