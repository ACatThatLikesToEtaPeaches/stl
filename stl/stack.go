package stl

import "container/list"

type StackInterface interface {
	Push(v interface{}) interface{}
	Pop()
	Top() interface{}
	Empty() bool
	Size() int
}

// Stack "栈" 结构，基于双向链表实现
type Stack struct {
	list *list.List
}

func NewStack() StackInterface {
	return &Stack{list: list.New()}
}

func (s *Stack) Push(v interface{}) interface{} {
	return s.list.PushFront(v).Value
}

func (s *Stack) Pop() {
	if s.list.Len() == 0 {
		return
	}
	s.list.Remove(s.list.Front())
}

func (s *Stack) Top() interface{} {
	if s.list.Len() == 0 {
		return nil
	}
	return s.list.Front().Value
}

func (s *Stack) Empty() bool {
	return s.list.Len() == 0
}

func (s *Stack) Size() int {
	return s.list.Len()
}