package app

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func createTreeNode(val int) *TreeNode {
	tn := TreeNode{}
	tn.val = val
	tn.left = nil
	tn.right = nil
	return &tn
}

func ConstructTree(nums []int, rootVal int) *TreeNode {
	node := createTreeNode(rootVal)
	for _, num := range nums {
		node.insertTree(num)
	}
	return node
}

func (tn *TreeNode) insertTree(val int) {
	if tn.val > val {
		if tn.left == nil {
			tn.left = createTreeNode(val)
		} else {
			tn.left.insertTree(val)
		}
	} else if tn.val < val {
		if tn.right == nil {
			tn.right = createTreeNode(val)
		} else {
			tn.right.insertTree(val)
		}
	} else {
		fmt.Println("Operation invalid:")
	}
}

func (tn *TreeNode) LevelOrderTrav() {
	queue := list.New()
	queue.PushBack(tn)
	for queue.Len() != 0 {
		num := queue.Len()
		for num != 0 {
			elm := queue.Front()
			node := elm.Value.(*TreeNode)
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

func (tn *TreeNode) Postorder() {
	stack := list.New()
	ptr := tn
	for {
		if ptr != nil {
			stack.PushBack(ptr)
			ptr = ptr.left
		} else {
			if stack.Len() == 0 {
				break
			}
			if stack.Back().Value.(*TreeNode).right == nil {
				elm := stack.Back()
				ptr = elm.Value.(*TreeNode)
				fmt.Println(ptr.val)
				stack.Remove(elm)
				for stack.Back().Value.(*TreeNode).right == ptr {
					elm = stack.Back()
					ptr = elm.Value.(*TreeNode)
					fmt.Println(ptr.val)
					stack.Remove(elm)
					if stack.Len() == 0 {
						break
					}
				}
			}
			if stack.Len() != 0 {
				ptr = stack.Back().Value.(*TreeNode).right
			} else {
				break
			}
		}

	}
}
