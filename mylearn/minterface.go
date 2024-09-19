package mylearn

import (
	"fmt"
)

type notfouctions  interface{}

func Minterface(){

	var nf notfouctions
	fmt.Printf("空接口=%v,空接口的类型：%T\n",nf,nf)
    nfs := 711
	nf = nfs
	judeType(nf)
	fmt.Printf("空接口保存的数据1=%v,空接口保存1的类型：%T\n",nf,nf)
	str := "hello"
	nf = str
	fmt.Printf("空接口保存的数据2=%v,空接口保存2的类型：%T\n",nf,nf)
    judeType(nf)
	s := []interface{}{"hello",88,true}
	fmt.Println(s)

    m := map[string]interface{}{} // 映射的声明
	m["name"] = "tom"
    m["age"] = 35
	m["year"] = 10
	fmt.Println(m)

	var car struct{
		price interface{}
	}
	
	car.price = 1000
	fmt.Printf("car.1的值== %v, car1.type类型 =%T \n",car.price,car.price)
	car.price = "hello,777"
	fmt.Printf("car.2的值== %v, car2.type类型 =%T \n",car.price,car.price)


	num := 101
	var nbf interface{} = num

	value,ok :=  nbf.(int)
    fmt.Println(value) // 101
    fmt.Println(ok) // true
    fmt.Printf("%v\n",nbf) // 101


	/*
	当空接口类型是基础类型、数组类型或结构体类型时，不论空接口的类型是否相同，都可以比较空接口保存的值。
	如果空接口类型是切片类型或映射类型，则无法比较空接口保存的值。
	*/

}

//类型分支
func judeType(v interface{}){
	switch v.(type){
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
	fmt.Println("default")
	}
}