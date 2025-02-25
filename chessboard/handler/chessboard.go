package handler

import (
	"chessgogame/api"
	"chessgogame/domain"
	"chessgogame/helpers"
	"fmt"
	"net/http"
)

type ChessBoardHandler struct {
	Myboard *helpers.MyBoard
}

func NewChessBoardHandler(r *http.ServeMux, myboard *helpers.MyBoard) {

	handler := &ChessBoardHandler{
		Myboard: myboard,
	}
	r.HandleFunc("GET /initialise-board", handler.InitialiseBoard)
	r.HandleFunc("GET /get-state", handler.GetState)
	r.HandleFunc("POST /make-move", handler.MakeMove)
	r.HandleFunc("GET /possible-moves", handler.PossibleMoves)
}

func (c *ChessBoardHandler) InitialiseBoard(w http.ResponseWriter, r *http.Request) {
	//rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR
	//7K/3N2qp/b5r1/2p1Q1N1/Pp4PK/7P/1P3p2/6r1
	fen := helpers.NewFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR")
	fen.DistributePieces(c.Myboard)

	boardVal := *c.Myboard.Board
	for i := 0; i < domain.ROWS; i++ {
		for j := 0; j < domain.COLUMNS; j++ {
			boardVal[i][j].CellCode = helpers.CellCodeMap[j][1] + helpers.CellCodeMap[i][0]
		}
	}

	fmt.Println(c.Myboard.Board)
	api.Success(w, 200, c.Myboard.Board)
}
func (c *ChessBoardHandler) PossibleMoves(w http.ResponseWriter, r *http.Request) {
	boardVal := *c.Myboard.Board
	movesMap := make(map[string][]string)
	for i := 0; i < domain.ROWS; i++ {
		for j := 0; j < domain.COLUMNS; j++ {
			piece := boardVal[i][j].CellPiece
			pieceType := piece.PieceType
			cellCode := boardVal[i][j].CellCode
			if pieceType != 0 {
				switch pieceType {
				case 1:
					{
						// fmt.Println("pawn")
						movesArr := piece.GetPawnMoves(i, j, c.Myboard.Board)
						movesMap[cellCode] = movesArr
					}
				case 2:
					{
						// fmt.Println("Knight")
						movesArr := piece.GetKnightMoves(i, j, c.Myboard.Board)
						movesMap[cellCode] = movesArr

					}
				case 3:
					{
						// fmt.Println("Bishop")
						movesArr := piece.GetBishopMoves(i, j, c.Myboard.Board)
						movesMap[cellCode] = movesArr
					}
				case 4:
					{
						// fmt.Println("Rook")
						movesArr := piece.GetRookMoves(i, j, c.Myboard.Board)
						movesMap[cellCode] = movesArr
					}
				case 5:
					{
						// fmt.Println("Queen")
						movesArr := piece.GetQueenMoves(i, j, c.Myboard.Board)
						movesMap[cellCode] = movesArr
					}
				case 6:
					{
						// fmt.Println("King")
						movesArr := piece.GetKingMoves(i, j, c.Myboard.Board)
						movesMap[cellCode] = movesArr
					}
				}
			}
		}
	}
	res := domain.PossibleMoveResponse{
		Turn:  "white",
		Moves: movesMap,
	}
	api.Success(w, 200, res)
}
func (c *ChessBoardHandler) GetState(w http.ResponseWriter, r *http.Request) {
}
func (c *ChessBoardHandler) MakeMove(w http.ResponseWriter, r *http.Request) {
}
