package main

import (
	"fmt"
	//"mutiresult"
	//"deferFunc"
	"mylearn"
)

const boilingF = 212.0 //常量

func main() {
	fmt.Printf("------%f-------\n", boilingF)

	//mutiresult.Error_debug()
	//mutiresult.Defertest()
	// deferFunc.DeferTud()
	//fmt.Println(deferFunc.Test())
	// fmt.Println(deferFunc.Test1())
	fmt.Println("=======1========")
	// mylearn.Functions()
	// mylearn.AnoyMain()
	fmt.Println("-------2--------")
	// mylearn.TwoFunc()
	fmt.Println("--------3----------")
	// mylearn.Whatever()
	fmt.Println("----------4----------")
	// mylearn.CClose()
	fmt.Println("--------5----------")
	//mylearn.CCmain()
	fmt.Println("--------6----------")
	mylearn.Foo(2, 0)

	mylearn.JobTestSruct()
	fmt.Println("------------7------------")
	mylearn.CarTest()
	fmt.Println("------------8------------")
	mylearn.Minterface()
	mylearn.OpenFil()
	fmt.Println("------------9------------")
    mylearn.QsqrtTest()
}
