package components

import (
	"github.com/a-h/templ"
	"github.com/pulsone21/powner/internal/entities"
)

type (
	TeamListAddon   func(teamID string) templ.Component
	MemberItemAddon func(m entities.Member) templ.Component
)

type MemberComponent func(m entities.Member) templ.Component
