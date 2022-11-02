package board

import "fmt"

func (b Board) showBoard() {

	fmt.Println()
	for _, sp := range b.Place {
		if sp.sel == true {
			fmt.Print("X  ")
		} else {
			fmt.Print("O  ")
		}

		if sp.posY == 1 {
			fmt.Println("")
		}
	}
	fmt.Println()
}
