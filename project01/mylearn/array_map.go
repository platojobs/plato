package mylearn

import (
	"fmt"
)

func Arraydtest(){
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)
    c := names
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
	fmt.Println(c)

	q := [ ]int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	sss := make(map[string]any)
	sss["name"] ="john"
	sss["age"] = 20
	fmt.Println(sss)
	
	m := make(map[string]int)

	m["答案"] = 42
	fmt.Println("值：", m["答案"])

	m["答案"] = 48
	fmt.Println("值：", m["答案"])

	delete(m, "答案")
	fmt.Println("值：", m["答案"])

	v, ok := m["答案"]
	fmt.Println("值：", v, "是否存在？", ok)
}
