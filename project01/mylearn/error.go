package mylearn

import (
	"errors"
	"fmt"
	"math"
	"os"

)



func OpenFil(){
	fmt.Println("------error---------")
   _,err := os.Open("/.rt.txt")
   if err == nil{
   	fmt.Println("open file success")
   }else{
	fmt.Println(err)
   }
}

func sqrt(num float64)(float64,error){
	if num < 0{
		//自定义错误信息最简单的方法就是调用errors包中的New()函数
		return -1,errors.New("负数没有平方根")
	}
	return math.Sqrt(num),nil
}

func QsqrtTest(){
	result,errr0r := sqrt(-1)
	if errr0r != nil{
		fmt.Println(errr0r)
	}else{
		fmt.Println(result)
	}
}


type sqrtError struct{
	num float64
}
func (se sqrtError) Error() string{
	return fmt.Sprintf("负数没有平方根：%v \n",se.num)
}

func dsqrt(snum float64)(float64,error){

    if snum <0 {
		return -1,sqrtError{num:snum}
	}
	return math.Sqrt(snum),nil

}

func SqrtErrorTest(){
	reslut , serror := dsqrt(-10)
	if serror != nil {
		fmt.Println(serror)
	}else{
		fmt.Println(reslut)
	}
}

func MerrorF(){
	const name,id = "rose",101
	err := fmt.Errorf("name = %v, id = %v 不存在",name,id)
	fmt.Println(err)
	
}


func area(length float64,width float64)(float64,error){
	if length < 0 || width < 0{
		return -1,errors.New("长度或宽度不能为负数")
	}
	return length * width,nil
}

func AreaError(){
	length := -10.0
	width := 10.0
	area,err := area(length,width)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(area)
	}
}

//Unwrap 函数
func UnwrapError(){
	err := errors.New("this is an error1")
	err2 := fmt.Errorf("this is an error2:%w",err)
	fmt.Println(errors.Unwrap(err2)) //解包
}

//通过errors包中的Is()函数可以判断两个error是否是同一个error。Is()函数
/* 如果err和target是同一个error则返回true。如果err是一个嵌套error，
而target也包含在这个嵌套error链中，那么也返回true。其他情况返回false。*/

func IsError(){
	err := errors.New("this is an error1")
	err2 := fmt.Errorf("this is an error2:%w",err)
	fmt.Println(errors.Is(err2,err)) // err是target
    // 如果err2 嵌套了err，那么返回true
}

// 宕机测试
func PanicTest(){
	defer func(){
		fmt.Println("free me")
	}()
	panic("this is a panic") // 宕机测试
}


//宕机恢复
//通过recover()函数可以获取panic信息，后面的程序可以正常按顺序执行。
func RecoverTest(){
	defer func(){
		if err := recover();err != nil{
			fmt.Println(err)
			fmt.Println("recover success")
		}
		fmt.Println("恢复执行")
	}()
	fmt.Println("hello")
	panic("this is a panic") // 宕机测试
	
}