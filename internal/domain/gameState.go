package domain

type GameState struct {
	//player card cycle
	Deck   *Deck
	Hand   *Hand
	Trash  *Trash
	Banish *Banish

	//enemy instance cycle
	EnemyTrash  *Trash
	EnemyBanish *Banish

	//shared
	Board *Board
}

func NewGameState() *GameState {
	return &GameState{
		//Initialize fields

		//Likely call initializaers of each component
		Deck:        NewDeck(),
		Board:       NewBoard(),
		Hand:        NewHand(),
		Trash:       NewTrash(),
		EnemyTrash:  NewTrash(),
		EnemyBanish: NewBanish(),
		Banish:      NewBanish(),
	}

}
