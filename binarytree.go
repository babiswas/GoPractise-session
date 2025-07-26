package app

import (
	"container/list"
	"fmt"
)

type Btree struct {
	val   int
	left  *Btree
	right *Btree
}

func createBtree(val int) *Btree {
	b := Btree{}
	b.val = val
	b.left = nil
	b.right = nil
	return &b
}

func (b *Btree) insertBtree(val int) {
	if b.val > val {
		if b.left == nil {
			b.left = createBtree(val)
		} else {
			b.left.insertBtree(val)
		}
	} else if b.val < val {
		if b.right == nil {
			b.right = createBtree(val)
		} else {
			b.right.insertBtree(val)
		}
	}
}

func BuildBtree(nums []int, rootVal int) *Btree {
	node := createBtree(rootVal)
	for _, num := range nums {
		node.insertBtree(num)
	}
	return node
}

func (b *Btree) LTrav() {
	queue := list.New()
	queue.PushBack(b)
	for queue.Len() != 0 {
		num := queue.Len()
		for num != 0 {
			elm := queue.Front()
			node := elm.Value.(*Btree)
			fmt.Printf("%d ", node.val)
			if node.left != nil {
				queue.PushBack(node.left)
			}
			if node.right != nil {
				queue.PushBack(node.right)
			}
			queue.Remove(elm)
			num = num - 1
		}
		fmt.Println("")
	}
}

func (b *Btree) DDlist() {
	stack := list.New()
	var head *Btree
	var prev *Btree
	ptr := b
	for {
		if ptr != nil {
			stack.PushBack(ptr)
			ptr = ptr.left
			continue
		}
		if stack.Len() != 0 {
			elm := stack.Back()
			ptr = elm.Value.(*Btree)
			stack.Remove(elm)
			if head == nil {
				head = ptr
				ptr.left = head
			} else {
				if ptr.left == prev {
					prev.right = ptr
				} else if prev.right == ptr {
					ptr.left = prev
				} else if ptr.left != prev && prev.right != ptr {
					ptr.left = prev
					prev.right = ptr
				}
			}
			prev = ptr
			ptr = ptr.right
		}
		if stack.Len() == 0 {
			break
		}
	}
	fmt.Println("Traversing the doubly linked list:")
	temp := head
	for temp != nil {
		fmt.Println(temp.val)
		temp = temp.right
	}
}
