package components

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/pulsone21/powner/internal/entities"
	"github.com/pulsone21/powner/internal/ui"
)

type (
	TeamListAddon   func(teamID string) templ.Component
	MemberItemAddon func(m entities.Member) templ.Component
)

type MemberComponent func(m entities.Member) templ.Component

type CSSMergerProp interface {
	MergeClasses() string
}

type ListItemProps struct {
	Class      string
	Header     string
	Footer     string
	ParentAttr templ.Attributes
}

func (p ListItemProps) MergeClasses() string {
	defaults := "min-h-10 justify-between flex flex-row items-center rounded-lg p-1 hover:opacity-70 transition"
	out := ui.CssMerger.Merge(fmt.Sprint(defaults + " " + p.Class))
	fmt.Println(out)
	return out
}

type ListProps struct {
	Class    string
	ListAttr templ.Attributes
}

func (p ListProps) MergeClasses() string {
	defaults := "w-full scroll-smooth h-full overflow-y-auto divide-y-2 divide-base flex flex-col gap-2"
	return ui.CssMerger.Merge(fmt.Sprint(defaults + " " + p.Class))
}
