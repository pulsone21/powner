package service

import (
	"fmt"
	"os"
	"testing"

	"github.com/pulsone21/powner/internal/database"
	"github.com/pulsone21/powner/internal/entities"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type TeamServiceTestSuite struct {
	suite.Suite
	service TeamService
}

func (s *TeamServiceTestSuite) SetupSuite() {
	db_file := "./testing/teamServiceTest.db"
	_, err := os.Create(db_file)
	if err != nil {
		panic(err)
	}

	db, err := database.CreateDB(db_file,
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		panic(err)
	}
	repo := database.NewTeamRepo(db)

	repo.Create(*entities.NewTeam("TestTeam", "Test Team description"))

	s.service = *NewTeamService(repo)
}

func (s *TeamServiceTestSuite) TearDownSuite() {
	os.Remove("./testing/teamServiceTest.db")
}

func TestTeamServiceSuite(t *testing.T) {
	suite.Run(t, new(TeamServiceTestSuite))
}

func (s *TeamServiceTestSuite) Test1CreateTeam() {
	testCases := []struct {
		TestCaseName string
		Name         string
		Description  string
		ExpectedErr  error
		ExpectedTeam *entities.Team
	}{
		{
			TestCaseName: "Successfull",
			Description:  "TestSkill 2 description",
			Name:         "TestTeam 2",
			ExpectedErr:  nil,
			ExpectedTeam: &entities.Team{
				Name:        "TestTeam 2",
				Description: "TestSkill 2 description",
				Model:       gorm.Model{ID: 2},
			},
		},
		{
			TestCaseName: "Failed - Name validation error",
			Description:  "Team description",
			Name:         "T",
			ExpectedErr:  BadRequest,
			ExpectedTeam: nil,
		},
		{
			TestCaseName: "Failed - Description validation error",
			Description:  "T",
			Name:         "Optimus Prime",
			ExpectedErr:  BadRequest,
			ExpectedTeam: nil,
		},
		{
			TestCaseName: "Failed - Name already Used",
			Description:  "TestSkill description",
			Name:         "TestTeam",
			ExpectedErr:  BadRequest,
			ExpectedTeam: nil,
		},
	}

	for _, tC := range testCases {
		s.Run(tC.TestCaseName, func() {
			req := entities.TeamRequest{
				Name:        tC.Name,
				Description: tC.Description,
			}

			t, err := s.service.CreateTeam(req)

			if tC.ExpectedErr != nil {
				s.ErrorContains(err, tC.ExpectedErr.Error())
				s.Nil(t)
			} else {
				s.Nil(err)
				s.Equal(tC.ExpectedTeam.ID, t.ID)
				s.Equal(tC.ExpectedTeam.Name, t.Name)
				s.Equal(tC.ExpectedTeam.Description, t.Description)
			}
		})
	}
}

func (s *TeamServiceTestSuite) Test2GetAllTeam() {
	t, err := s.service.GetTeams()
	s.Nil(err, "Error should be nil, but is not")
	s.Len(*t, 2, fmt.Sprintf("I should have 2 Teams in the database, but have: %b", len(*t)))
}

func (s *TeamServiceTestSuite) Test3GetTeamByID() {
	testCases := []struct {
		tcName        string
		ExpectedError error
		ExpectedTeam  *entities.Team
		ID            string
	}{
		{
			tcName:        "Successfull",
			ID:            "1",
			ExpectedError: nil,
			ExpectedTeam: &entities.Team{
				Name:        "TestTeam",
				Description: "Test Team description",
			},
		},
		{
			tcName:        "Team not found",
			ID:            "250",
			ExpectedError: nil,
			ExpectedTeam:  nil,
		},
	}

	for _, tC := range testCases {
		s.Run(tC.tcName, func() {
			m, err := s.service.GetTeamByID(tC.ID)

			if tC.ExpectedError != nil {
				s.ErrorContains(err, tC.ExpectedError.Error())
			} else {
				s.Nil(err)
			}

			if tC.ExpectedTeam != nil {
				s.Equal(tC.ExpectedTeam.Name, m.Name)
				s.Equal(tC.ExpectedTeam.Description, m.Description)
				s.Equal(tC.ID, fmt.Sprint(m.ID))
			}
		})
	}
}

func (s *TeamServiceTestSuite) Test4UpdateSkill() {
	testCases := []struct {
		tcName        string
		TeamID        string
		UpdatedTeam   *entities.Team
		ExpectedError error
		ExpectedTeam  *entities.Team
	}{
		{
			tcName: "Successfull - Name Update",
			UpdatedTeam: &entities.Team{
				Name:        "Team Name Update",
				Description: "Test Team description",
			},
			TeamID:        "1",
			ExpectedError: nil,
			ExpectedTeam: &entities.Team{
				Name:        "Team Name Update",
				Description: "Test Team description",
			},
		},

		{
			tcName: "Successfull - Description Update",
			UpdatedTeam: &entities.Team{
				Name:        "Team Name Update",
				Description: "Test Team description update",
			},
			TeamID:        "1",
			ExpectedError: nil,
			ExpectedTeam: &entities.Team{
				Name:        "Team Name Update",
				Description: "Test Team description update",
			},
		},
		{
			tcName: "Failed - No Changes",
			UpdatedTeam: &entities.Team{
				Name:        "Team Name Update",
				Description: "Test Team description update",
			},
			TeamID:        "1",
			ExpectedError: BadRequest,
			ExpectedTeam:  nil,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.tcName, func() {
			req := entities.TeamRequest{
				Name:        tc.UpdatedTeam.Name,
				Description: tc.UpdatedTeam.Description,
			}

			m, err := s.service.UpdateTeam(tc.TeamID, req)
			if tc.ExpectedError != nil {
				s.ErrorContains(err, tc.ExpectedError.Error(), fmt.Sprintf("Error should be: %v but is actually: %v", tc.ExpectedError, err))
			} else {
				s.Nil(err)
			}

			if tc.ExpectedTeam != nil {
				s.Equal(tc.ExpectedTeam.Name, m.Name)
				s.Equal(tc.ExpectedTeam.Description, m.Description)
			} else {
				s.Nil(m)
			}
		})
	}
}

func (s *TeamServiceTestSuite) Test5DeleteSkill() {
	testCases := []struct {
		Name          string
		TeamID        string
		ExpectedError error
	}{
		{
			Name:          "Successfull",
			TeamID:        "1",
			ExpectedError: nil,
		},
		{
			Name:          "Failed",
			TeamID:        "250",
			ExpectedError: BadRequest,
		},
	}

	for _, tC := range testCases {
		s.Run(tC.Name, func() {
			err := s.service.DeleteTeam(tC.TeamID)
			if tC.ExpectedError != nil {
				s.ErrorContains(err, tC.ExpectedError.Error(), fmt.Sprintf("Expected Error should be %v but is actually %v", tC.ExpectedError, err))
			} else {
				s.Nil(err)
			}
		})
	}
}
