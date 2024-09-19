package mylearn
import (
	"fmt"
	"os"
	"errors"
	"math"
	
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