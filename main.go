package main

import "fmt"

func main() {
	/* local variable definition */
	var a int = 100
	var b int = 200

	fmt.Printf("Before swap, value of a : %d\n", a )
	fmt.Printf("Before swap, value of b : %d\n", b )

	/* calling a function to swap the values */
	swap(a, b)

	fmt.Printf("After swap, value of a : %d\n", a )
	fmt.Printf("After swap, value of b : %d\n", b )
}

func multpleA(a string, b int) (string, int) {
	return a,b
}

/* function definition to swap the values */
func swap(x int, y int) int {
	var temp int

	temp = x /* save the value of x */
	x = y    /* put y into x */
	y = temp /* put temp into y */

	return temp
}