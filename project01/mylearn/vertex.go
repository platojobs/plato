package mylearn

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Vertt(){
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
	fmt.Println(v)
}


type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v 岁)", p.Name, p.Age)
}

func Personmain() {
	a := Person{"张三", 42}
	z := Person{"李四", 9001}
	fmt.Println(a, z)

	ip := IPAddr{127, 0, 0, 1}
	fmt.Println(ip.String())
     
	ErrNegativeSqrtmain()

}

type IPAddr [4]byte

func (a IPAddr)String() string{
	return fmt.Sprintf("%v.%v.%v.%v",a[0],a[1],a[2],a[3])
}



type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string{
     if e < 0 {
		return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
	 }else{
		 return fmt.Sprintf("%v",float64(e))
	 }
}


func ErrNegativeSqrtmain() {

	fmt.Println(ErrNegativeSqrt(-2).Error())

}