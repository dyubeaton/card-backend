package domain

type Trash struct {
	Cards []Card
}

func NewTrash() *Trash {
	return &Trash{
		Cards: make([]Card, 0),
	}
}
