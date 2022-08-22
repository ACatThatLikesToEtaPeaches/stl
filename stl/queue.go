package stl

import "container/list"

// Queue "队列"结构，基于双向链表实现
type Queue struct{ list *list.List }

func NewQueue() *Queue              { return &Queue{list: list.New().Init()} }

func (q *Queue) Push(v interface{}) { q.list.PushBack(v) }
func (q *Queue) Pop()               { q.list.Remove(q.list.Front()) }
func (q *Queue) Front() interface{} { return q.list.Front().Value }
func (q *Queue) Back() interface{}  { return q.list.Back().Value }
func (q *Queue) Empty() bool        { return q.list.Len() == 0 }
func (q *Queue) Size() int          { return q.list.Len() }
