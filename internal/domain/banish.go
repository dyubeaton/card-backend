package domain

type Banish struct {
	Cards []Card
}

func NewBanish() *Banish {
	return &Banish{
		Cards: make([]Card, 0),
	}
}
