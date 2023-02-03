package main

import (
	"fmt"
	"log"
)

func main() {
	//var s string = "a b c d"
	//s := "hello "
	//s := printString()
	s := deck{"a", "b", "c"}
	s = append(s, "d")
	log.Panic("adiu")
	s.printString()
}

type deck []string

func (d deck) printString() {
	for i, a := range d {
		fmt.Println(i, a)
		i = i + 1
	}
}
