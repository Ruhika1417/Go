//maps are unordered
package main

import "fmt"

func main() {
	//var =map[key type]value type{ key: value pairs}
	var a = map[string]int{"ruhika": 17, "megha": 03, "reshma": 8, "ruhika2": 0}
	a["reshma"] = 89
	delete(a, "megha")
	val, ok := a["ruhika"]
	//ref_var := a
	//fmt.Println(ref_var)

	fmt.Printf("a\t%v\n", a) // \t is the tab space
	fmt.Println(val, ok)

	//method 2 using make
	var b = make(map[string]int)
	b["ruhika"] = 17
	b["megha"] = 03
	b["reshma"] = 8

	fmt.Printf("b\t%v\n", b)

	//creating an empty map in 2 ways make()is the best way

	var c1 map[string]int
	var c2 = make(map[string]int)

	fmt.Println(c1 == nil)
	fmt.Println(c2 == nil)

	//iterating over maps

	for k, v := range a {
		fmt.Printf(" %v : %v, ", k, v)
	}

}
