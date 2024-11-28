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

type SkillServiceTestSuite struct {
	suite.Suite
	service SkillService
}

func (s *SkillServiceTestSuite) SetupSuite() {
	db_file := "./testing/skillServiceTest.db"
	_, err := os.Create(db_file)
	if err != nil {
		panic(err)
	}

	db, err := database.CreateDB(db_file,
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		panic(err)
	}
	repo := database.NewSkillRepo(db)

	repo.Create(*entities.NewSkill("Splunk", "Splunk Skill", entities.Hard, 2))
	repo.Create(*entities.NewSkill("SOAR", "SOAR Skill", entities.Hard, 2))
	s.service = *NewSkillService(repo)
}

func (s *SkillServiceTestSuite) TearDownSuite() {
	os.Remove("./testing/skillServiceTest.db")
}

func TestSkillServiceSuite(t *testing.T) {
	suite.Run(t, new(SkillServiceTestSuite))
}

func (s *SkillServiceTestSuite) Test1CreateSkill() {
	testCases := []struct {
		TestCaseName  string
		Name          string
		Description   string
		Type          int
		Importance    int
		ExpectedErr   error
		ExpectedSkill *entities.Skill
	}{
		{
			TestCaseName: "Successfull",
			Description:  "TestSkill description",
			Name:         "TestSkill",
			Type:         int(entities.Hard),
			Importance:   2,
			ExpectedErr:  nil,
			ExpectedSkill: &entities.Skill{
				Name:        "TestSkill",
				Description: "TestSkill description",
				Type:        int(entities.Hard),
				Importance:  2,
				Model:       gorm.Model{ID: 3},
			},
		},
		{
			TestCaseName:  "Failed - Invalid Type",
			Description:   "TestSkill description",
			Name:          "TestSkill",
			Type:          10,
			Importance:    2,
			ExpectedErr:   BadRequest,
			ExpectedSkill: nil,
		},
		{
			TestCaseName:  "Failed - Invalid Importance",
			Description:   "TestSkill description",
			Name:          "TestSkill",
			Type:          1,
			Importance:    20,
			ExpectedErr:   BadRequest,
			ExpectedSkill: nil,
		},
		{
			TestCaseName:  "Failed - Invalid Description",
			Description:   "T",
			Name:          "TestSkill",
			Type:          0,
			Importance:    2,
			ExpectedErr:   BadRequest,
			ExpectedSkill: nil,
		},
		{
			TestCaseName:  "Failed - Invalid Name",
			Description:   "TestSkill description",
			Name:          "T",
			Type:          1,
			Importance:    2,
			ExpectedErr:   BadRequest,
			ExpectedSkill: nil,
		},
		{
			TestCaseName:  "Failed - Name already Used",
			Description:   "TestSkill description",
			Name:          "Splunk",
			Type:          1,
			Importance:    2,
			ExpectedErr:   BadRequest,
			ExpectedSkill: nil,
		},
	}

	for _, tC := range testCases {
		s.Run(tC.TestCaseName, func() {
			req := entities.SkillRequest{
				Name:        tC.Name,
				Description: tC.Description,
				Type:        tC.Type,
				Importance:  tC.Importance,
			}

			m, err := s.service.CreateSkill(req)

			if tC.ExpectedErr != nil {
				s.ErrorIs(err, tC.ExpectedErr)
				s.Nil(m)
			} else {
				s.Nil(err)
				s.Equal(tC.ExpectedSkill.ID, m.ID)
				s.Equal(tC.ExpectedSkill.Name, m.Name)
				s.Equal(tC.ExpectedSkill.Description, m.Description)
				s.Equal(tC.ExpectedSkill.Type, m.Type)
				s.Equal(tC.ExpectedSkill.Importance, m.Importance)
			}
		})
	}
}

func (s *SkillServiceTestSuite) Test2GetAllSkills() {
	m, err := s.service.GetSkills()
	s.Nil(err, "Error should be nil, but is not")
	// - We pre creating 2 skills and the create skills runs first which also creates one
	s.Len(*m, 3, fmt.Sprintf("I should have 3 Skills in the database, but have: %v", len(*m)))
}

