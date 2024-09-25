package mymath

import (
	"fmt"
	"time"
)

func gf(from string){
	for idxi := 0; idxi < 3; idxi++ {
		fmt.Println(from ,":",idxi)
	}
}

func GoroutinesTest(){
	gf("direct")
	go gf("goroutine") //轻量级的执行线程
	go func(msg string){
		fmt.Println(msg)
	}("going")
	time.Sleep(time.Second)
	fmt.Println("done")

	/*
direct : 0
direct : 1
direct : 2
going
goroutine : 0
goroutine : 1
goroutine : 2
	*/




}

