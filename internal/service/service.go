package service

import (
	"github.com/pulsone21/powner/internal/repos"
)

func NewMemberService(mem repos.MemberRepository) *MemberService {
	return &MemberService{
		repo: mem,
	}
}

func NewTeamService(team repos.TeamRepository) *TeamService {
	return &TeamService{
		repo: team,
	}
}

func NewSkillService(skill repos.SkillRepository) *SkillService {
	return &SkillService{
		repo: skill,
	}
}

func NewSkillManagement(mem repos.MemberRepository, team repos.TeamRepository, skill repos.SkillRepository) *SkillManagement {
	return &SkillManagement{
		memberRepo: mem,
		teamRepo:   team,
		skillRepo:  skill,
	}
}

func NewMemberManagement(mem repos.MemberRepository, team repos.TeamRepository, skill repos.SkillRepository) *MemberManagementService {
	return &MemberManagementService{
		memberRepo: mem,
		teamRepo:   team,
		skillRepo:  skill,
	}
}
