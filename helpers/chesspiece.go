package helpers

import "chessgogame/domain"

type Piece struct {
	PieceType             int
	Color                 int
	DefaultMoveDirections [][2]int
}

func (p Piece) GetKnightMoves(i, j int, board *Board) []string {
	var movesArr []string
	boardval := *board
	moveLocations := [][2]int{
		{i + 2, j - 1},
		{i + 2, j + 1},
		{i - 2, j - 1},
		{i - 2, j + 1},
		{i + 1, j - 2},
		{i + 1, j + 2},
		{i - 1, j - 2},
		{i - 1, j + 2},
	}
	for _, val := range moveLocations {
		cellCode := GetCellCodeValue(val[0], val[1])
		if cellCode != "" {
			piece := boardval[val[0]][val[1]]
			if piece.CellPiece.Color == p.Color {
				continue
			}
			movesArr = append(movesArr, cellCode)
		}
	}
	return movesArr
}
func (p Piece) GetPawnMoves(i, j int, board *Board) []string {
	var movesArr []string
	boardval := *board
	moveLocations := GetPieceSlideMoves(p, i, j, 1, boardval)
	for _, val := range moveLocations {
		cellCode := GetCellCodeValue(val[0], val[1])
		if cellCode != "" {
			movesArr = append(movesArr, cellCode)
		}
	}
	return movesArr
}
func (p Piece) GetBishopMoves(i, j int, board *Board) []string {
	var movesArr []string
	boardval := *board
	moveLocations := GetPieceSlideMoves(p, i, j, -1, boardval)
	for _, val := range moveLocations {
		cellCode := GetCellCodeValue(val[0], val[1])
		if cellCode != "" {
			movesArr = append(movesArr, cellCode)
		}
	}
	return movesArr
}
func (p Piece) GetRookMoves(i, j int, board *Board) []string {
	var movesArr []string
	boardval := *board
	moveLocations := GetPieceSlideMoves(p, i, j, -1, boardval)
	for _, val := range moveLocations {
		cellCode := GetCellCodeValue(val[0], val[1])
		if cellCode != "" {
			movesArr = append(movesArr, cellCode)
		}
	}
	return movesArr
}
func (p Piece) GetQueenMoves(i, j int, board *Board) []string {
	var movesArr []string
	boardval := *board
	moveLocations := GetPieceSlideMoves(p, i, j, -1, boardval)

	for _, val := range moveLocations {
		cellCode := GetCellCodeValue(val[0], val[1])
		if cellCode != "" {
			movesArr = append(movesArr, cellCode)
		}
	}
	return movesArr
}
func (p Piece) GetKingMoves(i, j int, board *Board) []string {
	var movesArr []string
	boardval := *board
	moveLocations := GetPieceSlideMoves(p, i, j, 1, boardval)
	for _, val := range moveLocations {
		cellCode := GetCellCodeValue(val[0], val[1])
		if cellCode != "" {
			movesArr = append(movesArr, cellCode)
		}
	}
	return movesArr
}

func GetPieceSlideMoves(p Piece, i, j, slide int, board Board) [][2]int {
	var moveLocations [][2]int
	var k = 0
	directions := p.DefaultMoveDirections
	for _, d := range directions {
		newRow, newCol := i+d[0], j+d[1]
		k = 0
		for newRow >= 0 && newRow < domain.ROWS && newCol >= 0 && newCol < domain.COLUMNS {
			if slide >= 0 {
				if k >= slide {
					break
				}
			}
			piece := board[newRow][newCol]
			if piece.CellPiece.Color == p.Color {
				break
			} else if piece.CellPiece.PieceType != 0 {
				moveLocations = append(moveLocations, [2]int{newRow, newCol})
				break
			} else {
				moveLocations = append(moveLocations, [2]int{newRow, newCol})
			}
			newRow += d[0]
			newCol += d[1]
			k++
		}
	}
	return moveLocations
}
