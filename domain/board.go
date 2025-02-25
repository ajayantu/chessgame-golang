package domain

const (
	ROWS    = 8
	COLUMNS = 8
)
const (
	Empty = iota
	Pawn
	Knight
	Bishop
	Rook
	Queen
	King
	White
	Black
)

type PossibleMoveResponse struct {
	Turn  string              `json:"turn"`
	Moves map[string][]string `json:"moves"`
}
