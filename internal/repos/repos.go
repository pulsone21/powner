package repos

import "github.com/pulsone21/powner/internal/entities"

type TeamRepository interface {
	GetAll() (*[]entities.Team, error)
	GetByID(id uint) (*entities.Team, error)
	Create(newTeam entities.Team) (*entities.Team, error)
	Update(newTeam entities.Team) (*entities.Team, error)
	Delete(id uint) error
	RemoveMember(team entities.Team, mem entities.Member) (*entities.Team, error)
	AddMember(team entities.Team, mem entities.Member) (*entities.Team, error)
	AddSkill(team entities.Team, skill entities.Skill) (*entities.Team, error)
	RemoveSkill(team entities.Team, skill entities.Skill) (*entities.Team, error)
}

type MemberRepository interface {
	GetAll() (*[]entities.Member, error)
	GetByID(id uint) (*entities.Member, error)
	Create(newMember entities.Member) (*entities.Member, error)
	Update(newMember entities.Member) (*entities.Member, error)
	Delete(id uint) error
	AddSkill(member entities.Member, skill entities.Skill) (*entities.Member, error)
	UpdateSkillRating(skill_id uint, rating int) error
}

type SkillRepository interface {
	GetAll() (*[]entities.Skill, error)
	GetByID(id uint) (*entities.Skill, error)
	Create(newSkill entities.Skill) (*entities.Skill, error)
	Update(newSkill entities.Skill) (*entities.Skill, error)
	Delete(id uint) error
}
