package main

import (
	"computorV1/parser"
	"fmt"
	"log"
)

func main() {
	p, err := parser.Start()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Parsed polynome: ")
		p.PrintPolynome()
	}
	p.Solve()
}
