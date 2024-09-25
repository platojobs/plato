package mymath

import (
	"fmt"
)
type geometry interface{
	area()float64
	perim()float64
}

type srect struct{
	width,height float64
}

type circle struct {
	radius float64
}

func (ss srect)area() float64{
	return ss.width * ss.height
}
func (ss srect)perim() float64{
	return 2*(ss.width + ss.height)
}
func (cc circle)area() float64{
	return cc.radius * cc.radius * 3.14
}
func (cc circle)perim() float64{
	return 2 * 3.14 * cc.radius
}	

func (cc *circle) sperim() float64{
	return 2 * 3.14 * cc.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}


func TTest(){
	fmt.Println("------1------")
    r := srect{width: 3, height: 4}
	c := circle{radius: 5}
	measure(r)
	fmt.Println("------2------")
	measure(c)
	
}