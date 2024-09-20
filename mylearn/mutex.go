package mylearn
import(
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex
func printData(){
	mutex.Lock()

	for i := 30; i < 33; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
	mutex.Unlock()
}

// 互斥锁
func MutexTest(){
	go printData()
	mutex.Lock()
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
	mutex.Unlock()
	time.Sleep(time.Second * 10)
	fmt.Println("end")
}