package entities

type SkillHolder interface {
	HasSkill(uint) bool
	GetType() string
	GetID() uint
}
