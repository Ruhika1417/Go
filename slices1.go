//creating slices in go in 3 ways
package main

import (
	"fmt"
)

func main() {

	var array1 = [...]int{2, 3, 4}
	//method 1: from an array
	my_slice := array1[0:]
	my_slice = array1[1:] //changing length
	fmt.Printf("myslice = %v\n", my_slice)
	fmt.Printf("length = %d\n", len(my_slice))
	fmt.Printf("capacity = %d\n", cap(my_slice))

	//method2: default format
	my_slice1 := []int{1, 2, 3, 4, 5, 6, 7}
	my_slice1 = append(my_slice1, 8, 9) //appending

	//method3 : using make()
	my_slice2 := make([]int, 3, 6)

	my_slice3 := append(my_slice, my_slice2...) //appending

	//copy function to get useful slice int

	my_slice111 := make([]int, len(my_slice1)-5)
	copy(my_slice111, my_slice1)
	fmt.Println(my_slice111)

	fmt.Print(my_slice, "\n", my_slice1, "\n", my_slice2, "\n", my_slice3)
}
