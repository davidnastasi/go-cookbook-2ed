package main

import (
	"fmt"
	"github.com/go-programing-cookbook-2ed/chapter3/math"
)

func main() {
	math.Examples()
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", math.Fib(i))
	}
	fmt.Println()

}
