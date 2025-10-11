package domain

//Idea is to have this represent anything that can be on the board
//Furthermore: units (trrops and enemies) landmarks, obstacles

//TODO: change this to a proper id mechanism, likely needs to be moved to another file
type GameId string

type Permanent interface {
	GameObjectInstance

	GetPosition() []Position
	IsOnBoard() bool
}

type Unit interface {
	Permanent

	GetHealth() int
	GetBaseHealth() int
	CanAttack() bool
}
