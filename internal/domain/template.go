package domain

type GameObjectTemplate interface {
	GetTemplateID() string
}

type CardTemplate struct {
	GameObjectTemplate
}

type EnemyTemplate struct {
	GameObjectTemplate
}
