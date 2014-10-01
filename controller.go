package gothello

// Types that implement this interface are supported by the game.
// When it's this controller's turn, the game calls its Move method.
// Move is given a pointer to the Game, and it must return a move.
// Return Bitboard(0) to indicate no move.
// Game does not enforce move legality, so please do so within the controller.
type Controller interface {
	Move(*Game) Bitboard
}

// NullControl does not make any moves. Used for testing.
type NullControl struct{}

func NewNullControl() NullControl         { return NullControl{} }
func (NullControl) Move(g *Game) Bitboard { return Bitboard(0) }

// borrowed from Norvig's Paradigms of AI programming
var weights [64]int = [64]int{
	-120, -20, 20, 5, 5, 20, -20, 120, -20, -40, -5, -5, -5, -5, -40, -20,
	20, -5, 3, 3, 3, 3, -5, 20, 5, -5, 3, 3, 3, 3, -5, 5,
	-120, -20, 20, 5, 5, 20, -20, 120, -20, -40, -5, -5, -5, -5, -40, -20,
	20, -5, 3, 3, 3, 3, -5, 20, 5, -5, 3, 3, 3, 3, -5, 5,
}

func Evaluate(board Bitboard) int {
	score := 0
	for i, w := range weights {
		score += w * int((board>>uint(i))&1)
	}
	return score
}
