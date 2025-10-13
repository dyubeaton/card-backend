package domain

import "fmt"

//simple value types

type Position [2]int

func (p Position) String() string {
	return fmt.Sprintf("(%d , %d)", p[0], p[1])
}

type Zone int

const (
	ZoneHand Zone = iota
	ZoneTrash
	ZoneDeck
	ZoneBanish
	ZoneBoard
	ZoneEnemyTrash
	ZoneEnemyBanish
)

func (z Zone) String() string {
	switch z {
	case ZoneHand:
		return "Hand"
	case ZoneTrash:
		return "Trash"
	case ZoneDeck:
		return "Deck"
	case ZoneBanish:
		return "Banish"
	case ZoneBoard:
		return "Board"
	case ZoneEnemyTrash:
		return "EnemyTrash"
	case ZoneEnemyBanish:
		return "EnemyBanish"
	default:
		return "Unknown"
	}
}

//counts enemies too, I don't make the rules (wait I do?)
//Should I have separate types for enemies? (EnemyUnit, EnemyLandmark?)
type CardType int

const (
	TypeTroop CardType = iota
	TypeLand
	TypeEquip
	TypeSpell
	//Enemy Types (TODO: figure out if this is neccesary)
	TypeEnemyUnit
	TypeEnemyLand
	TypeEnemyEquip
)

func (c CardType) String() string {
	switch c {
	case TypeTroop:
		return "Troop"
	case TypeLand:
		return "Land"
	case TypeEquip:
		return "Equip"
	case TypeSpell:
		return "Spell"
	case TypeEnemyUnit:
		return "EnemyUnit"
	case TypeEnemyLand:
		return "EnemyLand"
	case TypeEnemyEquip:
		return "EnemyEquip"
	default:
		return "unknown"
	}
}
