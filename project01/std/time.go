package std

import(
	"fmt"
	"time"
)

func Time_Test(){
  stime()
}


func stime(){
    today := time.Now()
	fmt.Println(today)
	timestamp1 := today.Unix() //10位时间戳
	timestamp3 := today.UnixMilli() //毫秒
	timestamp4 := today.UnixMicro()//微秒
	timestamp5 := today.UnixNano() //纳秒
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp3:%v\n", timestamp3)
	fmt.Printf("current timestamp4:%v\n", timestamp4)
	fmt.Printf("current timestamp5:%v\n", timestamp5)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(timeFromTimestamp(1727248475))
}

func timeFromTimestamp(timestamp int64) (year, month, day, hour, minute, second int) {
	timeobj := time.Unix(timestamp, 0)
	year = timeobj.Year()     //年
	month = int(timeobj.Month()) //月
	day = timeobj.Day()      //日
	hour = timeobj.Hour()     //小时
	minute = timeobj.Minute() //分钟
	second = timeobj.Second() //秒
	return
}


//变量的批量声明
func VarTT(){

	var (
		name string
		age int
		married bool
	)
	// name = "jobs"
	// age = 10
	// married = false

	fmt.Printf("name:%s,age:%d,married:%t\n",name,age,married)
}
