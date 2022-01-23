package main

import (
	"fmt"
	"strings"
)

func main() {
	for i := 5; i >= 1; i-- {
		fmt.Printf("%5s\n", strings.Repeat("#", i)) //%6s left justified width 6
	}
	fmt.Println()

	for i := 1; i <= 5; i++ {
		fmt.Printf("%5s\n", strings.Repeat("#", i))
		//%6s left justified width 6
	}
	fmt.Println()

	for i := 1; i <= 5; i++ {
		fmt.Printf("%s\n", strings.Repeat("#", i))
	}

	//half triangle pattern with numbers

	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", j)
		}
		fmt.Println()
	}

}
