package exe

import (
	"fmt"
	"log"
)

var bitSize = 10

/*
Polynome struct
Super struct for polynome
considerating polynome => A + B + C = D + E + F
*/
type Polynome struct {
	Left  Hand
	Right Hand
}

type IPolynome interface {
	PrintPolynome()
	Solve() (err error)
	ToPolyII() (y *PolynomeII, err error)
}

func (p *Polynome) PrintPolynome() {
	p.Left.PrintHand()
	fmt.Printf(" = ")
	p.Right.PrintHand()
	fmt.Println()
}

/*
Solve func
Starting point for solver
*/
func (p *Polynome) Solve() (err error) {
	polyII, err := p.ToPolyII()
	if err != nil {
		log.Fatal(err)
		return
	}
	polyII.Print()
	if polyII.FindDelta() {
		polyII.PrintDelta()
		polyII.FindRac()
		polyII.PrintRac()
	} else {
		return polyII.Level1Poly()
	}
	return
}

/*
ToPolyII func
Switch to reduced polynome
*/
func (p *Polynome) ToPolyII() (y *PolynomeII, err error) {
	y = NewPolyII()
	if y.A, err = a(p.Left.Two, p.Right.Two); err != nil {
		return nil, err
	}
	if y.B, err = a(p.Left.One, p.Right.One); err != nil {
		return nil, err
	}
	if y.C, err = a(p.Left.Zero, p.Right.Zero); err != nil {
		return nil, err
	}
	return
}
