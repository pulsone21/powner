package entities

import (
	"time"

	"github.com/pulsone21/powner/internal/database/pwd"
)

type User struct {
	Firstname string
	Lastname  string
	Email     string
	Birthday  time.Time
	Password  string
	Type      UserType
}

type UserType int

const (
	TeamLeader UserType = 0
	TeamMember UserType = 1
)

func NewUser(fName, lName, email string, password pwd.HashSalt, bDay time.Time, t UserType) *User {
	return &User{
		Firstname: fName,
		Lastname:  lName,
		Email:     email,
		Password:  password.String(),
		Type:      t,
	}
}
