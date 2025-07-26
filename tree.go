package app

import (
	"container/list"
	"fmt"
)

type BStree struct {
	val   int
	left  *BStree
	right *BStree
}

func CreateNode(val int) *BStree {
	b := BStree{}
	b.val = val
	b.left = nil
	b.right = nil
	return &b
}

func (b *BStree) Insert(val int) {
	if b.val > val {
		if b.left == nil {
			b.left = CreateNode(val)
			return
		} else {
			b.left.Insert(val)
		}
	} else if b.val < val {
		if b.right == nil {
			b.right = CreateNode(val)
			return
		} else {
			b.right.Insert(val)
		}
	} else {
		fmt.Println("Operation not supported.")
	}
}

func (b *BStree) LevelOrder() {
	queue := list.New()
	queue.PushBack(b)
	for queue.Len() != 0 {
		elm := queue.Front()
		nd := elm.Value.(*BStree)
		fmt.Println(nd.val)
		queue.Remove(elm)
		if nd.left != nil {
			queue.PushBack(nd.left)
		}
		if nd.right != nil {
			queue.PushBack(nd.right)
		}
	}
}

func BuildTree(values []int, rootVal int) *BStree {
	node := CreateNode(rootVal)
	for _, val := range values {
		node.Insert(val)
	}
	return node
}

func (b *BStree) PreorderTrav() {
	if b != nil {
		fmt.Println(b.val)
		if b.left != nil {
			b.left.PreorderTrav()
		}
		if b.right != nil {
			b.right.PreorderTrav()
		}
	} else {
		return
	}
}

func (b *BStree) PostorderTrav() {
	if b.left != nil {
		b.left.PostorderTrav()
	}
	if b.right != nil {
		b.right.PostorderTrav()
	}
	fmt.Println(b.val)
}

func (b *BStree) InorderTrav() {
	if b.left != nil {
		b.left.InorderTrav()
	}
	fmt.Println(b.val)
	if b.right != nil {
		b.right.InorderTrav()
	}
}

func (b *BStree) Height() int {
	height := 0
	queue := list.New()
	queue.PushBack(b)
	for queue.Len() != 0 {
		num := queue.Len()
		height += 1
		for num != 0 {
			elm := queue.Front()
			nd := elm.Value.(*BStree)
			if nd.left != nil {
				queue.PushBack(nd.left)
			}
			if nd.right != nil {
				queue.PushBack(nd.right)
			}
			queue.Remove(elm)
			num -= 1
		}
		if queue.Len() == 0 {
			break
		}
	}
	return height
}

func (b *BStree) PreorderTraversal() {
	stack := list.New()
	stack.PushBack(b)
	for stack.Len() != 0 {
		elm := stack.Back()
		node := elm.Value.(*BStree)
		fmt.Println(node.val)
		stack.Remove(elm)
		if node.right != nil {
			stack.PushBack(node.right)
		}
		if node.left != nil {
			stack.PushBack(node.left)
		}
	}
}
