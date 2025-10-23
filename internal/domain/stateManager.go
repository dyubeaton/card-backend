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

	templateRegistry *TemplateRegistry
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

// Mid level function that moves any instance.
// Actually have this be called by other high level functions, like DestroyCard, PlayPermenantToBoard, etc
// Does not need to handle wrapping/unwrapping
func (gsm *StateManager) MoveInstance(inst *CardInstance, destination Zone) error {

	//Original collection
	originCollection, err := gsm.getCollectionFromZone(inst.Location)

	if err != nil {
		return err
	}

	destinationCollection, err := gsm.getCollectionFromZone(destination)

	if err != nil {
		return err
	}

	//Remove from original collection
	err = originCollection.RemoveInstance(inst.InstanceID)

	if err != nil {
		return err
	}

	err = destinationCollection.AddInstance(inst)

	if err != nil {
		return err
	}

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

// Wrap the provided instance using its type from its template
// TODO: Probably remove distinction between allied/enemy instances and add that as a field in instance
func (gsm *StateManager) wrapInstancePermenant(instance *CardInstance) Permanent {

	template := gsm.templateRegistry.GetTemplate(instance.InstanceID)

	switch template.CardType {
	case TypeAlliedUnit:
		return NewAlliedUnit(instance)
	case TypeAlliedLand:
		return NewAlliedLand(instance)
	case TypeAlliedEquip:
		return NewAlliedEquip(instance)
	case TypeEnemyUnit:
		return NewEnemyUnit(instance)
	case TypeEnemyEquip:
		return NewEnemyEquip(instance)
	case TypeEnemyLand:
		return NewEnemyLand(instance)
	default:
		return nil
	}

}

// Move permenant off board
// Unwrap logic
func (gsm *StateManager) MoveInstanceOffBoard(instance *CardInstance, destination Zone) error

// Wrap instance then move it onto the board
// Return the wrapper
func (gsm *StateManager) MoveInstanceOnBoard(instance *CardInstance) (Permanent, error)

/*
Thoughts on more methods, this should be the hub for any reasonable action I want events/event actions to take
Basically the "outcomes" of anything should be recorded through these methods

I've covered state changed regarding the placement of the cards, basically the "zone" field and of course reflecting the change on the game state
This whole game state is meant to capture the world of the instances and what holds them. I've most (i think) mapped out the "what holds them"
Thus, moving on, I should think about anything else that could modify an instance. (I still need to think about how persistent/status effects will work between the GSM and
controllers, hopefully I can put this off for now)

UnitTakesDamage
UnitHeals
UnitMoves
AddPermenant/StatusEffect?

*/
