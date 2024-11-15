package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestMemberHasChanges(t *testing.T) {
	testCases := []struct {
		OldMember Member
		NewName   string
		NewAge    int
		Expected  bool
	}{
		{
			OldMember: *NewMember("Peter", 60),
			NewName:   "Peter",
			NewAge:    60,
			Expected:  false,
		},
		{
			OldMember: *NewMember("Peter", 60),
			NewName:   "James",
			NewAge:    60,
			Expected:  true,
		},
		{
			OldMember: *NewMember("Peter", 60),
			NewName:   "Peter",
			NewAge:    20,
			Expected:  true,
		},
	}

	for _, tC := range testCases {
		newMember, changes := tC.OldMember.HasChanges(tC.NewName, tC.NewAge)

		assert.Equal(t, changes, tC.Expected, "Should be equal")
		assert.Equal(t, newMember.Name, tC.NewName, "Should be equal")
		assert.Equal(t, newMember.Age, tC.NewAge, "Should be equal")
	}
}

func TestMemberHasSkill(t *testing.T) {
	testCases := []struct {
		Member   Member
		Expected bool
		SkillID  uint
	}{
		{
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
			Member:   *NewMember("Test", 25),
			Expected: false,
			SkillID:  uint(1),
		},
	}
	for _, tC := range testCases {
		actual := tC.Member.HasSkill(tC.SkillID)
		assert.Equal(t, tC.Expected, actual)
	}
}

func TestMemberGetRatingBySkill(t *testing.T) {
	testCases := []struct {
		Member   Member
		SkillID  uint
		Expected *SkillRating
	}{
		{
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
			Member: Member{
				Skills: []SkillRating{},
			},
			SkillID:  uint(1),
			Expected: nil,
		},
	}

	for _, tC := range testCases {
		actual := tC.Member.GetSkillRatingBySkill(tC.SkillID)
		assert.Equal(t, tC.Expected, actual)
	}
}
