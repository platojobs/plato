package mylearn

import "fmt"

type carInMotion interface {
	move(speed int)
	brake() (int, int)
	park()
	consumeOil(flueleft float64, aver_construction float64) (distance float64)
}

type Car struct {
	color string
}

func (c *Car) move(speed int) {
	fmt.Println("car is moving at speed", speed)
}
func (c *Car) brake() (int, int) {
	fmt.Printf("car is braking %v\n", c.color)
	speedBefore := 60
	speedAfter := 0
	return speedBefore, speedAfter
}

func (c *Car) park() {
	fmt.Println("car is parking%v \n", c.color)

}

func (c *Car) consumeOil(flueleft float64, aver_construction float64) (distance float64) {
	fmt.Printf("%v车还剩余燃油%v,燃油平均值%v,还能行驶%v\n", c.color, flueleft, aver_construction, distance)
	return flueleft / aver_construction * 100
}



func CarTest() {
	var cim carInMotion
	cr := Car{"红色"}
	cim = &cr
	cim.move(60)
	speedBefore, speedAfter := cim.brake()
	fmt.Printf("speed before=%v,speedAfter=%v\n", speedBefore, speedAfter)
	cim.brake()
	distance := cim.consumeOil(27, 6.3)
	fmt.Println(distance)



	var a,b interface{}
	a = 711
	b = "hello"
	value_a, a_ok := a.(int)
	fmt.Printf("value_a=%v,a_ok = %v",value_a,a_ok) //value_a=711,a_ok = true
	
    value_b , b_ok := b.(string)
	fmt.Printf("value_b=%v,b_ok = %v",value_b,b_ok) //value_b=hello,b_ok = true


	var r runner
	pn := person{
		name: "zhangsan",
		age:  18,
		legs: 2,
	}
	r = &pn
	r.run()
	value_pn , ok_pn := r.(*person)
	fmt.Printf("value_pn=%v,ok_pn=%v",value_pn,ok_pn)
    ct := Car{"red"}
	r = &ct
    value_ct,ok_ct := r.(*person)
    fmt.Printf("value_ct=%v,ok_ct=%v",value_ct,ok_ct)

}





type runner interface {
	run()
}

type person struct {
	name string
	age int
	legs int
}

func (p *person) run() {
	fmt.Printf("%v is running\n,has %v legs age = %v",p.name,p.legs,p.age)
}

func(c *Car)run(){
	fmt.Println("%v is running")
}