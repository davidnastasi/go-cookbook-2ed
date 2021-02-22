package main

import (
	"fmt"
	"github.com/davidnastasi/go-programming-cookbook-2ed/chapter3/math"
)

func main() {
	math.Examples()
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", math.Fib(i))
	}
	fmt.Println()

}
