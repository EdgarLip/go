package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}
type shape interface {        // create new type "shape" which implemets 2 methods
	area() float64            //   which ever struct implemets those methods
	areaWithPointer() float64 // he will be also ( bbecomes also ) of type "shape".
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) areaWithPointer() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area", s.area())
	//fmt.Println("area", s.areaWithPointer())
}

func main() {
	c := circle{5}
	info(c) // not wroking
	//info(&c) // working !
}
