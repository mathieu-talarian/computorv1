package main

import (
	"computorV1/parser"
	"log"
)

func main() {
	if p, err := parser.Start(); err != nil {
		log.Fatal(err)
	} else {
		p.PrintPolynome()
	}
}
