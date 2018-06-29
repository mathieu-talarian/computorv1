package tools

import (
	"errors"
	"fmt"
)

/*
MyError function
just a tool for writing errors
*/
func MyError(str ...string) error {
	return errors.New(fmt.Sprintln(str))
}

func Power(x float64, y int) float64 {
	if y > 0 {
		return x * Power(x, y-1)
	}
	return 1
}

/*
DELTA const
For SQRT func
*/
const DELTA = 0.000000000001

/*
Sqrt func
Find root square
Babylonian method
https://en.wikipedia.org/wiki/Methods_of_computing_square_roots#Babylonian_method
*/
func Sqrt(x float64) float64 {
	if x <= 0 {
		return 0
	}
	s := x
	for (s - x/s) > DELTA {
		s = (s + x/s) / 2
	}
	return s
}
