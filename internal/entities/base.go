package entities

type SkillHolder interface {
	HasSkill(uint) bool
	GetType() string
	GetID() uint
}

type Entity interface {
	Member | Skill | SkillRating | Team
}
