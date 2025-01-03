package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestMemberHasChanges(t *testing.T) {
	testCases := []struct {
		Name      string
		OldMember Member
		NewName   string
		NewAge    int
		Expected  bool
	}{
		{
			Name:      "No Changes",
			OldMember: *NewMember("Peter", 60),
			NewName:   "Peter",
			NewAge:    60,
			Expected:  false,
		},
		{
			Name:      "Changes in Name",
			OldMember: *NewMember("Peter", 60),
			NewName:   "James",
			NewAge:    60,
			Expected:  true,
		},
		{
			Name:      "Changes in Age",
			OldMember: *NewMember("Peter", 60),
			NewName:   "Peter",
			NewAge:    20,
			Expected:  true,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.Name, func(t *testing.T) {
			newMember, changes := tC.OldMember.HasChanges(tC.NewName, tC.NewAge)

			assert.Equal(t, changes, tC.Expected, "Should be equal")
			assert.Equal(t, newMember.Name, tC.NewName, "Should be equal")
			assert.Equal(t, newMember.Age, tC.NewAge, "Should be equal")
		})
	}
}

func TestMemberHasSkill(t *testing.T) {
	testCases := []struct {
		Name     string
		Member   Member
		Expected bool
		SkillID  uint
	}{
		{
			Name: "Has Skill",
			Member: Member{
				Skills: []SkillRating{
					{
						SkillID: 1,
						Skill: Skill{
							Model: gorm.Model{ID: 1},
						},
					},
				},
			},
			Expected: true,
			SkillID:  uint(1),
		},
		{
			Name:     "Skill is missing",
			Member:   *NewMember("Test", 25),
			Expected: false,
			SkillID:  uint(1),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.Name, func(t *testing.T) {
			actual := tC.Member.HasSkill(tC.SkillID)
			assert.Equal(t, tC.Expected, actual)
		})
	}
}

func TestMemberGetRatingBySkill(t *testing.T) {
	testCases := []struct {
		Name     string
		Member   Member
		SkillID  uint
		Expected *SkillRating
	}{
		{
			Name: "Found skillrating by skill",
			Member: Member{
				Skills: []SkillRating{
					{
						SkillID: 1,
						Rating:  2,
						Skill: Skill{
							Model: gorm.Model{ID: 1},
							Name:  "TestSkill",
						},
					},
				},
			},
			SkillID: uint(1),
			Expected: &SkillRating{
				SkillID: 1,
				Rating:  2,
				Skill: Skill{
					Model: gorm.Model{ID: 1},
					Name:  "TestSkill",
				},
			},
		},
		{
			Name: "Didn't found skillrating by skill",
			Member: Member{
				Skills: []SkillRating{},
			},
			SkillID:  uint(1),
			Expected: nil,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.Name, func(t *testing.T) {
			actual := tC.Member.GetSkillRatingBySkill(tC.SkillID)
			assert.Equal(t, tC.Expected, actual)
		})
	}
}
