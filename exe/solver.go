package exe

import (
	"computorV1/tools"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type PolynomeII struct {
	A, B, C *Monome
	a, b, c float64
	Delta   float64
	Rac     interface{}
}

func (p *PolynomeII) Print() {
	p.A.PrintMonome(false)
	p.B.PrintMonome(true)
	p.C.PrintMonome(true)
	fmt.Println("= 0")
}

func (p *PolynomeII) PrintDelta() {
	fmt.Println("Le discriminant est :", p.Delta)
}

func (p *PolynomeII) PrintRac() {
	if s, ok := p.Rac.([]float64); ok {
		fmt.Println("L'equation a deux solutions dans ℝ:", s[0], s[1])
	} else if s, ok := p.Rac.(float64); ok {
		fmt.Println("Cette equation a une solution dans ℝ: ", s)
	} else if s, ok := p.Rac.(string); ok {
		fmt.Println(s)
	}
}

/*
Solve func
Starting point for solver
*/
func (p *Polynome) Solve() (err error) {
	fmt.Printf("%+v\n", p)
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

func (p *PolynomeII) Level1Poly() (err error) {
	if p.b != 0.0 {
		fmt.Println("ceci est un polynome de degre 1")
		if float64(int(p.c/p.b)) != p.c/p.b {
			fmt.Println("Le resultat de l'equation : \nx =", p.c, "/", p.b)
		} else {
			fmt.Println("Le resultat de l'equation : \nx =", p.c/p.b)

		}
	} else if p.b == 0.0 {
		fmt.Println("cette `equation` n'a pas d'inconnues")
	}
	return
}

/*
ToPolyII func
Switch to reduced polynome
*/
func (p *Polynome) ToPolyII() (y *PolynomeII, err error) {
	y = new(PolynomeII)
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
	return
}

func floatToInts(f float64) (i Ints, err error) {
	i = []int{0, 0}
	a := strings.Split(fmt.Sprintln(f), ".")
	if len(a) >= 1 {
		var tmp1 int
		if tmp1, err = strconv.Atoi(a[0]); err != nil {
			return
		}
		i[0] = tmp1
		if len(a) == 2 {
			var tmp2 int
			t := strings.Trim(a[1], "\n")
			fmt.Println(t)
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

func (p *PolynomeII) FindDelta() bool {
	p.a = forDelta(p.A)
	p.b = forDelta(p.B)
	p.c = forDelta(p.C)
	if p.a == 0.0 {
		return false
	}
	p.Delta = tools.Power(p.b, 2) - 4.0*p.a*p.c
	return true
}

func forDelta(m *Monome) (f float64) {
	if m != nil {
		return m.Val.Tofloat()
	}
	return 0.0
}

func (p *PolynomeII) FindRac() {
	if p.Delta > 0 {
		p.Rac = p.twoRac()
	} else if p.Delta == 0 {
		p.Rac = p.oneRac()
	} else {
		p.Rac = p.noneRac()
	}
}

func (p *PolynomeII) twoRac() (ret []float64) {
	ret = make([]float64, 2)
	ret[0] = (p.b*-1 + tools.Sqrt(p.Delta)) / 2 * p.a
	ret[1] = (p.b*-1 - tools.Sqrt(p.Delta)) / 2 * p.a
	return
}

func (p *PolynomeII) oneRac() (ret float64) {
	return (p.b * -1) / (2 * p.a)
}

func (p *PolynomeII) noneRac() (ret string) {
	return "This polynome does not have solution on ℝ"
}
