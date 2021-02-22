package main

import "github.com/davidnastasi/go-programing-cookbook-2ed/chapter1/tempfiles"

func main() {
	if err := tempfiles.WorkWithTemp(); err != nil {
		panic(err)
	}
}
