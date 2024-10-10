package main

import (
	"fmt"
	"maps"
	"net/http"
	"unicode/utf8"

	//"mutiresult"
	//"deferFunc"
	"project01/gee"
	"project01/mylearn"
	"project01/mymath"
	"project01/std"
	//"project01/pgin"
	//"project01/redis"
	//"project01/std"
	//"project01/dup01"
	//"project01/pgin"
	//"project01/gee"
)

func main() {
	
	//std.FMT()
	//std.Sprint()
    //mylearn.Vertt()
	//mylearn.Personmain()
	//pgin.HttpTT()
    //pgin.MySql()
	// pgin.CreateTable()
	//std.EncodeTT()
	//std.PrrrTT()
	//std.PPmain()
     std.ClosureTest()
	
}

func gintest(){
	engine := gee.New()
	engine.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	engine.GET("/hello", func(c *gee.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	engine.Run(gee.ListenAddrPort)
}



func othertest(){

	//	mapFunc()
	//sums("a","b","c")
//	fmt.Println(sums("a","b","c"))

	// next := insss()
	// fmt.Println(next())
	// fmt.Println(next())
	// fmt.Println(next())
	// newnext := insss()
	// fmt.Println(newnext())
	// fmt.Println(factorial(7))
    
	//mymath.TTest()
   // mymath.TestEnum()
	//mymath.StructIn()
	//mymath.GnericsTest()
	//mymath.EroorsTest()
    //mymath.CustomError()
	//mymath.GoroutinesTest()
	//mymath.ChanTest()
	mymath.EroorsTest()
	
	//redis.Redis()
}




func stes(){
	i := 10
	fmt.Println(i)
	zeroral(i)
	fmt.Println(i)
	zeroptr(&i)
	fmt.Println(i)
	fmt.Println("pointer:=",&i)
    stringDD()
}


func testUU(){
	const boilingF = 212.0 //常量
	fmt.Printf("------%f-------\n", boilingF)

	//mutiresult.Error_debug()
	//mutiresult.Defertest()
	// deferFunc.DeferTud()
	//fmt.Println(deferFunc.Test())
	// fmt.Println(deferFunc.Test1())
	fmt.Println("=======1========")
	// mylearn.Functions()
	// mylearn.AnoyMain()
	fmt.Println("-------2--------")
	// mylearn.TwoFunc()
	fmt.Println("--------3----------")
	// mylearn.Whatever()
	fmt.Println("----------4----------")
	// mylearn.CClose()
	fmt.Println("--------5----------")
	//mylearn.CCmain()
	fmt.Println("--------6----------")
	mylearn.Foo(2, 0)

	mylearn.JobTestSruct()
	fmt.Println("------------7------------")
	mylearn.CarTest()
	fmt.Println("------------8------------")
	//mylearn.Minterface()
	//mylearn.OpenFil()
	fmt.Println("------------9------------")
	//mylearn.QsqrtTest()
	fmt.Println("------------10------------")
	//mylearn.SqrtErrorTest()
	fmt.Println("------------11------------")
	//mylearn.MerrorF()
	fmt.Println("------------12------------")
	//mylearn.AreaError()

	fmt.Println("------------13------------")
	//mylearn.UnwrapError()
	fmt.Println("------------14------------")
	//mylearn.IsError()
	fmt.Println("------------15------------")
	//mylearn.PanicTest()
	fmt.Println("------------16------------")
	//mylearn.RecoverTest()
	fmt.Println("------------17------------")
	//mylearn.Goroutine()
	fmt.Println("------------18------------")
	//mylearn.ChannelTest()
	fmt.Println("------------19------------")
	//mylearn.ChanelSRtest()
	fmt.Println("------------20------------")
	//mylearn.NoTouchChanel()
	fmt.Println("------------21------------")
	//mylearn.MNchanelnum()
	fmt.Println("------------22------------")
	//mylearn.MchnelSelect()
	fmt.Println("------------23------------")
	//mylearn.AtmoicAdd()
	fmt.Println("------------24------------")
	//mylearn.MutexTest()
	fmt.Println("------------25------------")
	//mylearn.OperationFile()
	//mylearn.AppendFile()
	//mylearn.ReadFile()
	//mylearn.GobEncode()
	//mylearn.GobDecode()
	//mylearn.BinaryEncode()
	//mylearn.BinaryDecode()
	//mylearn.JsonEcodeTT()
	//mylearn.JsonDecode()
	//mylearn.XmlEcode()
	//mylearn.XmlDecode()
	const b = 1000
	fmt.Println(int64(b))
    //testArray()
	//slice()
}


func testArray(){
	a := [5]int{}
	fmt.Println(a)
	a[0] = 10 //改变数组a的首个元素
	fmt.Println(a)
	fmt.Println(len(a)) //数组的长度



	d := [2][3]int{}
	fmt.Println(d)
	d[0][0]= 2
	fmt.Println(d)
    d[0][1] = 3
	fmt.Println(d)
    d[1][1] = 4
	fmt.Println(d)

}


func slice(){

	a := make([]string,3)
	fmt.Println(a)
    a[0] = "a"
    a[1] = "b"
	a[2] = "c"
	fmt.Println(a)

	 a = append(a, "c")
	 fmt.Println(a)
	 fmt.Println(len(a)) // 长度为4

	s := append(a,"d") // 追加一个元素
	fmt.Println(s)
	cd := append(s,a...)  // s和a组合为一个新的cd切片
	fmt.Println(cd)
    
    l := cd[2:5] //c c d
	fmt.Println(l)
    
	l = cd[2:]   //c c d a b c c
	fmt.Println(l)
     

}

func mapFunc(){

	m := make(map[string]string)
	m["a"] = "jobs"
	m["b"] = "apple"
    m["c"] = "google"
	m["d"] = "facebook"
	fmt.Println(m)
    //删除
	delete(m,"a")
	fmt.Println(m)
	// 判断是否存在
    n := map[string]string{"a":"jobs"}
	if maps.Equal(m,n){
		fmt.Println("m和n相等")
	}else{
		fmt.Println("m和n不相等")
	}

	value,errors := m["c"]
	if errors {
		fmt.Println(value)
	}else{
		fmt.Println("不存在")
	}


	//clear(m)
	//fmt.Println(m)

    sint := add(2,3)
	fmt.Println(sint)

	for key,i := range m{
	    fmt.Printf("key=%s,value=%s\n",key,i)
	}

}




//参数类型相同，可以只写最后一个
func add(a,d int)int{
	return a + d
}

//函数的可变参数
func sums(nums ...string) string{
     s := ""
	 for _,i := range nums{
		 s += i
	 }
	 return s
}

// 闭包
func insss() func () int {
	i := 0
	return  func() int {
		i++
		//fmt.Println(i)
		return i
	}
	//闭包
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func zeroral(intarl int){
	intarl = 0
}
func zeroptr(intarl *int){
	*intarl = 0
}

func stringDD(){
	 const s = "สวัสดี"
	 fmt.Println("len=",len(s))
	 for i := 0; i < len(s); i++ {
		fmt.Printf("%v=%x ",i, s[i])
	 }
	 fmt.Println("Runne count :", utf8.RuneCountInString(s))

	 for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	 }
	 fmt.Println("\nUsing DecodeRuneInString.")
	 for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
		examineRune(runeValue)
	 }
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}