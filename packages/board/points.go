package board

func (b Board) points() int {

	var points int = 0

	for _, sp := range b.Place {
		if sp.sel == true {
			points++
		}
	}
	return points
}
