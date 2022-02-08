package main

import "fmt"

func general_function(a int, b string) (dob int, name string) {
	dob = a
	name = b + "Bulani"
	return
}

func testcount(x int) int {

	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return testcount(x + 1) //recursion

}

//max of 2 no's
/*
func max(num1, num2 int) int{

	var  result int

	if( num1 > num2) {
		result = num1
	}
	else {
		result = num2
	}

	return result
}
*/
func swap(x, y string) (a string, b string) {

	// a = y
	// b = x
	// return a, b//how to implement with this
	fmt.Println(y, x)
	return y, x    
}

func main() {
	_, b := general_function(17, "Ruhika") //to omit the dob
	fmt.Println(b)

	//2
	testcount(1)

	// ret = max(18.0001, 18.002)
	// fmt.Printf("max value is:  %d\n", ret)

	//4
	swap("ruhika2", "bulani1")
}
