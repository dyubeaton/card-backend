// Represent any object that impacts the game
// give each their own id and track where they are, essentially the top layer of abstraction
package domain

//Interface represents any "instance" that is any initialized game object that comes from a template
// type GameObjectInstance interface {
// 	GetInstanceID() string
// 	GetTemplate() GameObjectTemplate
// 	GetLocation() Zone
// }

//The broad, generic type that represents an instance of a game object
//It won't implement interfaces other than GameObjectInstance, rather it will be wrapped within more specific types when needed
//Ex: Troop takes an instance of CardInstance, and troop will implement unit.
type CardInstance struct {
	// GameObjectInstance
	InstanceID string
	Template   CardTemplate
	Location   Zone
	IsAlly     bool
	//TODO: Cost here?
	//TODO: Subtypes here? Other categorization like color?
}

//Consider: are all enemies neccesarily units? Perhaps they have landmarks as well?
//In that case the same separation should occur as in CardInstance
// type EnemyInstance struct {
// 	GameObjectInstance
// 	InstanceID string
// 	Template   EnemyTemplate
// 	Location   Zone
// }
