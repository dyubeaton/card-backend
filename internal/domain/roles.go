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
func NewAlliedUnit(ci *CardInstance) *AllyUnit {
	//TODO: Expand with other fields
	return &AllyUnit{
		Instance: ci,
	}
}

func (au *AllyUnit) GetPosition() []Position
func (au *AllyUnit) GetCurrentHealth() int
func (au *AllyUnit) GetCanAttack() bool

type EnemyUnit struct {
	Unit
	Instance      *CardInstance
	CurrentHealth int
	CanAttack     bool //not sure if this is neccesary
	Position      []Position
}

func NewEnemyUnit(ei *CardInstance) *EnemyUnit {
	return &EnemyUnit{
		Instance: ei,
	}
}

func (eu *EnemyUnit) GetPosition() []Position
func (eu *EnemyUnit) GetCurrentHealth() int
func (eu *EnemyUnit) GetCanAttack() bool

// type Spell struct {
// 	//does it need to wrap an instance?
// 	Instance *CardInstance
// }

// func NewSpell(ci *CardInstance) *Spell {
// 	return &Spell{
// 		Instance: &ci,
// 	}
// }

type AlliedLand struct {
	Permanent
	Instance *CardInstance
	Position []Position
}

func NewAlliedLand(ci *CardInstance) *AlliedLand {
	return &AlliedLand{
		Instance: ci,
	}
}
func (al *AlliedLand) GetPosition() []Position

type EnemyLand struct {
	Permanent
	Instance *CardInstance
	Position []Position
}

func NewEnemyLand(ci *CardInstance) *EnemyLand {
	return &EnemyLand{
		Instance: ci,
	}
}

func (el *EnemyLand) GetPosition() []Position

type AlliedEquip struct {
	Permanent
	Instance *CardInstance
	Holder   *Permanent //Maybe they can be equip to lands? Otherwise this should say Unit

}

func NewAlliedEquip(ci *CardInstance) *AlliedEquip {
	return &AlliedEquip{
		Instance: ci,
	}
}

//Probably just resturn Holder.GetPosition
func (ae *AlliedEquip) GetPosition() []Position

type EnemyEquip struct {
	Permanent
	Instance *CardInstance
	Holder   *Permanent
}

func NewEnemyEquip(ci *CardInstance) *EnemyEquip {
	return &EnemyEquip{
		Instance: ci,
	}
}

func (ee *EnemyEquip) GetPosition() []Position

//TODO:
//EnemyLand
//EnemyEquip
