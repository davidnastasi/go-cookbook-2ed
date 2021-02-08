package main

import "github.com/go-programing-cookbook-2ed/chapter2/confformat"

func main() {
	if err := confformat.MarshalAll(); err != nil {
		panic(err)
	}
	if err := confformat.UnmarshalAll(); err != nil {
		panic(err)
	}
	if err := confformat.OtherJSONExamples(); err != nil {
		panic(err)
	}
}