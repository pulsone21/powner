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

type MemberServiceTestSuite struct {
	suite.Suite
	db      *gorm.DB
	service MemberService
}

func (s *MemberServiceTestSuite) SetupSuite() {
	db_file := "./testing/memberServiceTest.db"
	_, err := os.Create(db_file)
	if err != nil {
		panic(err)
	}

	db, err := database.CreateDB(db_file,
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		panic(err)
	}
	s.db = db
	repo := database.NewMemberRepo(db)
	repo.Create(*entities.NewMember("Jon Doe", 30))
	repo.Create(*entities.NewMember("Foo Bar", 30))
	s.service = *NewMemberService(repo)
}

func (s *MemberServiceTestSuite) TearDownSuite() {
	os.Remove("./testing/memberServiceTest.db")
}

func TestMemberServiceTestSuite(t *testing.T) {
	suite.Run(t, new(MemberServiceTestSuite))
}

func (s *MemberServiceTestSuite) TestCreateMember() {
	testCases := []struct {
		TestCaseName   string
		Age            int
		Name           string
		ExpectedErr    error
		ExpectedMember *entities.Member
	}{
		{
			TestCaseName:   "Validation Error because of Age",
			Age:            14,
			Name:           "Jon Doe",
			ExpectedErr:    BadRequest,
			ExpectedMember: nil,
		},
		{
			TestCaseName: "Working Example",
			Age:          65,
			Name:         "Ida Boss",
			ExpectedErr:  nil,
			ExpectedMember: &entities.Member{
				Age:  65,
				Name: "Ida Boss",
			},
		},
		{
			TestCaseName:   "Validation Error because missing Name",
			Age:            25,
			Name:           "",
			ExpectedErr:    BadRequest,
			ExpectedMember: nil,
		},
	}

	for _, tC := range testCases {
		s.Run(fmt.Sprintf("Create Member: %v", tC.TestCaseName), func() {
			req := entities.MemberRequest{Name: tC.Name, Age: tC.Age}
			m, err := s.service.CreateMember(req)
			if err != nil {
				s.ErrorIs(err, tC.ExpectedErr)
				s.Nil(m)
			} else {
				s.Nil(err)
				s.Equal(tC.ExpectedMember.Name, m.Name)
				s.Equal(tC.ExpectedMember.Age, m.Age)
			}
		})
	}
}

func (s *MemberServiceTestSuite) TestDeleteMember() {
	testCases := []struct {
		Name          string
		MemberID      string
		ExpectedError error
	}{
		{
			Name:          "Successfull Deletion",
			MemberID:      "3",
			ExpectedError: nil,
		},
		{
			Name:          "Failed Deletion",
			MemberID:      "250",
			ExpectedError: nil,
		},
	}

	for _, tC := range testCases {
		s.Run(tC.Name, func() {
			err := s.service.DeleteMember(tC.MemberID)
			s.ErrorIs(err, tC.ExpectedError, fmt.Sprintf("Expected Error should be %v but is actually %v", tC.ExpectedError, err))
		})
	}
}

func (s *MemberServiceTestSuite) TestGetAllMembers() {
	m, err := s.service.GetMembers()
	s.Nil(err, "Error should be nil, but is not")
	s.Len(*m, 2, fmt.Sprintf("I should have 2 User in the database, but have: %b", len(*m)))
}

func (s *MemberServiceTestSuite) TestGetMemberByID() {
	testCases := []struct {
		tcName         string
		ExpectedError  error
		ExpectedMember *entities.Member
		ID             string
	}{
		{
			tcName:        "Get first User",
			ID:            "1",
			ExpectedError: nil,
			ExpectedMember: &entities.Member{
				Name: "Jon Doe",
				Age:  30,
			},
		},
		{
			tcName:         "User not found",
			ID:             "250",
			ExpectedError:  nil,
			ExpectedMember: nil,
		},
	}

	for _, tC := range testCases {
		s.Run(tC.tcName, func() {
			m, err := s.service.GetMemberByID(tC.ID)

			s.ErrorIs(tC.ExpectedError, err)
			if tC.ExpectedMember != nil {
				s.Equal(tC.ExpectedMember.Name, m.Name)
				s.Equal(tC.ExpectedMember.Age, m.Age)
			}
		})
	}
}

func (s *MemberServiceTestSuite) TestUpdateMember() {
	testCases := []struct {
		tcName         string
		MemberID       string
		Name           string
		Age            int
		ExpectedError  error
		ExpectedMember *entities.Member
	}{
		{
			tcName:        "Successfull Update of Member",
			Name:          "Bar Foo",
			Age:           45,
			MemberID:      "1",
			ExpectedError: nil,
			ExpectedMember: &entities.Member{
				Name: "Bar Foo",
				Age:  45,
			},
		},
		{
			tcName:         "Failed Updated of Member",
			Name:           "Bar Foo",
			Age:            45,
			MemberID:       "250",
			ExpectedError:  BadRequest,
			ExpectedMember: nil,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.tcName, func() {
			req := entities.MemberRequest{Name: tc.Name, Age: tc.Age}
			m, err := s.service.UpdateMember(tc.MemberID, req)
			s.ErrorIs(err, tc.ExpectedError, fmt.Sprintf("Error should be: %v but is actually: %v", tc.ExpectedError, err))
			if tc.ExpectedMember != nil {
				s.Equal(tc.ExpectedMember.Name, m.Name)
				s.Equal(tc.ExpectedMember.Age, m.Age)
			} else {
				s.Nil(m)
			}
		})
	}
}
