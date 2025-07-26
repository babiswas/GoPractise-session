package app

import (
	"container/list"
	"fmt"
)

func NewStack(num []int) {
	stack := list.New()
	for i := 0; i < len(num); i++ {
		stack.PushBack(num[i])
	}
	for e := stack.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value.(int))
	}
}
