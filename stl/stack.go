package stl

import "container/list"

type Stack interface {
	Push(v interface{}) interface{}
	Pop()
	Top() interface{}
	Empty() bool
	Size() int
}

// stackListImpl "栈" 基于双向链表实现。LIFO: 前插后删
type stackListImpl struct { list *list.List }

func NewStack() Stack {return &stackListImpl{list: list.New()}}
func (s *stackListImpl) Push(v interface{}) interface{} { return s.list.PushFront(v).Value}
func (s *stackListImpl) Pop() {s.list.Remove(s.list.Front())}
func (s *stackListImpl) Top() interface{} {return s.list.Front().Value}
func (s *stackListImpl) Empty() bool {return s.list.Len() == 0}
func (s *stackListImpl) Size() int {return s.list.Len()}

// stackArrImpl "栈" 基于数组实现。LIFO: 后插前删
type stackArrImpl []interface{}

func NewStackV2() Stack {return &stackListImpl{}}
func (s *stackArrImpl) Push(v interface{}) { *s = append(*s, v) }
func (s *stackArrImpl) Pop() { *s = (*s)[:len(*s)-1] }
func (s *stackArrImpl) Top() interface{} {return (*s)[len(*s)-1]}
func (s *stackArrImpl) Empty() bool {return len(*s) == 0}
func (s *stackArrImpl) Size() int {return len(*s)}