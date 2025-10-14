package domain

//The following types are rather similar, simple containers

//TODO: Enemy zones might need their own types since these reference CardInstances.
//So something like EnemyTrash []EnemyInstance

type Deck struct {
	Cards []CardInstance
}

func NewDeck() *Deck {

	return &Deck{
		Cards: make([]CardInstance, 0),
	}
}

type Banish struct {
	Cards []CardInstance
}

func NewBanish() *Banish {
	return &Banish{
		Cards: make([]CardInstance, 0),
	}
}

type Trash struct {
	Cards []CardInstance
}

func NewTrash() *Trash {
	return &Trash{
		Cards: make([]CardInstance, 0),
	}
}

type Hand struct {
	Cards []CardInstance
}

func NewHand() *Hand {
	return &Hand{
		Cards: make([]CardInstance, 0),
	}
}
