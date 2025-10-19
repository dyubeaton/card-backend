package domain

type Board struct {
	InstanceCollection
	Board [][]Tile
}

func NewBoard() *Board {
	return &Board{
		Board: make([][]Tile, 0),
		//TODO: should take in rows/cols and make each slice within the outer slice
	}
}

type Tile struct {
	Row int
	Col int
	//Should they hold cards specifically or something more abstract like a "Permanent" type?
	Permanent *Permanent
}
