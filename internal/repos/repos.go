package repos

import "github.com/pulsone21/powner/internal/entities"

type TeamRepository interface {
	GetAll() (*[]entities.Team, error)
	GetByID(id string) (*entities.Team, error)
	Create(newTeam entities.Team) (*entities.Team, error)
	Update(newTeam entities.Team) (*entities.Team, error)
	Delete(id string) error
	RemoveMember(team_id string, mem entities.Member) error
	AddMember(team_id string, mem entities.Member) error
	AddSkill(team_id string, skill entities.Skill) error
	RemoveSkill(team_id string, skill entities.Skill) error
}

type MemberRepository interface {
	GetAll() (*[]entities.Member, error)
	GetByID(id string) (*entities.Member, error)
	Create(newTeam entities.Member) (*entities.Member, error)
	Update(newTeam entities.Member) (*entities.Member, error)
	Delete(id string) error
}

type SkillRepository interface {
	GetAll() (*[]entities.Skill, error)
	GetByID(id string) (*entities.Skill, error)
	Create(newTeam entities.Skill) (*entities.Skill, error)
	Update(newTeam entities.Skill) (*entities.Skill, error)
	Delete(id string) error
}
