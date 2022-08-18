package stl

import (
	"container/heap"
)

// PriorityQueue "优先级队列" 结构
type PriorityQueue interface {
	Push(v int)
	Pop()
	Empty() bool
	Size() int
	Top() int
}

// PriorityQueueImpl 基于数组"堆调整"实现
type PriorityQueueImpl struct { data HeapImpl }
func NewPriorityQueue() PriorityQueue { return &PriorityQueueImpl{} }
func NewPriorityQueueBySlice(items []int) PriorityQueue {
	hp := make(HeapImpl, len(items))
	for i, v := range items {
		hp[i] = &Item{priority: v}
	}
	heap.Init(&hp)
	return &PriorityQueueImpl{data: hp}
}

func (pq *PriorityQueueImpl) Push(v int) { heap.Push(&pq.data, v) }
func (pq *PriorityQueueImpl) Pop() { heap.Pop(&pq.data) }
func (pq *PriorityQueueImpl) Empty() bool {return len(pq.data) == 0}
func (pq *PriorityQueueImpl) Size() int {return len(pq.data)}
func (pq *PriorityQueueImpl) Top() int { return pq.data[0].priority }


// HeapImpl 基于数组调整实现
type HeapImpl []*Item

type Item struct {
	value int
	priority int         // The priority of the item in the queue.
}

func (hp HeapImpl) Len() int { return len(hp) }
func (hp HeapImpl) Less(i, j int) bool { return hp[i].priority > hp[j].priority } // Less 默认升序，">" 表示降序
func (hp HeapImpl) Swap(i, j int) { hp[i], hp[j] = hp[j], hp[i]}
func (hp *HeapImpl) Push(v interface{}) { *hp = append(*hp, &Item{priority: v.(int)})}
func (hp *HeapImpl) Pop() (v interface{}) { a := *hp; *hp, v = a[:len(a)-1], a[len(a)-1]; return v}