package shared

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/a-h/templ"
)

func LabelToIcon(label string) templ.Component {
	val, ok := iconMap[strings.ToLower(label)]
	if !ok {
		return placeholderIcon()
	}
	return val
}

var iconMap = map[string]templ.Component{
	"skilladd":     skillAddIcon(true),
	"skillremove":  skillAddIcon(false),
	"memberadd":    memberAddIcon(true),
	"memberremove": memberAddIcon(false),
	"members":      memberIcon(),
	"teams":        teamIcon(),
	"skills":       skillIcon(),
	"settings":     settingsIcon(),
}

func TeamNameToBtnName(teamName string) string {
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

const (
	DefaultBtnClasses string = "w-full h-full inline-flex justify-center items-center rounded-md bg-base px-3 py-2 text-sm font-semibold  hover:bg-crust m-1 transition-colors duration-200"
)

func CapitalizeFirst(in string) string {
	first := strings.ToUpper(string(in[0]))
	return fmt.Sprintf("%v%v", first, string(in[1:]))
}
