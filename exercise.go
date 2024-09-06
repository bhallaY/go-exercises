package main

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

// Given two integers, adds them together
//
// [Addition Explanation]: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](n1, n2 T) T {
	return n1 + n2
}

func main() {

}
