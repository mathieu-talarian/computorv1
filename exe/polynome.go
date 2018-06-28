package exe

import (
	"computorV1/tools"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var bitSize = 10

/*
Polynome struct
Super struct for polynome
considerating polynome => A + B + C = D + E + F
*/
type Polynome struct {
	Left  *Hand
	Right *Hand
}

type PolynomeII struct {
	A, B, C Monome
}

func (p *Polynome) PrintPolynome() {
	p.Left.PrintHand()
	fmt.Printf(" = ")
	p.Right.PrintHand()
	fmt.Println()
}

type Hand struct {
	Zero, One, Two *Monome
}

func (h *Hand) PrintHand() {
	h.Zero.PrintMonome(false)
	h.One.PrintMonome(true)
	h.Two.PrintMonome(true)
}

func (m *Monome) PrintMonome(b bool) {
	if m != nil {
		if b {
			if m.Operator < 0 {
				fmt.Printf("- ")
			} else {
				fmt.Printf("+ ")
			}
		}
		m.Val.Print()
		fmt.Printf("*")
		fmt.Printf(" X^%d ", m.Power)
	}

}

func impl(l Monomes) (h *Hand, err error) {
	h = &Hand{}
	if len(l) > 3 {
		return nil, tools.MyError("Too much values on left hand")
	}
	for k, v := range l {
		if v.Power == 0 {
			if v.Val.Zero() {
				h.Zero = &l[k]
			}
		} else if v.Power == 1 {
			if v.Val.Zero() {
				h.One = &l[k]
			}
		} else if v.Power == 2 {
			if v.Val.Zero() {
				h.Two = &l[k]
			}
		}
	}
	return h, nil
}

/*
Monome struct
A polynome is composed by one or more monomes
*/
type Monome struct {
	Val      Ints
	Power    int
	Operator int
}

type Ints []int

func (i Ints) Tofloat() float64 {
	if len(i) == 0 {
		return 0
	}
	if len(i) == 1 {
		return float64(i[0])
	}
	tmp := tools.Power(10, len(strconv.Itoa(i[1])))
	return float64(i[0]) + float64(i[1])/float64(tmp)
}

func (i Ints) Print() {
	if len(i) == 1 {
		fmt.Printf("%d ", i[0])
	} else {
		fmt.Printf("%d.%d ", i[0], i[1])
	}
}

func (i Ints) Zero() bool {
	return i.Tofloat() != float64(0)
}

/*
Monomes is an array of monome
*/
type Monomes []Monome

func findPower(power string) (p int, err error) {
	if []byte(power)[0] != 'X' {
		return 0, errors.New(fmt.Sprintln("Issue with power", power))
	}
	tmp := strings.Split(power, "^")
	if len(tmp) <= 1 {
		err = tools.MyError("Issue with power", power)
	}
	if tmp[0] == "X" {
		p, err = strconv.Atoi(tmp[1])
	} else {
		p, err = strconv.Atoi(tmp[0])
	}
	return
}

/*
CreateMonome func
*/
func createMonome(a ...string) (m Monome, err error) {
	var power, val string
	var tmpStr []string
	var tmpInt []int
	if []byte(a[0])[0] == 'X' {
		power = a[0]
		val = a[1]
	} else {
		power = a[1]
		val = a[0]
	}
	tmpStr = strings.Split(val, ".")
	for _, v := range tmpStr {
		var tmp int
		if tmp, err = strconv.Atoi(v); err != nil {
			return
		}
		tmpInt = append(tmpInt, tmp)
	}
	m.Val = tmpInt
	if m.Val[0] < 0 {
		m.Val[0], m.Operator = m.Val[0]*-1, -1
	} else {
		m.Operator = 1
	}
	if m.Power, err = findPower(power); err != nil {
		return
	}
	return
}

func createMonomes(hand []string) (m Monomes, err error) {
	for k, v := range hand {
		if v == "*" {
			var monome Monome
			if monome, err = createMonome(hand[k-1], hand[k+1]); err != nil {
				return
			}
			m = append(m, monome)
		}
	}
	return
}

func buildPolynome(l, r Monomes) (p *Polynome, err error) {
	p = &Polynome{
		&Hand{},
		&Hand{},
	}
	if p.Left, err = impl(l); err != nil {
		return nil, err
	}
	if p.Right, err = impl(r); err != nil {
		return nil, err
	}
	return
}

/*
CreatePolynome func
Way to create polynome from 2 strings from parser
*/
func CreatePolynome(lefthand, righthand []string) (p *Polynome, err error) {
	var rightMonomes, leftMonomes Monomes
	if rightMonomes, err = createMonomes(righthand); err != nil {
		return
	}
	if leftMonomes, err = createMonomes(lefthand); err != nil {
		return
	}
	return buildPolynome(leftMonomes, rightMonomes)
}
