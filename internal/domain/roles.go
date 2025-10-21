// Structs that wrap Instance layer and sometiems implement behavioral interfaces
// Are still "instances", just more specific
/*
Some General Guidelines:
	These roles basically only apply when the instance is "active" or on board
	(active is basically just for spells)

	They are wrapped when activated/put on board

	Dewrapped when sent off board/completed activation

	Shouldn't actually need to implment the unwrap functionality here?
*/
package domain

type AllyUnit struct {
	Unit
	Instance      *CardInstance
	CurrentPower  int
	CurrentHealth int
	CanAttack     bool //not sure if this is neccesary
	Position      []Position
}

//Wrap given card instance
func NewAlliedUnit(ci CardInstance) *AllyUnit

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

func NewEnemyUnit(ci EnemyInstance) *EnemyUnit

func (eu *EnemyUnit) GetPosition() []Position
func (eu *EnemyUnit) GetCurrentHealth() int
func (eu *EnemyUnit) GetCanAttack() bool

type Spell struct {
	//does it need to wrap an instance?
	Instance *CardInstance
}

func NewSpell(ci CardInstance) *Spell

type AlliedLand struct {
	Permanent
	Instance *CardInstance
	Position []Position
}

func NewAlliedLand(ci CardInstance) *AlliedLand
func (l *AlliedLand) GetPosition() []Position

type AlliedEquip struct {
	//Not a Permenant
	Holder *Permanent //Maybe they can be equip to lands? Otherwise this should say Unit

}

func NewAlliedEquip(ci CardInstance) *AlliedEquip

//TODO:
//EnemyLand
//EnemyEquip
