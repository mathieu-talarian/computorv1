package exe

import (
	"computorV1/tools"
	"fmt"
)

type PolynomeII struct {
	A, B, C *Monome
	a, b, c float64
	Delta   float64
	Rac     interface{}
}

type IPolynomeII interface {
	Print()
	PrintDelta()
	PrintRac()
	Level1Poly() (err error)
	FindDelta() bool
	FindRac()
	twoRac() (ret []float64)
	oneRac() (ret float64)
	noneRac() (ret string)
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

func (p *PolynomeII) FindDelta() bool {
	fmt.Println(p.A)
	p.a = forDelta(p.A)
	p.b = forDelta(p.B)
	p.c = forDelta(p.C)
	if p.a == 0.0 {
		return false
	}
	fmt.Println(p.a, p.b, p.c)
	p.Delta = tools.Power(p.b, 2) - (4.0 * p.a * p.c)
	return true
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
	ret[0] = ((p.b + tools.Sqrt(p.Delta)) / (2 * p.a)) * -1
	ret[1] = ((p.b - tools.Sqrt(p.Delta)) / (2 * p.a)) * -1
	return
}

func (p *PolynomeII) oneRac() (ret float64) {
	return (p.b * -1) / (2 * p.a)
}

func (p *PolynomeII) noneRac() (ret string) {
	return "This polynome does not have solution on ℝ"
}
