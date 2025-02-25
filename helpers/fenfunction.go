package helpers

import (
	"chessgogame/domain"
	"fmt"
	"unicode"
)

type Fen struct {
	FenString string
}

func NewFen(fen string) *Fen {
	return &Fen{
		FenString: fen,
	}
}

func (f Fen) DistributePieces(board *MyBoard) {
	fenToMap := map[rune]Piece{
		'r': {
			PieceType: domain.Rook,
			Color:     domain.Black,
			DefaultMoveDirections: [][2]int{
				{1, 0},
				{-1, 0},
				{0, 1},
				{0, -1},
			},
		},
		'n': {
			PieceType:             domain.Knight,
			Color:                 domain.Black,
			DefaultMoveDirections: [][2]int{},
		},
		'b': {PieceType: domain.Bishop, Color: domain.Black, DefaultMoveDirections: [][2]int{
			{1, 1},
			{-1, -1},
			{-1, 1},
			{1, -1},
		}},
		'q': {PieceType: domain.Queen, Color: domain.Black, DefaultMoveDirections: [][2]int{
			{1, 1},
			{-1, -1},
			{-1, 1},
			{1, -1},
			{1, 0},
			{-1, 0},
			{0, 1},
			{0, -1},
		}},
		'k': {PieceType: domain.King, Color: domain.Black, DefaultMoveDirections: [][2]int{
			{1, 1},
			{-1, -1},
			{-1, 1},
			{1, -1},
			{1, 0},
			{-1, 0},
			{0, 1},
			{0, -1},
		}},
		'p': {PieceType: domain.Pawn, Color: domain.Black, DefaultMoveDirections: [][2]int{
			{1, 0},
		}},
		'R': {
			PieceType: domain.Rook,
			Color:     domain.White,
			DefaultMoveDirections: [][2]int{
				{1, 0},
				{-1, 0},
				{0, 1},
				{0, -1},
			},
		},
		'N': {
			PieceType:             domain.Knight,
			Color:                 domain.White,
			DefaultMoveDirections: [][2]int{},
		},
		'B': {PieceType: domain.Bishop, Color: domain.White, DefaultMoveDirections: [][2]int{
			{1, 1},
			{-1, -1},
			{-1, 1},
			{1, -1},
		}},
		'Q': {PieceType: domain.Queen, Color: domain.White, DefaultMoveDirections: [][2]int{
			{1, 1},
			{-1, -1},
			{-1, 1},
			{1, -1},
			{1, 0},
			{-1, 0},
			{0, 1},
			{0, -1},
		}},
		'K': {PieceType: domain.King, Color: domain.White, DefaultMoveDirections: [][2]int{
			{1, 1},
			{-1, -1},
			{-1, 1},
			{1, -1},
			{1, 0},
			{-1, 0},
			{0, 1},
			{0, -1},
		}},
		'P': {PieceType: domain.Pawn, Color: domain.White, DefaultMoveDirections: [][2]int{
			{-1, 0},
		}},
	}
	fmt.Println("fen is", f.FenString)
	boardVal := *board.Board
	row := 0
	col := 0
	for _, val := range f.FenString {
		if val == '/' {
			row++
			col = 0
		} else if unicode.IsDigit(val) {
			col += int(val - '0')
		} else {
			piece := fenToMap[val]
			boardVal[row][col].CellPiece = piece
			col++
		}
	}
}
