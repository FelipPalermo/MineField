package board

import "fmt"

func Play(b Board, b2 *Board) {

	var uI int

	for {

		if b.points() >= 3 {
			fmt.Print("You win!")
			break
		}

		b.showBoard()

		fmt.Println("Select the coordinate")
		fmt.Scanf("%d", &uI)

		if b.colision(uI) == true {
			fmt.Print("\n")
			break
		}

	}
}
