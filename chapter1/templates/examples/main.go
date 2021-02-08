package main

import (
	"github.com/go-programing-cookbook-2ed/chapter1/templates"
)

func main() {
	if err := templates.RunTemplate(); err != nil {
		panic(err)
	}
	if err := templates.InitTemplates(); err != nil {
		panic(err)
	}
	if err := templates.HTMLDifferences(); err != nil {
		panic(err)
	}

}