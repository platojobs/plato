package mylearn

import (
	"fmt"
	"project01/mymath"
	"strings"
	"errors"
)

func Stud() {

	const a int = 10
	fmt.Println(a)

	v := 1
	Incr(&v)
	fmt.Println(Incr(&v))

	fmt.Println(Fu())
	fmt.Println(Fu())

	p := new(int)
	fmt.Println(*p)

	fmt.Println(mymath.Mhello("Jobs"))

	str := "hello,world"
	nstr := str[0:2]
	fmt.Println(nstr)

	fmt.Println(len([]rune(str)))
	fmt.Println([]rune(str))
	nesstring := strings.Map(func(r rune) rune {
		fmt.Println(r)
		if r == 'l' {
			return 'L'
		}
		return r
	}, str)
	fmt.Println(nesstring)

	fm := "六三wer"  // 有中文需要转化
	fmy := fm[0:2] //22= �
	fmt.Println("22=", fmy)
	srn := []rune(fm)
	str2 := srn[0:2]
	fmt.Println("77=", string(str2))

	name := "Plato Jobs"
	strarr := strings.Fields(name) //空格分割字符串
	fmt.Println("strarr =", strarr, strarr[0])
}

func Incr(p *int) int {
	*p++ // 非常重要：只是增加p指向的变量的值，并不改变p指针！！！
	fmt.Println(*p, p)
	return *p
}

func Fu() *int {
	v := 1
	fmt.Println(&v)
	return &v
}

// 切片学习
func Slice_stu() {
	// 定义一个切片
	var numbers = make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(numbers), cap(numbers), numbers)
	numbers = append(numbers, 1, 2, 6, 9, 9999)
	fmt.Println(numbers)
	s2 := []string{}
	fmt.Println(s2)
	sarr := [4]int{1, 2, 3, 4}
	s6 := sarr[0:3]
	fmt.Println(s6)
	fmt.Println(s6[:])
	fmt.Println(s6[1:])
	fmt.Println(s6[:3])

	var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var slice0 []int = arr[2:8]
	var slice1 []int = arr[0:6]        //可以简写为 var slice []int = arr[:end]
	var slice2 []int = arr[5:10]       //可以简写为 var slice[]int = arr[start:]
	var slice3 []int = arr[0:len(arr)] //var slice []int = arr[:]
	var slice4 = arr[:len(arr)-1]      //去掉切片的最后一个元素
	slice66 := arr[0:]                 //等价slice3

	fmt.Printf("su：arr %v\n", arr)
	fmt.Printf("su：slice0 %v\n", slice0)
	fmt.Printf("su：slice1 %v\n", slice1)
	fmt.Printf("su：slice2 %v\n", slice2)
	fmt.Printf("su：slice3 %v\n", slice3)
	fmt.Printf("su：slice4 %v\n", slice4)
	fmt.Printf("su：slice66 %v\n", slice66)
	fmt.Printf("-----------------------------------\n")
	arr2 := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	slice5 := arr[2:8]
	slice6 := arr[0:6]         //可以简写为 slice := arr[:end]
	slice7 := arr[5:10]        //可以简写为 slice := arr[start:]
	slice8 := arr[0:len(arr)]  //slice := arr[:]
	slice9 := arr[:len(arr)-1] //去掉切片的最后一个元素
	fmt.Printf("局部变量： arr2 %v\n", arr2)
	fmt.Printf("局部变量： slice5 %v\n", slice5)
	fmt.Printf("局部变量： slice6 %v\n", slice6)
	fmt.Printf("局部变量： slice7 %v\n", slice7)
	fmt.Printf("局部变量： slice8 %v\n", slice8)
	fmt.Printf("局部变量： slice9 %v\n", slice9)

	s10086 := []int{1, 2, 3, 4, 5, 6: 7, 8, 9, 10}
	fmt.Println(s10086)
	pzslice := &s10086[3]
	*pzslice += 110
	fmt.Println(s10086)

	data := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(data)

}

// make函数搞切片
func Make_slice() {
	// make
	slice0 := make([]string, 3, 5)
	slice0[0] = "a"
	// fmt.Println(slice0)

	d := [5]struct {
		x int
	}{}

	s := d[:]

	d[1].x = 10
	s[2].x = 20

	fmt.Println(d)
	fmt.Printf("%p, %p\n", &d, &d[0])
	//[{0} {10} {20} {0} {0}]
	//0xc0000b2000, 0xc0000b2000

}

