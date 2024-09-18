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

}
