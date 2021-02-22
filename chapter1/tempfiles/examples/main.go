package main

import "github.com/davidnastasi/go-programming-cookbook-2ed/chapter1/tempfiles"

func main() {
	if err := tempfiles.WorkWithTemp(); err != nil {
		panic(err)
	}
}
