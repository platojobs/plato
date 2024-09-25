package mylearn

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int64
	wg sync.WaitGroup
)

func calculate( number0 int ){
	for i := 0; i < 2; i++ {
		atomic.AddInt64(&counter,1)
		runtime.Gosched()
	}
    wg.Done() // 对wg执行减1操作，直到位0，解除主程序的阻塞状态
}

//计数
func AtmoicAdd(){
	wg.Add(2)  // 设置并发程序的等待数量
	go calculate(1) // 启动第一个并发程序
	go calculate(2) // 启动第二个并发程序
	wg.Wait() // 主程序进如阻塞状态
	fmt.Println(counter) // 4
	
}