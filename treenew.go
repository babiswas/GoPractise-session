package app

import (
	"container/list"
	"fmt"
)

type TreeNew struct {
	val   int
	left  *TreeNew
	right *TreeNew
}

func CreateNewTree(val int) *TreeNew {
	return &TreeNew{
		val:   val,
		left:  nil,
		right: nil,
	}
}

func BuildNewTree(nums []int, rootVal int) *TreeNew {
	node := CreateNewTree(rootVal)
	for _, num := range nums {
		node.InsertValue(num)
	}
	return node
}

func (t *TreeNew) InsertValue(val int) {
	if t.val > val {
		if t.left == nil {
			t.left = CreateNewTree(val)
		} else {
			t.left.InsertValue(val)
		}
	} else if t.val < val {
		if t.right == nil {
			t.right = CreateNewTree(val)
		} else {
			t.right.InsertValue(val)
		}
	}
}

func (t *TreeNew) LevelorderT() {
	queue := list.New()
	queue.PushBack(t)
	for queue.Len() != 0 {
		num := queue.Len()
		for num != 0 {
			elm := queue.Front()
			node := elm.Value.(*TreeNew)
			fmt.Printf("%d ", node.val)
			queue.Remove(elm)
			if node.left != nil {
				queue.PushBack(node.left)
			}
			if node.right != nil {
				queue.PushBack(node.right)
			}
			num -= 1
		}
		fmt.Println("")
	}
}

func (t *TreeNew) PostorderTraversal() {
	stack := list.New()
	ptr := t
	for {
		if ptr != nil {
			stack.PushBack(ptr)
			ptr = ptr.left
		} else {
			if stack.Len() == 0 {
				break
			}
			if stack.Back().Value.(*TreeNew).right == nil {
				elm := stack.Back()
				ptr = elm.Value.(*TreeNew)
				fmt.Println(ptr.val)
				stack.Remove(elm)
				for stack.Back().Value.(*TreeNew).right == ptr {
					elm := stack.Back()
					ptr = elm.Value.(*TreeNew)
					fmt.Println(ptr.val)
					stack.Remove(elm)
					if stack.Len() == 0 {
						break
					}
				}
			}
			if stack.Len() != 0 {
				ptr = stack.Back().Value.(*TreeNew).right
			} else {
				break
			}
		}
	}
}

func stackToList(stack *list.List) []int {
	arr := make([]int, 0)
	for e := stack.Back(); e != nil; e = e.Prev() {
		arr = append(arr, e.Value.(*TreeNew).val)
	}
	return arr
}

func (t *TreeNew) RootToLeaf() [][]int {
	stack := list.New()
	allPath := make([][]int, 0)
	ptr := t
	for {
		if ptr != nil {
			stack.PushBack(ptr)
			ptr = ptr.left
		} else {
			if stack.Len() == 0 {
				break
			}
			if stack.Back().Value.(*TreeNew).right == nil {
				elm := stack.Back()
				ptr = elm.Value.(*TreeNew)
				path := stackToList(stack)
				allPath = append(allPath, path)
				stack.Remove(elm)
				for stack.Back().Value.(*TreeNew).right == ptr {
					elm := stack.Back()
					ptr = elm.Value.(*TreeNew)
					path := stackToList(stack)
					allPath = append(allPath, path)
					stack.Remove(elm)
					if stack.Len() == 0 {
						break
					}
				}
			}
			if stack.Len() != 0 {
				ptr = stack.Back().Value.(*TreeNew).right
			} else {
				break
			}
		}
	}
	return allPath
}

func (t *TreeNew) ReverseLevelOrder() {
	queue := list.New()
	stack := list.New()
	queue.PushBack(t)
	for queue.Len() != 0 {
		num := queue.Len()
		for num != 0 {
			elm := queue.Front()
			node := elm.Value.(*TreeNew)
			stack.PushBack(node)
			queue.Remove(elm)
			if node.right != nil {
				queue.PushBack(node.right)
			}
			if node.left != nil {
				queue.PushBack(node.left)
			}
			num -= 1
		}
	}
	fmt.Println("Displaying reverse level order traversal:")
	for stack.Len() != 0 {
		elm := stack.Back()
		node := elm.Value.(*TreeNew)
		fmt.Println(node.val)
		stack.Remove(elm)
	}
}
