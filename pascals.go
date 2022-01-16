package main

import "fmt"

func main() {

	var rows int = 7
	var temp int = 1

	//for row 1 i=0;row 2 i=1 ;row 3 i=2

	for i := 0; i < rows; i++ {
		for j := 1; j <= rows-i; j++ { //rows =7
			fmt.Print(" ")

		}
		//eg for row 2 i=1 so: k=0;k<=2;k++ will iterate 2 times
		for k := 0; k <= i; k++ {
			if k == 0 || i == 0 {
				temp = 1
			} else {
				temp = temp * (i - k + 1) / k
			}

			fmt.Printf("%d ", temp)

		}

		fmt.Println(" ")
	}
}
