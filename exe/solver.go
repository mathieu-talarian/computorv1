package exe

import (
	"computorV1/tools"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func a(l, r *Monome) (m *Monome, err error) {
	m = &Monome{}
	if l != nil && r != nil {
		return retTwo(l, r)
	} else if l != nil && r == nil {
		return retLeft(l)
	} else if l == nil && r != nil {
		return retRight(r)
	} else if l == nil && r == nil {
		return nil, nil
	}
	return nil, tools.MyError("Issue while converting Initial polynome to reduced polynome")
}

func retTwo(l, r *Monome) (m *Monome, err error) {
	m = &Monome{}
	tmp := l.Val.Tofloat()*float64(l.Operator) + r.Val.Tofloat()*float64(r.Operator*-1)
	fmt.Println(tmp)
	if tmp < 0.0 {
		m.Operator = -1
	} else {
		m.Operator = 1
	}
	if m.Val, err = floatToInts(tmp); err != nil {
		return
	}
	m.Power = l.Power
	return
}

func floatToInts(f float64) (i Ints, err error) {
	i = []int{0, 0}
	a := strings.Split(fmt.Sprintf("%.2f", f), ".")
	if len(a) >= 1 {
		var tmp1 int
		if tmp1, err = strconv.Atoi(string(a[0])); err != nil {
			return
		}
		i[0] = tmp1
		if len(a) == 2 {
			var tmp2 int
			t := strings.Trim(a[1], "\n")
			if tmp2, err = strconv.Atoi(t); err != nil {
				return
			}
			i[1] = tmp2
		}
	}
	return
}

func retLeft(l *Monome) (m *Monome, err error) {
	if l == nil {
		return nil, tools.MyError("Issue with pointer")
	}
	return l, nil
}

func retRight(r *Monome) (m *Monome, err error) {
	if r == nil {
		return nil, tools.MyError("Issue with pointer")
	}
	m = r
	m.Operator = m.Operator * -1
	return
}

func forDelta(m *Monome) (f float64) {
	if m != nil {
		return m.Val.Tofloat() * float64(m.Operator)
	}
	return 0.0
}

func impl(l Monomes) (h Hand, err error) {
	h = Hand{}
	if len(l) > 3 {
		return h, tools.MyError("Too much values on left hand")
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
		Hand{},
		Hand{},
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
