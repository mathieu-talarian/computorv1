package exe

import (
	"computorV1/tools"
	"fmt"
	"strconv"
)

type Ints []int

type IInts interface {
	ToFloat()
	Print()
	Zero()
}

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
