package main

import "fmt"

type shape interface{ getArea() float64 }

type triangle struct{
	height float64
	base float64
}
type square struct{
	sideLength float64
}

func main() {

	sq := square{2.5}
	tri := triangle{3.0, 3.0}


	printArea(sq)
	printArea(tri)

}

func (s square) getArea() float64{
	
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64{ 
	return t.height * t.base * 0.5
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}