package repos

import "github.com/pulsone21/powner/internal/entities"

type TeamRepository interface {
	GetAll() (*[]entities.Team, error)
	GetByID(id uint) (*entities.Team, error)
	Create(newTeam entities.Team) (*entities.Team, error)
	Update(newTeam entities.Team) (*entities.Team, error)
	Delete(id uint) error
	RemoveMember(team_id uint, mem entities.Member) (*entities.Team, error)
	AddMember(team_id uint, mem entities.Member) (*entities.Team, error)
	AddSkill(team_id uint, skill entities.Skill) (*entities.Team, error)
	RemoveSkill(team_id uint, skill entities.Skill) (*entities.Team, error)
}

type MemberRepository interface {
	GetAll() (*[]entities.Member, error)
	GetByID(id uint) (*entities.Member, error)
	Create(newTeam entities.Member) (*entities.Member, error)
	Update(newTeam entities.Member) (*entities.Member, error)
	Delete(id uint) error
}

type SkillRepository interface {
	GetAll() (*[]entities.Skill, error)
	GetByID(id uint) (*entities.Skill, error)
	Create(newTeam entities.Skill) (*entities.Skill, error)
	Update(newTeam entities.Skill) (*entities.Skill, error)
	Delete(id uint) error
}
