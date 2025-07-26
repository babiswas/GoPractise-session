package app

import (
	"container/list"
	"fmt"
)

type Tree struct {
	val   int
	left  *Tree
	right *Tree
}

func CreateTree(val int) *Tree {
	t := Tree{}
	t.val = val
	t.left = nil
	t.right = nil
	return &t
}

func (t *Tree) Debug() {
	fmt.Println(t.val)
	if t.left != nil {
		fmt.Println(t.left.val)
	}
	if t.right != nil {
		fmt.Println(t.right.val)
	}
}

func (t *Tree) InsertNode(num int) {
	if t.val > num {
		if t.left == nil {
			t.left = CreateTree(num)
			return
		} else {
			t.left.InsertNode(num)
		}
	} else if t.val < num {
		if t.right == nil {
			t.right = CreateTree(num)
			return
		} else {
			t.right.InsertNode(num)
		}
	} else {
		fmt.Println("Operation not allowed.")
	}
}

func TreeBuild(nums []int, rootElement int) *Tree {
	node := CreateTree(rootElement)
	for _, num := range nums {
		node.InsertNode(num)
	}
	return node
}

func (t *Tree) LevelOrderT() {
	queue := list.New()
	queue.PushBack(t)
	for queue.Len() != 0 {
		num := queue.Len()
		for num != 0 {
			elm := queue.Front()
			node := elm.Value.(*Tree)
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

func (t *Tree) Inorder() {
	stack := list.New()
	ptr := t
	for {
		if ptr != nil {
			stack.PushBack(ptr)
			ptr = ptr.left
			continue
		}
		if stack.Len() != 0 {
			elm := stack.Back()
			node := elm.Value.(*Tree)
			fmt.Println(node.val)
			stack.Remove(elm)
			ptr = node.right
		} else {
			break
		}

	}
}
