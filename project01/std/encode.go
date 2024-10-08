package std

import (
	"encoding/json"
	"fmt"
	"time"
)

type Result struct {
	Code int   `json:"code"`
	Msg string `json:"msg"`
	Data User //`json:"data"`
}

type User struct {
	Name string `json:"name"`
	Age int     `json:"age"`
	Salary int  `json:"salary"`
}

func EncodeTT(){
	var res Result
	res.Code = 100
	res.Msg = "success"
	res.Data = User{
		Name:   "唐春",
		Age:    22,
		Salary: 100000,
	}
    
    jsons,err := json.Marshal(res)
	if err != nil {
		fmt.Println("json err:", err)
	}
    fmt.Println("jsons:",string(jsons))

	var res2 Result
	err = json.Unmarshal(jsons, &res2)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println("res2:",res2)

}

func PrrrTT(){
 person := [3]string{"tom","jack","marry"}
 fmt.Println(person)
 
 for k,v := range person{
	fmt.Printf("key:%d,value:%s\n",k,v)
 }
 
 //使用空白符 -取值
 for _,v := range person{
	fmt.Printf("value:%s\n",v)
 }
 //取索引
 for k := range person{
	fmt.Printf("key:%d\n",k)
 }
 //索引取值
 for i := 0; i < len(person);i++{
    fmt.Printf("person[%d]=%s\n",i,person[i])
 }

}


func PPPerson(){
	person := []string{"tom","jack","marry"}
	fmt.Println("person:",person)
	fmt.Println("person len:",len(person))

}


func producer(ch chan string) {
	fmt.Println("producer start")
	ch <- "a"
	ch <- "b"
	ch <- "c"
	ch <- "d"
	fmt.Println("producer end")
}

func customer(ch chan string) {
	for {
		msg := <- ch
		fmt.Println(msg)
	}
}

func PPmain() {
	fmt.Println("main start")
	ch := make(chan string, 3)
	go producer(ch)
	//go customer(ch)

	time.Sleep(1 * time.Second)
	fmt.Println("main end")
}


func TTTRR(){
	// 创建一个3个元素缓冲大小的整型通道
	ch := make(chan int, 3)
	// 查看当前通道的大小
	fmt.Println(len(ch))
	// 发送3个整型元素到通道
	for i := 0; i < 3; i++ {
		 ch <- i
	}
	// 查看当前通道的大小
	fmt.Println(len(ch))
	for i := 0; i < 3; i++ {
		 fmt.Println(<-ch)
	}
	// 查看当前通道的大小
	fmt.Println(len(ch))
	// 查看当前通道的容量
	fmt.Println(cap(ch))


	

}