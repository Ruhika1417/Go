package main

import "fmt"

func main() {
	var b int = 10
	var a int //by default value of a is 0
	fmt.Println("value of a", a)
	numbers := [6]int{1, 2, 3, 5}

	/* for loop execution */
	for a := 0; a < 5; a++ {

		if a == 3 {
			continue
		}

		fmt.Printf("value of a is loop 1: %d\n", a)
	}

	fmt.Println("value of a is ", a)

	for a < b {
		a++
		fmt.Printf("value of a: %d\n", a)
	}
	for i, x := range numbers {
		fmt.Printf("value of x = %d at %d\n", x, i)
	}
}
