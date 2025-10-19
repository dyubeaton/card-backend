// Structs that wrap Instance layer and sometiems implement behavioral interfaces
// Are still "instances", just more specific
package domain

type AllyUnit struct {
	Unit
	Instance      *CardInstance
	CurrentPower  int
	CurrentHealth int
	CanAttack     bool //not sure if this is neccesary
	Position      []Position
}

func (au *AllyUnit) GetPosition() []Position
func (au *AllyUnit) GetCurrentHealth() int
func (au *AllyUnit) GetCanAttack() bool

type EnemyUnit struct {
	Unit
	Instance      *EnemyInstance
	CurrentHealth int
	CanAttack     bool //not sure if this is neccesary
	Position      []Position
}

func (eu *EnemyUnit) GetPosition() []Position
func (eu *EnemyUnit) GetCurrentHealth() int
func (eu *EnemyUnit) GetCanAttack() bool

type Spell struct {
	//does it need to wrap an instance?
	Instance *CardInstance
}

type Land struct {
	Permanent
	Instance *CardInstance
	Position []Position
}

func (l *Land) GetPosition() []Position

type Equip struct {
	//Not a Permenant
	Holder *Permanent //Maybe they can be equip to lands? Otherwise this should say Unit

}

//TODO:
//EnemyLand
//EnemyEquip
