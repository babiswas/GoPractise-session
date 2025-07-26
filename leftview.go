package app

import (
	"container/list"
	"fmt"
)

type LeftView struct {
	val   int
	left  *LeftView
	right *LeftView
}

func CreateTreeNode(val int) *LeftView {
	leftView := LeftView{}
	leftView.val = val
	leftView.left = nil
	leftView.right = nil
	return &leftView
}

func (l *LeftView) InsertLetNode(val int) {
	if l.val > val {
		if l.left == nil {
			l.left = CreateTreeNode(val)
		} else {
			l.left.InsertLetNode(val)
		}
	} else {
		if l.right == nil {
			l.right = CreateTreeNode(val)
		} else {
			l.right.InsertLetNode(val)
		}
	}
}

func BuildLeftView(nums []int, root int) *LeftView {
	node := CreateTreeNode(root)
	for _, num := range nums {
		node.InsertLetNode(num)
	}
	return node
}

func (l *LeftView) ShowLeftView() {
	queue := list.New()
	queue.PushBack(l)
	for queue.Len() != 0 {
		num := queue.Len()
		fmt.Println(queue.Front().Value.(*LeftView).val)
		for num != 0 {
			elm := queue.Front()
			node := elm.Value.(*LeftView)
			queue.Remove(elm)
			if node.left != nil {
				queue.PushBack(node.left)
			}
			if node.right != nil {
				queue.PushBack(node.right)
			}
			num -= 1
		}
	}
}

func (l *LeftView) LevelView() {
	queue := list.New()
	queue.PushBack(l)
	for queue.Len() != 0 {
		num := queue.Len()
		for num != 0 {
			elm := queue.Front()
			node := elm.Value.(*LeftView)
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

func (l *LeftView) RightView() {
	queue := list.New()
	queue.PushBack(l)
	for queue.Len() != 0 {
		num := queue.Len()
		fmt.Println(queue.Back().Value.(*LeftView).val)
		for num != 0 {
			elm := queue.Front()
			node := elm.Value.(*LeftView)
			queue.Remove(elm)
			if node.left != nil {
				queue.PushBack(node.left)
			}
			if node.right != nil {
				queue.PushBack(node.right)
			}
			num -= 1
		}
	}
}
