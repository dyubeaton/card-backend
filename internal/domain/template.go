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
	CardType   CardType

	//Card specific fields?
	Cost int
}

func (ct *CardTemplate) GetTemplateID() string

func (ct *CardTemplate) GetName() string

func (ct *CardTemplate) GetCardType() CardType

// type EnemyTemplate struct {
// 	GameObjectTemplate
// 	TemplateId string
// 	Name       string
// 	CardType   CardType
// }
