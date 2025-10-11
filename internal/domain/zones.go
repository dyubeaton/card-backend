package domain

//The following types are rather similar, simple containers

type Deck struct {
	Cards []Card
}

func NewDeck() *Deck {

	return &Deck{
		Cards: make([]Card, 0),
	}
}

type Banish struct {
	Cards []Card
}

func NewBanish() *Banish {
	return &Banish{
		Cards: make([]Card, 0),
	}
}

type Trash struct {
	Cards []Card
}

func NewTrash() *Trash {
	return &Trash{
		Cards: make([]Card, 0),
	}
}

type Hand struct {
	Cards []Card
}

func NewHand() *Hand {
	return &Hand{
		Cards: make([]Card, 0),
	}
}
