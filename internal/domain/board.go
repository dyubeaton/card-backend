package domain

type Board struct {
	Board [][]Tile
}

func NewBoard() *Board {
	return &Board{
		Board: make([][]Tile, 0),
		//TODO: should take in rows/cols and make each slice within the outer slice
	}
}