func (s *SkillServiceTestSuite) Test3GetSkillByID() {
	testCases := []struct {
		tcName        string
		ExpectedError error
		ExpectedSkill *entities.Skill
		ID            string
	}{
		{
			tcName:        "Successfull",
			ID:            "2", // we are deleting 1
			ExpectedError: nil,
			ExpectedSkill: &entities.Skill{
				Name:        "SOAR",
				Description: "SOAR Skill",
				Type:        int(entities.Hard),
				Importance:  2,
			},
		},
		{
			tcName:        "Skill not found",
			ID:            "250",
			ExpectedError: nil,
			ExpectedSkill: nil,
		},
	}

	for _, tC := range testCases {
		s.Run(tC.tcName, func() {
			m, err := s.service.GetSkillByID(tC.ID)

			s.ErrorIs(tC.ExpectedError, err)
			if tC.ExpectedSkill != nil {
				s.Equal(tC.ExpectedSkill.Name, m.Name)
				s.Equal(tC.ExpectedSkill.Description, m.Description)
				s.Equal(tC.ExpectedSkill.Type, m.Type)
				s.Equal(tC.ExpectedSkill.Importance, m.Importance)
				s.Equal(tC.ID, fmt.Sprint(m.ID))
			}
		})
	}
}

func (s *SkillServiceTestSuite) Test4UpdateSkill() {
	testCases := []struct {
		tcName        string
		SkillID       string
		UpdatedSkill  *entities.Skill
		ExpectedError error
		ExpectedSkill *entities.Skill
	}{
		{
			tcName: "Successfull - Name Update",
			UpdatedSkill: &entities.Skill{
				Name:        "SOAR Name",
				Description: "SOAR Skill",
				Importance:  2,
				Type:        int(entities.Hard),
			},
			SkillID:       "2",
			ExpectedError: nil,
			ExpectedSkill: &entities.Skill{
				Name:        "SOAR Name",
				Description: "SOAR Skill",
				Importance:  2,
				Type:        int(entities.Hard),
			},
		},
		{
			tcName: "Successfull - Description Update",
			UpdatedSkill: &entities.Skill{
				Name:        "SOAR",
				Description: "SOAR Skill Description",
				Importance:  2,
				Type:        int(entities.Hard),
			},
			SkillID:       "2",
			ExpectedError: nil,
			ExpectedSkill: &entities.Skill{
				Name:        "SOAR",
				Description: "SOAR Skill Description",
				Importance:  2,
				Type:        int(entities.Hard),
			},
		},
		{
			tcName: "Succsessfull - Importance Update",
			UpdatedSkill: &entities.Skill{
				Name:        "Splunk",
				Description: "Splunk Skill",
				Importance:  3,
				Type:        int(entities.Hard),
			},
			SkillID:       "1",
			ExpectedError: nil,
			ExpectedSkill: &entities.Skill{
				Name:        "Splunk",
				Description: "Splunk Skill",
				Importance:  3,
				Type:        int(entities.Hard),
			},
		},
		{
			tcName: "Failed - No Changes",
			UpdatedSkill: &entities.Skill{
				Name:        "Splunk",
				Description: "Splunk Skill",
				Importance:  2,
				Type:        int(entities.Hard),
			},
			SkillID:       "1",
			ExpectedError: BadRequest,
			ExpectedSkill: nil,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.tcName, func() {
			req := entities.SkillRequest{
				Name:        tc.UpdatedSkill.Name,
				Description: tc.UpdatedSkill.Description,
			}

			m, err := s.service.UpdateSkill(tc.SkillID, req)
			s.ErrorIs(err, tc.ExpectedError, fmt.Sprintf("Error should be: %v but is actually: %v", tc.ExpectedError, err))

			if tc.ExpectedSkill != nil {
				s.Equal(tc.ExpectedSkill.Name, m.Name)
				s.Equal(tc.ExpectedSkill.Description, m.Description)
			} else {
				s.Nil(m)
			}
		})
	}
}

func (s *SkillServiceTestSuite) Test5DeleteSkill() {
	testCases := []struct {
		Name          string
		SkillID       string
		ExpectedError error
	}{
		{
			Name:          "Successfull",
			SkillID:       "1",
			ExpectedError: nil,
		},
		{
			Name:          "Failed",
			SkillID:       "250",
			ExpectedError: BadRequest,
		},
	}

	for _, tC := range testCases {
		s.Run(tC.Name, func() {
			err := s.service.DeleteSkill(tC.SkillID)
			s.ErrorIs(err, tC.ExpectedError, fmt.Sprintf("Expected Error should be %v but is actually %v", tC.ExpectedError, err))
		})
	}
}
