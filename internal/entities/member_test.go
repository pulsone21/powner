package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
	// TODO: Not sure how to set this up without having a actual db implementation
}

func TestMemberGetRatingBySkill(t *testing.T) {
	// TODO: Not sure how to set this up without having a actual db implementation
}
