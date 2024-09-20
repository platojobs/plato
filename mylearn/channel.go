package mylearn

import (
	"fmt"
	"time"
)

func ChannelTest(){
	chel := make(chan int)
	go func ()  {
		for i := 0; i < 10; i++{
			chel <- i
			time.Sleep(time.Second)
		}
		close(chel)
	}()

	for data := range chel{
		fmt.Println(data)
		if data == 5{
			break
		}
	}
}

//只能发送数据的通道
func electrify(chel chan <- string){
	chel <- "接通电源！"
	close(chel)
}

func start(ch1 <- chan string, ch2 chan <- string){
     data := <- ch1
	 ch2 <- data + "启动"
	 close(ch2)
}

func drive(ch <- chan string){
	for data := range ch{
		fmt.Printf("%s,准备好了,开始出发",data)
	}
}

func ChanelSRtest(){
	ch1 := make(chan string)
	ch2 := make(chan string)
	go electrify(ch1)
	go start(ch1,ch2)
    drive(ch2)
}


//无缓冲的通道
func read(bookname string,ch chan bool){
	fmt.Printf("开始阅读%s\n",bookname)
	ch <- true
}
func NoTouchChanel(){
	cha := make(chan bool)
	go read("《西游记》",cha)
	<- cha
}

//有缓冲的通道
func read2(ch chan int){
	fmt.Println("开始",len(ch))
	ch <- 0
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("end",len(ch))
}
//有缓冲的通道,内部存储数据
func MNchanelnum(){
	chanl := make(chan int,6)
	read2(chanl)
}