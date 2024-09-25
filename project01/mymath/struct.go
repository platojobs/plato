package mymath

import "fmt"

type rect struct{
	width int
	height int
}

// 面积
func (r *rect) area() int {
	return r.width * r.height
}

// 周长
func (r rect) perim() int {
	return 2*(r.width + r.height)
}


func MMRest(){
	r := rect{width: 10, height: 5}
	fmt.Println("面积: ", r.area())
	fmt.Println("周长:", r.perim())

	rp := &r
	fmt.Println("面积: ", rp.area())
	fmt.Println("周长:", rp.perim())
	
}
