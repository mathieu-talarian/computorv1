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

func Power(x, y int) int {
	if y > 0 {
		return x * Power(x, y-1)
	}
	return 1

}
