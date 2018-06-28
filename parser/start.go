package parser

import (
	"computorV1/exe"
	"errors"
	"fmt"
	"os"
	"strings"
)

/*
Start function
Entrypoint for parsing
*/
func Start() (*exe.Polynome, error) {
	args := os.Args
	if len(args) < 2 {
		return nil, errors.New(fmt.Sprintln("Not enough arguments"))
	}

	return fillPolynome(os.Args[1])
}

func fillPolynome(poly string) (p *exe.Polynome, err error) {
	p = new(exe.Polynome)
	var lefthand []string
	var righthand []string
	arr := strings.Fields(poly)
	for k, v := range arr {
		if v == "=" {
			lefthand = append(lefthand, arr[0:k]...)
			righthand = append(righthand, arr[k+1:]...)
		}
	}
	lefthand = makeOperators(lefthand)
	righthand = makeOperators(righthand)
	return exe.CreatePolynome(lefthand, righthand)
}

func makeOperators(hand []string) []string {
	for k, v := range hand {
		if len([]byte(v)) == 1 && []byte(v)[0] == '-' {
			hand[k+1] = fmt.Sprintf("%s%s", hand[k], hand[k+1])
			hand = append(hand[0:k], hand[k+1:]...)
			hand = makeOperators(hand)
			break
		}
	}
	return hand
}
