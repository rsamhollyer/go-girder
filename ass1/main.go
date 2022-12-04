/*
 * Create a new package
 * Creat a slice of ints from 0 - 10
 * Iterate through slice and check if the value is even or odd
 * Print that value
 */

package main

import "fmt"

type numbers []int

func (n numbers) print() {
	for _, number := range n {
		if number%2 == 0 {
			fmt.Println("The number", number, ",is Even")
		} else {
			fmt.Println("The number", number, ",is Odd")
		}
	}
}

func main() {
	numbers := numbers{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numbers.print()
}
