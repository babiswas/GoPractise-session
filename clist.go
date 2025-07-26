package app

import (
	"container/list"
	"fmt"
)

func AddNumbers(nums []int) int {
	sum := 0
	l := list.New()
	for _, val := range nums {
		l.PushBack(val)
	}
	for e := l.Front(); e != nil; e = e.Next() {
		sum += e.Value.(int)
	}
	return sum
}

func AddFrontBack(nums []int) {
	l := list.New()
	for _, val := range nums {
		if (val % 2) == 0 {
			l.PushFront(val)
		} else {
			l.PushBack(val)
		}
	}

	fmt.Println("Displaying from the front:")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}

	fmt.Println("Displaying from the back:")
	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value.(int))
	}
}

func SubList(nums []int) {
	fmt.Println("Example to display how sublists works:")
	l1 := list.New()
	l2 := list.New()

	fmt.Println("Collect all the odd numbers:")
	for _, val := range nums {
		if val%2 != 0 {
			l1.PushBack(val)
		}
	}

	fmt.Println("Collect all the even numbers:")
	for _, val := range nums {
		if val%2 == 0 {
			l2.PushBack(val)
		}
	}

	fmt.Println("Display odd numbers:")
	for e := l1.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}

	fmt.Println("Display even numbers:")
	for e := l2.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}

	fmt.Println("Creating bigger list:")
	l := list.New()
	l.PushBackList(l1)
	l.PushBackList(l2)

	fmt.Println("Traversing the list from front:")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}

	fmt.Println("Traversing the list from back:")
	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value.(int))
	}
}

func ShrinkList(nums []int) {
	fmt.Println("Displaying how to shrink the list:")
	l := list.New()
	for _, val := range nums {
		l.PushBack(val)
	}
	fmt.Println("Shrinking the list")
	for l.Len() != 0 {
		fmt.Println("Length of the list before removal:", l.Len())
		v := l.Back()
		fmt.Println(v.Value.(int))
		l.Remove(v)
		fmt.Println("Length of the list after removal:", l.Len())
	}
}
