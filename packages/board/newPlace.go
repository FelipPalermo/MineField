package board

import (
	"math/rand"
	"time"
)

func rn() bool {

	rand.Seed(time.Now().UnixNano())
	if rand.Intn(2) == 0 {
		return true
	} else {
		return false
	}

}

func newPlace() place {

	p := place{}
	p.posX = 0
	p.posY = 0
	p.bomb = rn()
	p.sel = false

	return p
}
