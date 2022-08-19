package stl

import "container/list"

type Stack interface {
	Push(v interface{}) interface{}
	Pop()
	Top() interface{}
	Empty() bool
	Size() int
}

// stackImpl "栈" 结构，基于双向链表实现
type stackImpl struct {
	list *list.List
}

func NewStack() Stack {
	return &stackImpl{list: list.New()}
}

func (s *stackImpl) Push(v interface{}) interface{} {
	return s.list.PushFront(v).Value
}

func (s *stackImpl) Pop() {
	if s.list.Len() == 0 {
		return
	}
	s.list.Remove(s.list.Front())
}

func (s *stackImpl) Top() interface{} {
	if s.list.Len() == 0 {
		return nil
	}
	return s.list.Front().Value
}

func (s *stackImpl) Empty() bool {
	return s.list.Len() == 0
}

func (s *stackImpl) Size() int {
	return s.list.Len()
}
