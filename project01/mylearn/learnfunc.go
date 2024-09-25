package mylearn

import (
	"fmt"
	"strconv"
)

type job struct {
	position string
	year int
}

type info struct{
	name string
	age int
	record job
}

func newJob(position string, year int) *job{
	return &job{
		position: position, 
		year: year,
	}
}

func newInfo(name string, age int, position string, year int) *info{
	return &info{
		name: name,
		age: age,
		record: *newJob(
			position,
			year,
		),
	}
}

func JobTestSruct(){
	p := newInfo("张三",32,"艺术总监",5)
	fmt.Printf("姓名=%v\n",p.name)
	fmt.Printf("年龄=%v\n",p.age)
	fmt.Printf("职位=%v\n",p.record.position)
	fmt.Printf("工作年限=%v\n",p.record.year)
	fmt.Println(p.get_info("CTO",10))
}


func(p info) get_info(position string, year int) string {
	p.record.position = position
	return p.name + strconv.Itoa(p.age) + p.record.position + strconv.Itoa(year)
}

