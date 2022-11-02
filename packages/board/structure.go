package board

type place struct {
	posX int
	posY int
	bomb bool
	sel  bool
}

type Board struct {
	Place []place
}
