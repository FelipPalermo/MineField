package main

import bd "minefield/packages/board"

func main() {

	board := bd.Board{
		Place: bd.ArrPlace(bd.Dificulty())}

	bd.Play(board, &board)

}
