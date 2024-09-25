package mymath

import (
	"errors"
	"fmt"
)

type argError struct{
	arg int
	message string
}

func (e *argError) Error()string{
	return fmt.Sprintf("%d -- %s",e.arg,e.message)
}

func ff(arg int)(int,error){
    if arg == 42 {
		return -1,&argError{arg,"can't work with it"}
	}
	return arg +3,nil
}

func CustomError(){
	_,ererrors := ff(42)
	var ae *argError
	if errors.As(ererrors,&ae){
       fmt.Println(ae.arg)
	   fmt.Println(ae.message)
	}else{
		fmt.Println("err doesn't match argError")
	}
}