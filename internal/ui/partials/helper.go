package partials

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/a-h/templ"
)

func teamNameToBtnName(teamName string) string {
	// can i split it by space?
	splitted := strings.Split(teamName, " ")
	if len(splitted) > 1 {
		first := strings.ToUpper(string(splitted[0][0]))
		second := strings.ToUpper(string(splitted[1][0]))
		return fmt.Sprintf("%s%s", first, second)
	}

	// is it camel case
	// TODO: implement camel case detection

	// has it a number in it?
	re := regexp.MustCompile(`\d`)
	if re.MatchString(teamName) {
		int := re.FindString(teamName)
		return fmt.Sprintf("%s%v", strings.ToUpper(string(teamName[0])), int)
	}

	// Fallback use first two letters
	return fmt.Sprintf("%s%s", strings.ToUpper(string(teamName[0])), strings.ToUpper(string(teamName[1])))
}

type ButtonProps struct {
	Label       string
	Classes     string
	Type        string
	Hyperscript string
	HxAttr      *HXAttributes
}

func (p *ButtonProps) LabelToIcon() templ.Component {
	val, ok := IconMap[p.Label]
	if !ok {
		return placeholderIcon()
	}
	return val
}

type HXAttributes struct {
	Method     HxMethod
	Action     string
	Swap       string
	Trigger    string
	ReplaceURL string
	Target     string
}

type HxMethod string

var IconMap = map[string]templ.Component{
	"SkillAdd":     skillIcon(true),
	"SkillRemove":  skillIcon(false),
	"MemberAdd":    memberIcon(true),
	"MemberRemove": memberIcon(false),
}

const (
	GET               HxMethod = "get"
	POST              HxMethod = "post"
	DELETE            HxMethod = "delete"
	DefaultBtnClasses string   = "w-full h-full inline-flex justify-center items-center rounded-md bg-base px-3 py-2 text-sm font-semibold  hover:bg-crust m-1 transition-colors duration-200"
)

func (h *HXAttributes) GetAttributes() templ.Attributes {
	out := make(map[string]any)
	out["hx-trigger"] = h.Trigger
	out["hx-target"] = h.Target
	out["hx-swap"] = h.Swap
	out[fmt.Sprintf("hx-%s", h.Method)] = h.Action
	if h.ReplaceURL != "" {
		out["hx-replace-url"] = h.ReplaceURL
	}
	return out
}

func capitalizeFirst(in string) string {
	first := strings.ToUpper(string(in[0]))
	return fmt.Sprintf("%v%v", first, string(in[1:]))
}
