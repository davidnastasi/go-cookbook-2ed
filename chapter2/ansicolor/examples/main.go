package main

import (
	"fmt"
	"github.com/go-programing-cookbook-2ed/chapter2/ansicolor"
)

func main() {
	r := ansicolor.ColorText {
		TextColor : ansicolor.Red,
		Text: "I'm red!",
	}
	fmt.Println(r.String())

	r = ansicolor.ColorText {
		TextColor : ansicolor.Green,
		Text: "Now I'm green!",
	}
	fmt.Println(r.String())

	r = ansicolor.ColorText {
		TextColor : ansicolor.None,
		Text: "Back to normal...!",
	}
	fmt.Println(r.String())
}
