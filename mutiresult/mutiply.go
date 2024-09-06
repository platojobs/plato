package mutiresult

import (
	"fmt"
	"math"
	"strconv"
)

func Error_debug() {

	orig := "ABC"
	var newS string
	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	an, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
		return
	}
	fmt.Printf("The integer is %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)

}

// 为函数准备充足的返回值
func Error_re() {
	t ,k := mySqrt(25.0)
	fmt.Println(t)
	fmt.Println(k)
}

func mySqrt(f float64) (v float64, ok bool) {
	if f < 0 {
		return
	}
	return math.Sqrt(f), true
}
