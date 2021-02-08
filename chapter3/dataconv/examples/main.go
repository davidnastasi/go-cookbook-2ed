package main

import "github.com/go-programing-cookbook-2ed/chapter3/dataconv"

func main() {
	dataconv.ShowConv()
	if err := dataconv.Strconv(); err != nil {
		panic(err)
	}
	dataconv.Iterfaces()
}
