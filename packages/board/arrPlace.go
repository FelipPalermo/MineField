package board

func ArrPlace(amnt, div int) []place {

	px := 2
	py := 0

	board := make([]place, 0)

	for i := 0; i < amnt; i++ {

		board = append(board, newPlace())

		board[len(board)-1].posX = px
		board[len(board)-1].posY = py

		if px%div == 0 {
			px = 0
			py = 1
		}
		if px > 0 && py > 0 {
			py = 0
		}
		px++
	}
	return board
}
