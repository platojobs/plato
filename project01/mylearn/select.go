package mylearn

import (
	"fmt"
	//"time"
)

func MchnelSelect(){
     chi1 := make(chan int)
	 chi2 := make(chan int)
	 go func ()  {
		v := 1
		chi1 <- v
		v2 := <-chi2
		fmt.Println("v-%w,v2-%w",v,v2)
	 }()

	 vv := 2
	 var vv2 int
	 select{
		case chi2 <- vv:
			fmt.Print("-1-\n")
		case vv2 = <- chi1:
		    fmt.Print("-2-\n")
			
	 }
	 fmt.Println(vv,vv2)

}