package mylearn

import (
	"fmt"
	"time"
)

func getoff(){

	for i := 5;i > 0; i--{
        fmt.Printf("还有 %v 位乘客下车\n",i)
		time.Sleep(1*time.Second)
	}
}


func Goroutine(){
    go getoff()

	for i := 1; i < 6; i++ {
		fmt.Printf("第 %v 位乘客上车\n",i)
		time.Sleep(1*time.Second)
	}
}


