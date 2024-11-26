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

type SkillManagementTestSuite struct {
	suite.Suite
	service SkillManagement
}

func (s *SkillManagementTestSuite) SetupSuite() {
	// ensuring that TearDown is called no matter if setup was successfull
	setupDone := false
	defer func() {
		if !setupDone {
			s.TearDownSuite()
		}
	}()

	db_file := "./testing/skillManagementTest.db"
	_, err := os.Create(db_file)
	if err != nil {
		panic(err)
	}
	db, err := database.CreateDB(db_file,
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		panic(err)
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
	s1, _ := s_repo.Create(*entities.NewSkill("Splunk", "Splunk", entities.Hard, 2))
	skill, err := s_repo.Create(*entities.NewSkill("SOAR", "SOAR", entities.Hard, 2))
	if err != nil {
		log.Println(err)
		return
	}

	t_repo.AddSkill(*t, *s1)
	t_repo.AddSkill(*t, *skill)
	t_repo.AddMember(*t, *m)

	s.service = *NewSkillManagement(m_repo, t_repo, s_repo)
	setupDone = true
}

func (s *SkillManagementTestSuite) TearDownSuite() {
	os.Remove("./testing/skillManagementTest.db")
}

func TestSkillManagementTestSuite(t *testing.T) {
	suite.Run(t, new(SkillManagementTestSuite))
}

func (s *SkillManagementTestSuite) TestAddSKillToTeam() {
	testCases := []struct {
		Name        string
		SkillID     uint
		TeamID      uint
		ExpectedErr error
		HasSkill    bool
	}{
		{
			Name:        "Successfully add Skill",
			SkillID:     1,
			TeamID:      1,
			ExpectedErr: nil,
			HasSkill:    true,
		},
		{
			Name:        "Skill already on Team",
			TeamID:      1,
			ExpectedErr: nil,
			SkillID:     1,
			HasSkill:    true,
		},
		{
			Name:        "Skill does not exist",
			TeamID:      1,
			ExpectedErr: BadRequest,
			SkillID:     250,
			HasSkill:    false,
		},
		{
			Name:        "Team does not exist",
			TeamID:      1250,
			ExpectedErr: BadRequest,
			SkillID:     1,
			HasSkill:    false,
		},
	}

	for _, tC := range testCases {
		s.Run(tC.Name, func() {
			t, err := s.service.AddSkillToTeam(fmt.Sprint(tC.TeamID), fmt.Sprint(tC.SkillID))

			s.ErrorIs(err, tC.ExpectedErr)

			if tC.ExpectedErr == nil {
				s.Equal(tC.HasSkill, t.HasSkill(tC.SkillID))
				if err != nil {
					panic(err)
				}

				if t != nil { // we have a test case with unknown team
					s.Equal(tC.HasSkill, t.HasSkill(tC.SkillID))
					for _, m := range t.Members {
						s.True(
							m.HasSkill(tC.SkillID),
							fmt.Sprintf("Member: %v, should have the Skill: %v by now", m.ID, tC.SkillID))
					}
				}
			}

			if !tC.HasSkill {
				s.Nil(t)
			}
		})
	}
}

func (s *SkillManagementTestSuite) TestRemoveSkillFromTeam() {
	testCases := []struct {
		Name          string
		SkillID       uint
		TeamID        uint
		ExpectedError error
		SkillLen      int
	}{
		{
			Name:          "Successfully remove skill from team",
			SkillID:       1,
			TeamID:        1,
			ExpectedError: nil,
			SkillLen:      0,
		},
		{
			Name:          "Failed - skill not found",
			SkillID:       250,
			TeamID:        1,
			ExpectedError: BadRequest,
			SkillLen:      1,
		},
		{
			Name:          "Failed - Team not found",
			SkillID:       1,
			TeamID:        250,
			ExpectedError: BadRequest,
			SkillLen:      1,
		},
		{
			Name:          "Failed - skill not in Team",
			SkillID:       1, // we already removed him, in the first test case
			TeamID:        1,
			ExpectedError: BadRequest,
			SkillLen:      1,
		},
		{
			Name:          "Handling emptying the skill list",
			SkillID:       2,
			TeamID:        1,
			ExpectedError: nil,
			SkillLen:      0,
		},
	}

	for _, tC := range testCases {
		s.Run(tC.Name, func() {
			t, err := s.service.RemoveSkillToTeam(fmt.Sprint(tC.TeamID), fmt.Sprint(tC.SkillID))
			s.ErrorIs(err, tC.ExpectedError)

			if t != nil { // we have a testcase for a team which dosen't exists
				s.Equal(tC.SkillLen, len(t.Skills))
				s.False(t.HasSkill(tC.SkillID))
			}
		})
	}
}

func (s *SkillManagementTestSuite) TestAddSkillToMember() {
	testCases := []struct {
		Name          string
		MemberID      uint
		SkillID       uint
		ExpectedError error
	}{
		{
			Name:          "Successfull",
			MemberID:      1,
			SkillID:       1,
			ExpectedError: nil,
		},
		{
			Name:          "Member unknown",
			MemberID:      250,
			SkillID:       1,
			ExpectedError: BadRequest,
		},
		{
			Name:          "Skill unknown",
			MemberID:      1,
			SkillID:       250,
			ExpectedError: BadRequest,
		},
		{
			Name:          "Skill already on member",
			MemberID:      1,
			SkillID:       1,
			ExpectedError: BadRequest,
		},
	}

	for _, tC := range testCases {
		s.Run(tC.Name, func() {
			m, err := s.service.AddSkillToMember(fmt.Sprint(tC.MemberID), fmt.Sprint(tC.SkillID), 1)

			s.ErrorIs(err, tC.ExpectedError)

			if tC.ExpectedError == nil {
				s.True(m.HasSkill(tC.SkillID))
			} else {
				s.Nil(m)
			}
		})
	}
}

func (s *SkillManagementTestSuite) TestUpdateSkillRating() {
	testCases := []struct {
		Name          string
		MemberID      uint
		SkillID       uint
		Rating        int
		ExpectedError error
	}{
		{
			Name:          "Successfull",
			MemberID:      1,
			SkillID:       1,
			Rating:        3,
			ExpectedError: nil,
		},
		{
			Name:          "Member unknown",
			MemberID:      250,
			SkillID:       1,
			Rating:        0,
			ExpectedError: BadRequest,
		},
		{
			Name:          "Skill unknown",
			MemberID:      1,
			SkillID:       250,
			Rating:        0,
			ExpectedError: BadRequest,
		},
		{
			Name:          "Skill not on member",
			MemberID:      1,
			SkillID:       1,
			Rating:        0,
			ExpectedError: BadRequest,
		},
		{
			Name:          "Rating to low",
			MemberID:      1,
			SkillID:       1,
			Rating:        0,
			ExpectedError: BadRequest,
		},
		{
			Name:          "Rating to high",
			MemberID:      1,
			SkillID:       1,
			Rating:        6,
			ExpectedError: BadRequest,
		},
	}

	for _, tC := range testCases {
		s.Run(tC.Name, func() {
			m, err := s.service.UpdateSkillRating(fmt.Sprint(tC.MemberID), fmt.Sprint(tC.SkillID), tC.Rating)

			s.ErrorIs(err, tC.ExpectedError)

			if tC.ExpectedError == nil {
				s.True(m.HasSkill(tC.SkillID))
				sR := m.GetSkillRatingBySkill(tC.SkillID)
				s.Equal(tC.Rating, sR.Rating)
			} else {
				s.Nil(m)
			}
		})
	}
}
