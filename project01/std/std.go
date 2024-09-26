package std

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func FMT() {
	fmt.Print("在终端打印信息\n")
	name := "jobs"
	fmt.Printf("我是:%s\n", name)
	fmt.Println("在终端单独显示一行")
}

func Sprint() {
	s1 := "jobs"
	name := "jobs"
	age := 10
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	s3 := "Plato"
	fmt.Println(s1, s2, s3)
	s5 := person{
		name: "jobs",
		age:  10,
	}

	s4 := fmt.Sprintf("%+v", s5) // 输出带有字段名
	s6 := fmt.Sprintf("%v", s5)  //输出不带字段名
	fmt.Println(s4)
	fmt.Println(s6)
	fmt.Printf("%#v\n", s5) //输出结构体字段类型
	fmt.Printf("%T\n", s5)  //输出结构体类型

	o := struct{ name string }{"枯藤"}
	fmt.Printf("%T\n", o)
}

func Errorf() {
	err := fmt.Errorf("这是一个错误")
	fmt.Println(err)
}

func Placeholder() {
	n := 65
	fmt.Printf("%b\n", n)
	fmt.Printf("%c\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)
	fmt.Printf("%X\n", n)

	f := 12.34
	fmt.Printf("%b\n", f)
	fmt.Printf("%e\n", f)
	fmt.Printf("%E\n", f)
	fmt.Printf("%f\n", f)
	fmt.Printf("%g\n", f)
	fmt.Printf("%G\n", f)

	s := "留学"
	fmt.Printf("%s\n", s)
	fmt.Printf("%q\n", s)
	fmt.Printf("%x\n", s)
	fmt.Printf("%X\n", s)

	a := 18
	fmt.Printf("%p\n", &a)
	fmt.Printf("%#p\n", &a)

	sn := 88.88
	fmt.Printf("%f\n", sn)
	fmt.Printf("%9f\n", sn)
	fmt.Printf("%.2f\n", sn)
	fmt.Printf("%9.2f\n", sn)
	fmt.Printf("%9.f\n", sn)

}

func Other() {

	s := "hello"
	fmt.Printf("%s\n", s)
	fmt.Printf("%5s\n", s)
	fmt.Printf("%-5s\n", s)
	fmt.Printf("%5.7s\n", s)
	fmt.Printf("%-5.7s\n", s)
	fmt.Printf("%5.2s\n", s)
	fmt.Printf("%05s\n", s)
}

func Scan(){
	var (
		name string
		age  int
	)
	fmt.Scan("name:",&name,"age:",&age)
	fmt.Println("结果:",name, age)

}