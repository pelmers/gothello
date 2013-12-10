package gothello

// Types that implement this interface are supported by the game.
type Controller interface {
    // When it's this controller's turn, the game calls its GetMove method.
    // GetMove is given a pointer to the Board, and it must return a move.
    // Return Bitboard(0) to indicate no move.
    // Board does not enforce move legality, so please
    // do so within the controller.
    GetMove(*Board) Bitboard
}

// NullControl does not make any moves. Used for testing.
type NullControl struct{}

func NewNullControl() NullControl             { return NullControl{} }
func (NullControl) GetMove(b *Board) Bitboard { return Bitboard(0) }
