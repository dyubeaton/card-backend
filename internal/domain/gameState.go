package domain

type GameState struct {
	Deck   *Deck
	Board  *Board
	Hand   *Hand
	Trash  *Trash
	Banish *Banish
}

func NewGameState() *GameState {
	return &GameState{
		//Initialize fields

		//Likely call initializaers of each component
		Deck:   NewDeck(),
		Board:  NewBoard(),
		Hand:   NewHand(),
		Trash:  NewTrash(),
		Banish: NewBanish(),
	}

}
