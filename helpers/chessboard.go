package helpers

import "chessgogame/domain"

var CellCodeMap = map[int][]string{
	0: {"8", "a"},
	1: {"7", "b"},
	2: {"6", "c"},
	3: {"5", "d"},
	4: {"4", "e"},
	5: {"3", "f"},
	6: {"2", "g"},
	7: {"1", "h"},
}

type Cell struct {
	CellCode  string
	CellPiece Piece
}

type Board [][]Cell

type MyBoard struct {
	Board *Board
}

func NewBoard() MyBoard {
	board := make(Board, domain.ROWS)
	myboard := MyBoard{Board: &board}
	for i := range board {
		board[i] = make([]Cell, domain.COLUMNS)
	}
	return myboard
}

func GetCellCodeValue(i, j int) string {
	v1, exists1 := CellCodeMap[i]
	v2, exists2 := CellCodeMap[j]
	if exists1 && exists2 {
		return v2[1] + v1[0]
	}
	return ""
}
