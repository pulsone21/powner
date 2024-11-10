package service

import (
	"github.com/pulsone21/powner/internal/repos"
)

type Service struct {
	TeamRepo   repos.TeamRepository
	MemberRepo repos.MemberRepository
	SkillRepo  repos.SkillRepository
}

func NewService(team repos.TeamRepository, member repos.MemberRepository, skill repos.SkillRepository) *Service {
	return &Service{
		TeamRepo:   team,
		MemberRepo: member,
		SkillRepo:  skill,
	}
}
