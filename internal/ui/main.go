package ui

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	merge "github.com/tylantz/go-tailwind-merge"
)

var CssMerger = createMerger()

//go:emebed ./public/static/styles/base.css
var css string

func createMerger() *merge.Merger {
	cssFile, err := os.ReadFile("./public/static/styles/base.css")
	if err != nil {
		log.Fatalln(errors.Join(fmt.Errorf("Couldn't read base.css file"), err))
	}
	cssString := string(cssFile)
	fmt.Println("CSS Merger is called")
	m := merge.NewMerger(nil, true)
	m.AddRules(strings.NewReader(cssString), false)
	return m
}
