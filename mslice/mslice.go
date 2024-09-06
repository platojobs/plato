package mslice

import (
	"fmt"
)

// 删除切片的元素
func DeleteSlice() {

	sli := []int{1, 2, 3 ,4, 5, 6, 7, 8, 9}
	fmt.Println(sli)
	sli = append(sli[:2], sli[3:]...)
	fmt.Println(sli)

	ss := sli[:len(sli)-2]
	fmt.Println(ss)

	/*
		[1 2 3 4 5 6 7 8 9]
		[1 2 4 5 6 7 8 9]
		[1 2 4 5 6 7]
	*/

}

func InitSlicestu(){
	arr := [4]int{1,2,3,4}
    slice := arr[:]
	fmt.Println(slice)
	fmt.Printf("slice = %v\n",slice)

	str := "hello"
	fmt.Println(len(str))

	numbers := []int{}
    for i := 0; i < 10; i++ {
        numbers = append(numbers, i)
		fmt.Println(numbers)
        fmt.Printf("len: %d  cap: %d pointer: %p\n", len(numbers), cap(numbers), numbers)
    }
    fmt.Println(numbers)
}

func CopyOperation(){

	slice_1 := []int{1, 2, 3, 4, 5}  // 目标切片
	slice_2 := []int{5, 4, 3} //源切片
	copy(slice_1, slice_2) //把2逐个塞到1中去
	fmt.Println(slice_1) //[5 4 3 4 5]
	slice_3 := []int{}
	copy(slice_3,slice_2) //把2塞到3中，因为3此时是空切片，所以塞不进去
	fmt.Println(slice_3)//[]
	copy(slice_2,slice_3) //把空切片塞到2中去，目标切片不会变化
	fmt.Println(slice_2) //[5 4 3]
 
}

func GetTotal()func(nme string,price float64) float64 {
	var total float64
	return func (name string,price float64) float64 {
		fmt.Printf("name=%s,price=%f\n",name,price)
		total = total + price
		return total
	}
}