func Append_slice() {
	a := []int{1, 2, 3, 4: 990}
	b := []int{4, 5, 6}
	c := append(a, b...)
	fmt.Println(c) //[1 2 3 4 5 6]
	d := append(c, 4)
	fmt.Println(d) //[1 2 3 4 5 6 4]
	e := append(d, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(e) //[1 2 3 4 5 6 4 4 5 6 7 8 9 10]
	data := [...]int{0, 1, 2, 3, 4, 10: 0}
	fmt.Printf("cap = %d\n", cap(data))
	sdata := append(data[:], 100)
	fmt.Println(sdata)

	ss := sdata[:2:3] // 3-0 =cap ,len = 2-1 其实索引为0
	fmt.Println(ss)   //[0 1]

}

// slice中cap重新分配规律
func Cap_reviewMemerty() {
	s := make([]int, 0, 10)
	c := cap(s)
	fmt.Println(c) // 10 cap
	for i := 0; i < 20; i++ {
		s = append(s, i)
		fmt.Println(s)
		if n := cap(s); n > c {
			fmt.Printf("cap: %d -> %d\n", c, n)
			c = n
			fmt.Printf("len: %d\n c = %d", len(s), c)
		}
	}
}

//格式化工具  Sprintf

func Fmt_str() {
	greetstring := "hello,world"
	ss := fmt.Sprintf("%s,Jobs", greetstring)
	fmt.Println(ss)

	var arr_1 [5]int
	fmt.Println(arr_1)
	arr_1[0] = 101010
	fmt.Println(arr_1)

	arr_5 := [5]int{0: 3, 1: 5, 4: 6}
	fmt.Println(arr_5)

}

func ModifyArr(a [5]int) {
	a[0] = 200
	fmt.Println(a)
}



func Functions(){

	fn := func ()  {
		fmt.Println("hello-func")
	}
	fn()

	fns := []func(){
		func() {
			fmt.Println("hello-func1")
		},
		func() {
			fmt.Println("hello-func2")
		},
		func() {
			fmt.Println("hello-func3")
		},
	}
	fns[0]()
   
	d := struct {
		fn func() string
	}{
		fn: func() string { 
			return "hello" 
		},
	}
    fmt.Println(d.fn())

	fc := make(chan func() string,2)
	fc <- func() string { return "hello" }
	fmt.Println(<- fc)
	fmt.Println(fc)
	
	

}


func anonymous() func() int {
	i := 0
	b := func() int {
		i++
		fmt.Println(i)
		return i
	}
	return b
}

func AnoyMain() {
	
   c := anonymous()
   c()
   c()
   c()

   anonymous()

}


// 返回2个函数类型的返回值
func test01(base int) (func(int) int, func(int) int) {
    // 定义2个函数，并返回
    // 相加
    add := func(i int) int {
        base += i
        return base
    }
    // 相减
    sub := func(i int) int {
        base -= i
        return base
    }
    // 返回
    return add, sub
}


func TwoFunc(){

	f1, f2 := test01(10)
    // base一直是没有消
    fmt.Println(f1(1), f2(2))
    // 此时base是9
    fmt.Println(f1(3), f2(4))

	/*
	11 9
    12 8
	*/
}

func Whatever(){
	var whatever [5]struct{}
	 for i := range whatever {
		defer func() { 
			fmt.Printf("helllo==%d\n",i)
			fmt.Println(i) 
			}()
	}
	// defer语句的变量在defer声明时就已经确定了
    // 先进后出
	/*
	4
	3
	2
	1
	0
	*/


}


type Test struct {
	name string
}

func(t *Test)close(){
  fmt.Println(t.name,"closed")
}

func CClose(){
	ts := []Test{{"a"},{"b"},{"c"}}
	for _,t := range ts {
		defer t.close()
	}
}


func ttest() {
    x, y := 10, 20

    defer func(i int) {
        println("defer:", i, y) // y 闭包引用
    }(x) // x 被复制

    x += 10
    y += 100
    println("x =", x, "y =", y)

}

func CCmain() {
    ttest()
	/*
	x = 20 y = 120
    defer: 10 120
	*/
}

/*
third defer err divided by zero!
second defer err <nil>
first defer err <nil>
*/

func Foo(a, b int) (i int, err error) {
	defer fmt.Printf("first defer err %v\n", err)
	defer func(err error) { fmt.Printf("second defer err %v\n", err) }(err)
	defer func() { fmt.Printf("third defer err %v\n", err) }()
	defer func(){ fmt.Println(err) }()
	if b == 0 {
		err = errors.New("divided by zero")
		return
	}
  
	i = a / b
	return
	
  }

