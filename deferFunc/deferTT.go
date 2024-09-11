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
	fmt.Println("我是third")
	return num + 30
}

func DeferTud() {

	num := 10
	fmt.Println("程序开始")
	defer first(num)               //10
	defer fmt.Println(first(num))  //20
	defer fmt.Println(second(num)) //30
	defer fmt.Println("num1", num)
	num = 20
	defer fmt.Println("num2", num)
	defer fmt.Println(third(num)) //50
	fmt.Println("程序结束")

	/*
	程序开始
	我是first
	我是second
	我是third
	程序结束
	50
	num2 20
	num1 10
	30
	20
	我是first
	*/

}

// 匿名函数的延迟执行
func AnonymityFunc() {
	fmt.Println("程序开始")
	defer func() {
		fmt.Println("程序延迟调用")
	}()
	fmt.Println("程序结束")

	/*
			程序开始
		程序结束
		程序延迟调用

	*/

}

// 返回值未设置变量名
func Test() int {
	var i int // 0
	defer func() {
		i++
		fmt.Println(i)
		fmt.Println("defer2:", i) //2
	}()

	defer func() {
		i++
		fmt.Println(i)
		fmt.Println("defer1:", i) //1
	}()

	return i

	/*
		defer1: 1
		defer2: 2
		0
	*/

}
