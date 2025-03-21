package entities

type SkillRating struct {
	ID       uint
	SkillID  uint  `json:"skill_id"`
	Skill    Skill `json:"skill"`
	MemberID uint  `json:"member_id"`
	Rating   int   `json:"rating"`
}

func NewSkillRating(mem_id uint, skill Skill) *SkillRating {
	return &SkillRating{
		SkillID:  skill.ID,
		Skill:    skill,
		MemberID: mem_id,
		Rating:   1,
	}
}
