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
	TypeAlliedUnit CardType = iota
	TypeAlliedLand
	TypeAlliedEquip
	TypeSpell
	//Enemy Types (TODO: figure out if this is neccesary)
	TypeEnemyUnit
	TypeEnemyLand
	TypeEnemyEquip
)

func (c CardType) String() string {
	switch c {
	case TypeAlliedUnit:
		return "AlliedUnit"
	case TypeAlliedLand:
		return "AlliedLand"
	case TypeAlliedEquip:
		return "AlliedEquip"
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

type CardColor int

const (
	ColorRed CardColor = iota
	ColorBlue
	ColorGreen
	ColorBlack
)

func (c CardColor) String() string {
	switch c {
	case ColorRed:
		return "ColorRed"
	case ColorBlue:
		return "ColorBlue"
	case ColorGreen:
		return "ColorGreen"
	case ColorBlack:
		return "ColorBlack"
	default:
		return "unknown"
	}
}

type CardSubtype int

const (
	SubtypeNone CardSubtype = iota
	SubtypeKnight
	SubtypeDemon
	SubtypeSprite
	SubtypeMilitia
	//To be exapnded
)
