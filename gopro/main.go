package main

import (
	"fmt"
	//"mutiresult" 
	"deferFunc"
)

const boilingF = 212.0 //常量

func main() {
 fmt.Printf("------%f-------\n", boilingF)
 
 //mutiresult.Error_debug()
  //mutiresult.Defertest()
  deferFunc.DeferTud()
  fmt.Println(deferFunc.Test())

}

