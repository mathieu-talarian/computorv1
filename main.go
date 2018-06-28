package main

import (
	"computorV1/parser"
	"fmt"
	"log"
)

func main() {
	if p, err := parser.Start(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Parsed polynome: ")
		p.PrintPolynome()
	}
	p.Solve()
}
