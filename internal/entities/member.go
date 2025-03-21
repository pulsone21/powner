package entities

import (
	"net/mail"
	"time"

	"github.com/pulsone21/powner/internal/errx"
)

type Member struct {
	ID        uint          `json:"id"`
	Firstname string        `json:"firstname"`
	Lastname  string        `json:"lastname"`
	Email     string        `json:"email"`
	Birthday  time.Time     `json:"birthday"`
	Skills    []SkillRating `json:"skills"`
}

func NewMember(fName, lName, email string, bDay time.Time) *Member {
	return &Member{
		Firstname: fName,
		Lastname:  lName,
		Email:     email,
		Birthday:  bDay,
		Skills:    []SkillRating{},
	}
}

func (m Member) HasSkill(skillID uint) bool {
	for _, sR := range m.Skills {
		if sR.Skill.ID == skillID {
			return true
		}
	}
	return false
}

func (m Member) GetType() string {
	return "member"
}

func (m Member) GetID() uint {
	return m.ID
}

func (m Member) GetSkillRatingBySkill(id uint) *SkillRating {
	for _, sR := range m.Skills {
		if sR.Skill.ID == id {
			return &sR
		}
	}
	return nil
}

func (m *Member) HasChanges(newM *Member) bool {
	return m != newM
}

type memberSort []Member

func (s memberSort) Len() int           { return len(s) }
func (s memberSort) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s memberSort) Less(i, j int) bool { return s[i].ID > s[j].ID }
func (s memberSort) toMember() []Member { return []Member(s) }

type MemberRequest struct {
	Firstname string `json:"name" schema:"name"`
	Lastname  string `json:"lastname"`
	Birthday  int    `json:"age" schema:"age"`
	Email     string `json:"email"`
}

func (m MemberRequest) ValidateFields() errx.ErrorMap {
	var validationErr errx.ErrorMap

	now := time.Now()
	bDay := time.Unix(int64(m.Birthday), 0)

	diff := bDay.Sub(now)

	if (diff.Hours() / 24 / 365) < 16 {
		validationErr.Set("age", "Age must be bigger then 16, no kids labor allowed in here...")
	}

	if len(m.Firstname) < 1 {
		validationErr.Set("firstname", "Fastname must be >= 1 chars")
	}

	if len(m.Lastname) < 1 {
		validationErr.Set("lastname", "Lastname must be >= 1 chars")
	}

	if _, err := mail.ParseAddress(m.Email); err != nil {
		validationErr.Set("email", err.Error())
	}

	return validationErr
}
