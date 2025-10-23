package domain

// type GameObjectTemplate interface {
// 	GetTemplateID() string
// 	GetName() string
// 	GetCardType() CardType
// }

type CardTemplate struct {
	// GameObjectTemplate
	TemplateId string
	Name       string
	CardType   CardType //spell, unit, land, etc

	//Card specific fields?
	Cost    int //0 for enemies?
	Health  int
	Attack  int
	Color   CardColor
	Subtype CardSubtype

	/*
		Eventually:
			Activate Main List
			Trigger List
	*/

}

// type EnemyTemplate struct {
// 	GameObjectTemplate
// 	TemplateId string
// 	Name       string
// 	CardType   CardType
// }
