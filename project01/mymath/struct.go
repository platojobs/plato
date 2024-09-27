package mymath

import (
    "fmt"
	"encoding/json"
)

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


type User struct {
	Name string `json:"username"`
	Age int     `json:"userage"`
	Salary int  `json:"usersalary"`
}


func TTuser(){

	myself := User{
		Name:   "唐春",
		Age:    22,
		Salary: 10,
	}
	// json.Marshal()方法作用就是把结构体转换为json,对应的字段名为标签对应的值
	jsdata,err := json.Marshal(myself)
	if err != nil {
		fmt.Println("格式错误")
	} else {
		fmt.Printf("User结构体转json:%s\n",jsdata)
	}
}

