// Behavorial Layer: Implemented by role wrappers
package domain

//Represent anything that has a position on the board
type Permanent interface {
	//shouldnt need to implement GameObjInstance
	GetPosition() []Position
	// IsOnBoard() bool //is this neccesary?
}

//More specific, anything that can engage in combat
type Unit interface {
	Permanent

	GetCurrentHealth() int
	// GetBaseHealth() int Do we need base functions in here? Should it be left to templates?
	CanAttack() bool

	//certainly many other methods

	//Specific methods for change?
	//TakeDamage() int ?

}
