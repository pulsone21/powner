package service

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/pulsone21/powner/internal/database"
	"github.com/pulsone21/powner/internal/entities"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MemberManagementTestSuite struct {
	suite.Suite
	service MemberManagementService
}

func TestMemberManagementTestSuite(t *testing.T) {
	suite.Run(t, new(MemberManagementTestSuite))
}

func (s *MemberManagementTestSuite) SetupSuite() {
	// ensuring that TearDown is called no matter if setup was successfull
	setupDone := false
	defer func() {
		if !setupDone {
			s.TearDownSuite()
		}
	}()

	db_file := "./testing/membereManagementTest.db"
	_, err := os.Create(db_file)
	if err != nil {
		log.Println(err)
		return
	}

	db, err := database.CreateDB(db_file, &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		log.Println(err)
		return
	}

	m_repo := database.NewMemberRepo(db)
	m_repo.Create(*entities.NewMember("Jon Doe", 30))
	m, err := m_repo.Create(*entities.NewMember("Foo Bar", 30))
	if err != nil {
		log.Println(err)
		return
	}

	t_repo := database.NewTeamRepo(db)
	t, err := t_repo.Create(*entities.NewTeam("TestTeam", "TestTeam"))
	if err != nil {
		log.Println(err)
		return
	}

	s_repo := database.NewSkillRepo(db)
	skill, err := s_repo.Create(*entities.NewSkill("Test Skill", "Test Skill", entities.Hard, 2))
	if err != nil {
		log.Println(err)
		return
	}

	t_repo.AddSkill(*t, *skill)
	t_repo.AddMember(*t, *m)

	s.service = *NewMemberManagement(m_repo, t_repo, s_repo)
	setupDone = true
}

func (s *MemberManagementTestSuite) TearDownSuite() {
	os.Remove("./testing/membereManagementTest.db")
}

func (s *MemberManagementTestSuite) TestAddMemberToTeam() {
	testCases := []struct {
		Name        string
		MemberID    uint
		TeamID      string
		ExpectedErr error
		HasMember   bool
		SkillID     uint
		HasSkill    bool
	}{
		{
			Name:        "Successfully add Member",
			MemberID:    1,
			TeamID:      "1",
			ExpectedErr: nil,
			HasMember:   true,
			SkillID:     1,
			HasSkill:    true,
		},
		{
			Name:        "Member already on Team",
			MemberID:    2,
			TeamID:      "1",
			ExpectedErr: nil,
			HasMember:   true,
			SkillID:     1,
			HasSkill:    true,
		},
		{
			Name:        "Member does not exist",
			MemberID:    250,
			TeamID:      "1",
			ExpectedErr: BadRequest,
			HasMember:   false,
			SkillID:     1,
			HasSkill:    false,
		},
		{
			Name:        "Team does not exist",
			MemberID:    2,
			TeamID:      "1250",
			ExpectedErr: BadRequest,
			HasMember:   false,
			SkillID:     1,
			HasSkill:    false,
		},
	}

	for _, tC := range testCases {
		s.Run(tC.Name, func() {
			t, err := s.service.AddMemberToTeam(tC.TeamID, fmt.Sprint(tC.MemberID))

			s.ErrorIs(err, tC.ExpectedErr)

			if tC.ExpectedErr == nil {
				s.Equal(tC.HasMember, t.HasMember(tC.MemberID))
				m, err := s.service.memberRepo.GetByID(tC.MemberID)
				if err != nil {
					panic(err)
				}
				s.Equal(tC.HasSkill, m.HasSkill(tC.SkillID))
			}

			if !tC.HasMember {
				s.Nil(t)
			}
		})
	}
}

func (s *MemberManagementTestSuite) TestRemoveMemberFromTeam() {
	testCases := []struct {
		Name          string
		MemberID      uint
		TeamID        uint
		ExpectedError error
		MemberLen     int
	}{
		{
			Name:          "Successfully remove member from team",
			MemberID:      2,
			TeamID:        1,
			ExpectedError: nil,
			MemberLen:     1,
		},
		{
			Name:          "Failed to remove member from team - Member not found",
			MemberID:      250,
			TeamID:        1,
			ExpectedError: BadRequest,
			MemberLen:     1,
		},
		{
			Name:          "Failed to remove member from team - Team not found",
			MemberID:      1,
			TeamID:        250,
			ExpectedError: BadRequest,
			MemberLen:     1,
		},
		{
			Name:          "Failed to remove member from team - Member not in Team",
			MemberID:      2, // we already removed him, in the first test case
			TeamID:        1,
			ExpectedError: BadRequest,
			MemberLen:     1,
		},
		{
			Name:          "Handling emptying the member list",
			MemberID:      1,
			TeamID:        1,
			ExpectedError: nil,
			MemberLen:     0,
		},
	}

	for _, tC := range testCases {
		s.Run(tC.Name, func() {
			_, err := s.service.RemoveMemberToTeam(fmt.Sprint(tC.TeamID), fmt.Sprint(tC.MemberID))
			s.ErrorIs(err, tC.ExpectedError)

			t, err := s.service.teamRepo.GetByID(tC.TeamID)
			if err != nil {
				panic(err)
			}

			if t != nil { // we have a testcase for a team which dosen't exists
				s.Equal(tC.MemberLen, len(t.Members))
				s.False(t.HasMember(tC.MemberID))
			}
		})
	}
}
