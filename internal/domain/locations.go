package domain

import (
	"fmt"
)

// All instance collections will implement this
// Will need basic wrappers so I can make generic high level methods
type InstanceCollection interface {
	ContainsInstance(instanceID string) bool
	AddInstance(instance *CardInstance) error
	RemoveInstance(instanceID string) error
	GetAll() []*CardInstance
	Size() int
}

// Generic unordered collection. Used for trash, banish, and hand
type CardInstanceCollection struct {
	Cards []*CardInstance
}

// Make new CIC
func NewCardInstanceCollection() *CardInstanceCollection {
	return &CardInstanceCollection{
		Cards: make([]*CardInstance, 0),
	}
}

func (cic *CardInstanceCollection) ContainsInstance(instanceID string) bool {
	for _, card := range cic.Cards {
		if card.InstanceID == instanceID {
			return true
		}
	}
	return false
}

func (cic *CardInstanceCollection) AddInstance(instance *CardInstance) error {

	cic.Cards = append(cic.Cards, instance)
	return nil
}

func (cic *CardInstanceCollection) RemoveInstance(instanceID string) error {
	for i, card := range cic.Cards {
		if card.InstanceID == instanceID {
			// Remove by swapping with last element and truncating (order doesn't matter)
			cic.Cards[i] = cic.Cards[len(cic.Cards)-1]
			cic.Cards = cic.Cards[:len(cic.Cards)-1]
			return nil
		}
	}
	return fmt.Errorf("instance with ID %s not found in collection", instanceID)
}

func (cic *CardInstanceCollection) GetAll() []*CardInstance {
	return cic.Cards
}

func (cic *CardInstanceCollection) Size() int {
	return len(cic.Cards)
}

// type EnemyInstanceCollection struct {
// 	InstanceCollection
// 	Enemies []*CardInstance
// }

// func NewEnemyInstanceCollection() *EnemyInstanceCollection {
// 	return &EnemyInstanceCollection{
// 		Enemies: make([]*EnemyInstance, 0),
// 	}
// }

// func (eic *EnemyInstanceCollection) ContainsInstance(instanceID string) bool {
// 	for _, enemy := range eic.Enemies {
// 		if enemy.GetInstanceID() == instanceID {
// 			return true
// 		}
// 	}
// 	return false
// }

// func (eic *EnemyInstanceCollection) AddInstance(instance CardInstance) error {
// 	enemyInstance, ok := instance.(*EnemyInstance)
// 	if !ok {
// 		return fmt.Errorf("expected *EnemyInstance, got %T", instance)
// 	}
// 	eic.Enemies = append(eic.Enemies, enemyInstance)
// 	return nil
// }

// func (eic *EnemyInstanceCollection) RemoveInstance(instanceID string) error {
// 	for i, enemy := range eic.Enemies {
// 		if enemy.GetInstanceID() == instanceID {
// 			// Remove by swapping with last element and truncating
// 			eic.Enemies[i] = eic.Enemies[len(eic.Enemies)-1]
// 			eic.Enemies = eic.Enemies[:len(eic.Enemies)-1]
// 			return nil
// 		}
// 	}
// 	return fmt.Errorf("instance with ID %s not found in collection", instanceID)
// }

// func (eic *EnemyInstanceCollection) GetAll() []CardInstance {
// 	result := make([]CardInstance, len(eic.Enemies))
// 	for i, enemy := range eic.Enemies {
// 		result[i] = enemy
// 	}
// 	return result
// }

// func (eic *EnemyInstanceCollection) Size() int {
// 	return len(eic.Enemies)
// }

type Deck struct {
	InstanceCollection
	Cards []*CardInstance
}

func NewDeck() *Deck {

	return &Deck{
		Cards: make([]*CardInstance, 0),
	}
}

func (d *Deck) ContainsInstance(instanceID string) bool {
	for _, card := range d.Cards {
		if card.InstanceID == instanceID {
			return true
		}
	}
	return false
}

func (d *Deck) AddInstance(instance *CardInstance) error {
	// cardInstance, ok := instance.(*CardInstance)
	// if !ok {
	// 	return fmt.Errorf("expected *CardInstance, got %T", instance)
	// }
	d.Cards = append(d.Cards, instance)
	return nil
}

func (d *Deck) RemoveInstance(instanceID string) error {
	for i, card := range d.Cards {
		if card.InstanceID == instanceID {
			// For deck, preserve order by shifting elements
			d.Cards = append(d.Cards[:i], d.Cards[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("instance with ID %s not found in deck", instanceID)
}

func (d *Deck) GetAll() []*CardInstance {

	return d.Cards
}

func (d *Deck) Size() int {
	return len(d.Cards)
}
