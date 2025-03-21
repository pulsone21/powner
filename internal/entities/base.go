package entities

type SkillHolder interface {
	HasSkill(uint) bool
	Entity
}

type Entity interface {
	GetType() string
	GetID() uint
}
