package board

import "fmt"

func Dificulty() (int, int) {

	var uI int

	for {

		fmt.Println("Chose your dificulty!\n1 - easy\n2 - medium\n3 - hard.")
		fmt.Scanf("%d", &uI)

		switch {

		case uI == 1:
			return 9, 3

		case uI == 2:
			return 16, 4

		case uI == 3:
			return 25, 5

		default:
			fmt.Println("Invalid value.")
		}
	}
}
