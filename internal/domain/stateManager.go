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

import (
	"fmt"
)

type StateManager struct {
	GameState   *GameState
	GameContext *GameContext
	//Probably needs a reference to the orchestrater above a round (like what tracks the current deck, enemies, and what not) I imagine this process will link together
	//relevant templates and instances

}

// Init new GSM with context.
func NewStateManagerWithContext(g *GameContext) *StateManager {
	return &StateManager{
		GameContext: g,
	}
}

// Initialize GS field within GSM.
// Should estbalish template linking using context/other template tools
func (gsm *StateManager) InitGameState() {
	gsm.GameState = NewGameState()

	//TODO: actually fill out init logic using context information (include template linking and what not)
}

// High level function that moves any instance.
// Should handle any wrappping/similar state changes
func (gsm *StateManager) MoveInstance(inst GameObjectInstance, destination Zone) error {

	//Original collection
	originCollection, err := gsm.getCollectionFromZone(inst.GetLocation())

	if err != nil {
		return err
	}

	destinationCollection, err := gsm.getCollectionFromZone(destination)

	if err != nil {
		return err
	}

	//Remove from original collection
	err = originCollection.RemoveInstance(inst.GetInstanceID())

	if err != nil {
		return err
	}

	//TODO: Per instance dewrapping logic must occur around here

	err = destinationCollection.AddInstance(inst)

	if err != nil {
		return err
	}

	//TODO: Wrapping logic needs to happen somewhere around here

	return nil
}

func (gsm *StateManager) getCollectionFromZone(location Zone) (InstanceCollection, error) {
	switch location {
	case ZoneHand:
		return gsm.GameState.Hand, nil
	case ZoneTrash:
		return gsm.GameState.Trash, nil
	case ZoneDeck:
		return gsm.GameState.Deck, nil
	case ZoneBoard:
		return gsm.GameState.Board, nil
	case ZoneBanish:
		return gsm.GameState.Banish, nil
	case ZoneEnemyTrash:
		return gsm.GameState.EnemyTrash, nil
	case ZoneEnemyBanish:
		return gsm.GameState.EnemyBanish, nil
	default:
		return nil, fmt.Errorf("invalid zone:  %d ", location)
	}
}
