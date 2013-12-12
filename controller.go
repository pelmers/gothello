package gothello

// Types that implement this interface are supported by the game.
type Controller interface {
    // When it's this controller's turn, the game calls its Move method.
    // Move is given a pointer to the Game, and it must return a move.
    // Return Bitboard(0) to indicate no move.
    // Game does not enforce move legality, so please
    // do so within the controller.
    Move(*Game) Bitboard
}

// NullControl does not make any moves. Used for testing.
type NullControl struct{}

func NewNullControl() NullControl             { return NullControl{} }
func (NullControl) Move(g *Game) Bitboard { return Bitboard(0) }
