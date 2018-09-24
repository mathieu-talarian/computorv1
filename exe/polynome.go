package exe

import (
	"computorV1/exe/tools"
	"errors"
	"fmt"
	"math"
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

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func newVal(l, r *Monome) (m Monome) {
	fmt.Println(l.Val, r.Val)
	fmt.Println(float64(r.Operator))
	m.ValF = l.Val.Tofloat()*float64(l.Operator) + r.Val.Tofloat()*float64(r.Operator)*-1.0
	m.ValF = toFixed(m.ValF, 2)
	m.Operator = 1
	if m.ValF < 0.000000 {
		m.Operator = -1
	}
	m.Power = l.Power
	return
}

func newHand(l, r *Monome) *Monome {
	if l != nil {
		if r != nil {
			m := newVal(l, r)
			return &m
		}
		return l
	}
	if r != nil {
		return r
	}
	return nil
}

func (p *Polynome) PrintReducedValue() {
	fmt.Printf("Reduced form => ")
	p.Left.Two = newHand(p.Left.Two, p.Right.Two)
	p.Left.One = newHand(p.Left.One, p.Right.One)
	p.Left.Zero = newHand(p.Left.Zero, p.Right.Zero)
	p.Left.PrintHand2()
}

func (p *Polynome) PrintPolynome() {
	p.Left.PrintHand()
	fmt.Printf(" = ")
	p.Right.PrintHand()
	fmt.Println()
	p.PrintReducedValue()

}

type Hand struct {
	Zero, One, Two *Monome
}

func (h *Hand) PrintHand() {
	h.Two.PrintMonome()
	h.One.PrintMonome()
	h.Zero.PrintMonome()
}

func (h *Hand) PrintHand2() {
	h.Two.PrintMonomeF()
	h.One.PrintMonomeF()
	h.Zero.PrintMonomeF()

}

func (m *Monome) PrintMonomeF() {
	if m != nil {
		fmt.Printf("%c%.2f * X^%d", m.Operator, m.ValF, m.Power)
	}
}

func (m *Monome) PrintMonome() {
	if m != nil {
		m.Val.Print(m.Operator)
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
	ValF     float64
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

func (i Ints) Print(o int) {
	if len(i) == 1 {
		fmt.Printf("%d ", i[0]*o)
	} else {
		fmt.Printf("%d.%d ", i[0]*o, i[1])
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
