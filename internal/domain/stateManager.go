//Formaliztion on what this should do:
/*
Should either implement or call the most fine grained operations I've specified for my gameState. All the operations will be on Card Instances and their relation
to their storage components in GameState.
Considerations: (Should this handle the action AND its consequence?) Example: For something like "move card from deck to hand", there are two parts, moving the instance
	in storage(so from *Deck to *Hand) and changing the information on the instance itself (instance.Location updated to Hand). These are two different "fine grained"
	functions. The question is should I implement the broader function?


Answer: I should cover the action and its consequence. If it were to be handled by EventActions, then the EventActions need to "know" implementation logic in the data layer,
	breaking that separation.
*/
package domain

type StateManager struct {
	GameState   *GameState
	GameContext *GameContext
	//Probably needs a reference to the orchestrater above a round (like what tracks the current deck, enemies, and what not) I imagine this process will link together
	//relevant templates and instances

}

//Init new GSM with context.
func NewStateManagerWithContext(g *GameContext) *StateManager {
	return &StateManager{
		GameContext: g,
	}
}

//Initialize GS field within GSM.
//Should estbalish template linking using context/other template tools
func (gsm *StateManager) InitGameState() {
	gsm.GameState = NewGameState()

	//TODO: actually fill out init logic using context information (include template linking and what not)
}
