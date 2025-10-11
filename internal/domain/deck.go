package domain

type Deck struct {
	Cards []Card
}

func NewDeck() *Deck {

	return &Deck{
		Cards: make([]Card, 0),
	}
}
