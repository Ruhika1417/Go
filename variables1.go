//GO FORMATTING VERBS
package main

import (
	"fmt"
)

func main() {
	var (
		name    string = "Spider-man"
		a, b, i        = 31, true, 151.854
	)

	fmt.Println("STRING FORMATTING VERBS")
	fmt.Printf("%s\n", name)
	fmt.Printf("%q\n", name)
	fmt.Printf("%8s\n", name)
	fmt.Printf("%-8s\n", name)
	fmt.Printf("%x\n", name)
	fmt.Printf("% x\n  \n", name)

	fmt.Println("INTEGER FORMATTING VERBS")
	fmt.Printf("%b\n %b\n %d\n %+d\n  %x\n  %X\n  \n", a, a, a, a, a, a)

	fmt.Println("FLOAT FORMATTING VERBS")
	fmt.Printf("%e\n %f\n .%.2f\n %6.2f %g\n \n ", i, i, i, i, i)

	fmt.Println("BOOLEAN FORMATTING VERBS")
	fmt.Printf("%t\n", b)
}
