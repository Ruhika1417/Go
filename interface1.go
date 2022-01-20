package main

import "fmt"

type Tank interface {
	Tarea() float64 //method signatures
	TVolume() float64
}

type Tarea struct {
	radius float64
	height float64
}

type TVolume struct {
	radius float64
	height float64
}

//implementing methods of tank interface and creating the receier type above
func (tank1 Tarea) Tarea() float64 {
	return 3.14 * tank1.height * tank1.radius
}
func (tank2 TVolume) Volume() float64 {  //can also name func as TVolume 
	return tank2.height * tank2.radius
}

func main() {
	//accessing elements of tank interface
	tank_area := Tarea{
		radius: 100.89,
		height: 98.6790,
	}

	tank_volume := TVolume{
		radius: 78.0,
		height: 1.0,
	}

	fmt.Printf("Tank Area is : %f", tank_area.Tarea())
	fmt.Printf("Tank Volume is : %f", tank_volume.Volume())
}
