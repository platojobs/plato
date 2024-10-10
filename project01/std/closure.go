package std

import (
	"fmt"
	"path"
)

func makeClosurFunc(suffix string) func(name string)string {

   return func(name string)string{
		if path.Ext(name) == ""{
			return name + suffix
		}else{
			return "2024:" + name
		}
	}

}


func ClosureTest(){

   mp4func := makeClosurFunc(".mp4")
   txtfunc := makeClosurFunc(".txt")
   fmt.Println(mp4func("jielun.mp4"))
   fmt.Println(txtfunc("jielun"))

}

