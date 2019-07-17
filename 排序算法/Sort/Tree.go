package Sort

import (
	"container/list"
	"fmt"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (root *Node) PreTravesal() {
	if root == nil {
		return
	}

	s := stack.NewStack()
	s.push(root)

	for !s.Empty() {
		cur := s.Pop().(*Node)
		fmt.Println(cur.Val)

		if cur.Right != nil {
			s.Push(cur.Right)
		}
		if cur.Left != nil {
			s.Push(cur.Left)
		}
	}
}

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	list := list.New()
	return &Stack{list}
}
func (stack *Stack) Push(value interface{}) {
	stack.list.PushBack(value)
}
func (stack *Stack) Pop() interface{} {
	if e := stack.list.Back(); e != nil {
		stack.list.Remove(e)
		return e.Value
	}

	return nil
}
func (stack *Stack) Len() int {
	return stack.list.Len()
}

func (stack *Stack) Empty() bool {
	return stack.Len() == 0
}
