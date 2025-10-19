package domain

//The following types are rather similar, simple containers

//TODO: Enemy zones might need their own types since these reference CardInstances.
//So something like EnemyTrash []EnemyInstance

//Consideration: Since all these do the same thing, perhaps I just lump them as something like a CardInterface container? They will have unique names in GameState
//so I can make more generic functions that move from one collection to the other?

type CardInstanceCollection struct {
	Cards []CardInstance
}

func NewCardInstanceCollection() *CardInstanceCollection {
	return &CardInstanceCollection{
		Cards: make([]CardInstance, 0),
	}
}

//Basic wrapper for adding card instance to a collection
func (cic *CardInstanceCollection) AddCard(ci CardInstance) {
	_ = append(cic.Cards, ci)
}

type EnemyInstanceCollection struct {
	Enemies []EnemeyInstance
}

func NewEnemyInstanceCollection() *EnemyInstanceCollection {
	return &EnemyInstanceCollection{
		Enemies: make([]EnemeyInstance, 0),
	}
}

//Basic wrapper for adding enemy instance to a collection
func (eic *EnemyInstanceCollection) AddEnemy(ei EnemeyInstance) {
	_ = append(eic.Enemies, ei)
}

type Deck struct {
	Cards []CardInstance
}

func NewDeck() *Deck {

	return &Deck{
		Cards: make([]CardInstance, 0),
	}
}

// type Banish struct {
// 	Cards []CardInstance
// }

// func NewBanish() *Banish {
// 	return &Banish{
// 		Cards: make([]CardInstance, 0),
// 	}
// }

// type Trash struct {
// 	Cards []CardInstance
// }

// func NewTrash() *Trash {
// 	return &Trash{
// 		Cards: make([]CardInstance, 0),
// 	}
// }

// type Hand struct {
// 	Cards []CardInstance
// }

// func NewHand() *Hand {
// 	return &Hand{
// 		Cards: make([]CardInstance, 0),
// 	}
// }

// type EnemyTrash struct {
// 	Enemies []EnemeyInstance
// }

// type EnemyBanish struct {
// 	Enemies []EnemeyInstance
// }
