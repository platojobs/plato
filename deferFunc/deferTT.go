package deferFunc

import "fmt"

func first(num int) int {
	fmt.Println("我是first")
	return num + 10
}

func second(num int) int {
	fmt.Println("我是second")
	return num + 20
}

func third(num int) int {
	fmt.Println("程序开始")
	return num + 30
}

func DeferTud() {

	num := 10
	fmt.Println("程序开始")
	defer first(num)               //10
	defer fmt.Println(first(num))  //20
	defer fmt.Println(second(num)) //30
	num = 20
	defer fmt.Println(third(num)) //50
	fmt.Println("程序结束")

	/*
	程序开始
	我是first
	我是second
	程序开始
	程序结束
	50
	30
	20
	我是first
	*/

}
