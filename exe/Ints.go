package exe

import (
	"computorV1/tools"
	"fmt"
	"strconv"
)

type Ints []int

type iInts interface {
	ToNimber()
	Print()
	Zero()
}

func (i Ints) ToNumber() interface{} {
	if len(i) == 0 {
		return int(0)
	}
	if len(i) == 1 {
		return int(i[0])
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
	switch i.ToNumber().(type) {
	case int:
		return i.ToNumber().(int) != 0
	case float64:
		return i.ToNumber().(float64) != 0.0
	}
	return false
}
