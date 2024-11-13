package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeamHasChanges(t *testing.T) {
	testCases := []struct {
		Team           Team
		NewName        string
		NewDescription string
		NewSkills      []Skill
		NewMembers     []Member
		Expected       bool
	}{
		{
			Team:           *NewTeam("TestTeam", "TestDesc"),
			NewName:        "TestTeam",
			NewDescription: "TestDesc",
			NewSkills:      []Skill{},
			NewMembers:     []Member{},
			Expected:       false,
		},
		{
			Team:           *NewTeam("TestTeam", "TestDesc"),
			NewName:        "TestTeam",
			NewDescription: "TestDesc1",
			NewSkills:      []Skill{},
			NewMembers:     []Member{},
			Expected:       true,
		},
		{
			Team:           *NewTeam("TestTeam", "TestDesc"),
			NewName:        "TestTeam",
			NewDescription: "TestDesc",
			NewSkills:      []Skill{*NewSkill("Test", "TestSkill", SkillType(0), 1)},
			NewMembers:     []Member{},
			Expected:       true,
		},
		{
			Team:           *NewTeam("TestTeam", "TestDesc"),
			NewName:        "TestTeam",
			NewDescription: "TestDesc",
			NewSkills:      []Skill{},
			NewMembers:     []Member{*NewMember("Hans", 24)},
			Expected:       true,
		},
	}

	for _, tC := range testCases {
		nT, changes := tC.Team.HasChanges(tC.NewName, tC.NewDescription, &tC.NewSkills, &tC.NewMembers)
		assert.Equal(t, tC.Expected, changes)
		assert.Equal(t, tC.NewName, nT.Name)
		assert.Equal(t, tC.NewDescription, nT.Description)
		assert.Equal(t, tC.NewSkills, nT.Skills)
		assert.Equal(t, tC.NewMembers, nT.Members)
	}
}
