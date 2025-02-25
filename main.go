package main

import (
	"chessgogame/chessboard/handler"
	"chessgogame/helpers"
	"net/http"
)

const (
	PORT = ":8080"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	mux := http.NewServeMux()
	mux.Handle("GET /", fs)
	myboard := helpers.NewBoard()
	handler.NewChessBoardHandler(mux, &myboard)

	http.ListenAndServe(PORT, mux)
}
