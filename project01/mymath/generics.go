package mymath

import (
	"fmt"
	"iter"
	"slices"
)

func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
    for i := range s {
        if v == s[i] {
            return i
        }
    }
    return -1
}

type List[T any] struct {
    head, tail *element[T]
}
type element[T any] struct {
    next *element[T]
    val  T
}


func (lst *List[T]) Push(v T) {
    if lst.tail == nil {
        lst.head = &element[T]{val: v}
        lst.tail = lst.head
    } else {
        lst.tail.next = &element[T]{val: v}
        lst.tail = lst.tail.next
    }
}

func (lst *List[T]) AllElements() []T {
    var elems []T
    for e := lst.head; e != nil; e = e.next {
        elems = append(elems, e.val)
    }
    return elems
}

//迭代器
func (lst *List[T]) All() iter.Seq[T]{
	return func(yield func(T) bool) {
		for w := lst.head; w != nil; w = w.next{
			if !yield(w.val){
				return
			}
		}
	}
}

func getFib() iter.Seq[int]{
	return func(yield func(int) bool) {
		a,b := 1,1
		for{
			if !yield(a){
				return
			}
			a,b = b , a+b
		}
	}
}


func GnericsTest() {
    var s = []string{"foo", "bar", "zoo"}

    fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))

    _ = SlicesIndex[[]string, string](s, "zoo")
    lst := List[int]{}
    lst.Push(10)
    lst.Push(13)
    lst.Push(23)
	rt := lst.AllElements()
    fmt.Println("list1:", lst.AllElements())
    fmt.Println("list2:", lst.All())
	fmt.Println("list3:",&rt)
	for e := range lst.All(){
        fmt.Println("all:",e)
	}
    all := slices.Collect(lst.All())
	fmt.Println("all3:",all)

	for n := range getFib(){
		if n >= 10{
			break
		}
		fmt.Println(n)
	}

}