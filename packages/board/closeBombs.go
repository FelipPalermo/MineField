package board

import "strconv"

func (b Board) closeBombs(uI int) string {

	// centro
	//b.Place[uI].bomb

	var bCounter int = 0

	const x int = 5
	n := uI

	for i := 0; i < 3; i++ {
		y := -1

		for j := 0; j < 3; j++ {

			if i == 0 && uI > 5 && b.Place[n-x+y].bomb == true {
				y++
				bCounter++
			}

			if i == 1 && b.Place[n+y].bomb == true {

				y++
				bCounter++

			}

			if i == 2 && uI < 20 && b.Place[n+x+y].bomb == true {

				y++
				bCounter++
			}

		}
	}
	tZ := strconv.Itoa(bCounter)
	return tZ
}
