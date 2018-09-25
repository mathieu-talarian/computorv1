package exe

import "fmt"

/*
Monome struct
A polynome is composed by one or more monomes
*/
type Monome struct {
	Val      Ints
	Power    int
	Operator int
}

/*
Monomes is an array of monome
*/
type Monomes []Monome

/*
IMonome interface for struct monome
*/
type IMonome interface {
	PrintMonome(b bool)
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
