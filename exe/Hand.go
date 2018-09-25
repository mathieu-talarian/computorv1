package exe

type Hand struct {
	Zero, One, Two *Monome
}

type IHand interface {
	PrintHand()
}

func (h *Hand) PrintHand() {
	h.Zero.PrintMonome(false)
	h.One.PrintMonome(true)
	h.Two.PrintMonome(true)
}
