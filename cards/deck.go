package main

import "fmt"

type deck []string

/* Receiver on a function */
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
