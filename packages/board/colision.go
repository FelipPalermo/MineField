package board

import "fmt"

func (b *Board) colision(uI int) bool {

	if b.Place[uI].bomb == true {
		fmt.Print("\nBomb, you died!")
		return true
	} else {
		b.Place[uI].sel = true
		return false
	}
}
