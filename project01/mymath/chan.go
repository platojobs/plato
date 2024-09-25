package mymath
import (
	"fmt"
	
)


func ChanTest(){
	messages  := make(chan string)
	go func(){
		messages <- "ping"
	}()

	masg := <-messages
	fmt.Println(masg)
}



func ping(pings chan<- string, msg string) {
    pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

func Pingsmain() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "passed message") 
    pong(pings, pongs)
    fmt.Println(<-pongs)
}