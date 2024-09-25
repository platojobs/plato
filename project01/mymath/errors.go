package mymath

import (
	"errors"
	"fmt"
)


func f (arg int)(int,error){
	if arg == 35{
		return -1,errors.New("35岁,是个死局")
	}
	return arg + 35 ,nil
}

var ErrOutTea = fmt.Errorf("no more tea")
var ErrPower = fmt.Errorf("can't power")

func makeTea(arg int)error {
	if arg == 2{
		return ErrOutTea
	}else if arg == 4 {
		return fmt.Errorf("making tea: %w",ErrPower)
	}
	return nil
}

func EroorsTest(){
	for _,i := range []int{7,35,42,45}{
		if r,e := f(i); e != nil{
			fmt.Println("f failed:",e)
		}else{
			fmt.Println("f worked:",r)
		}
	}


	for i := range 5 {
		if err := makeTea(i); err != nil{
			if errors.Is(err,ErrOutTea){
				fmt.Println("we should buy new tea")
			}else if errors.Is(err,ErrPower){
				fmt.Println("Now it is dark")
			}else{
				fmt.Println("unknown error:",err)
			}
			continue
		}
		fmt.Println("Tea is ready")
	}


}