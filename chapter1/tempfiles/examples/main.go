package main

import "github.com/go-programing-cookbook-2ed/chapter1/tempfiles"

func main() {
	if err := tempfiles.WorkWithTemp(); err != nil {
		panic(err)
	}
}
