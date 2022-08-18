package stl


// CircularQueue "循环队列" 结构
type CircularQueue struct {
	front, tail int
	elements    []int
}

func NewCircularQueue(k int) *CircularQueue {
	// 在循环队列中，当队列为空，可知 front=rear；而当所有队列空间全占满时，也有 front=rear。
	// 为了区别这两种情况，假设队列使用的数组有 capacity 个存储空间，则此时规定循环队列最多只能有capacity−1 个队列元素，
	// 当循环队列中只剩下一个空存储单元时，则表示队列已满。
	//	 **根据以上可知，队列判空的条件是 front=rear，而队列判满的条件是 front=(rear+1) mod capacity。
	//   **对于一个固定大小的数组，只要知道队尾 rear 与队首 front，即可计算出队列当前的长度：(rear−front+capacity) mod capacity
	return &CircularQueue{elements: make([]int, k+1)}
}

func (q *CircularQueue) EnQueue(value int) bool {
	if q.IsFull() {
		return false
	}
	q.elements[q.tail] = value
	q.tail = (q.tail + 1) % len(q.elements)
	return true
}

func (q *CircularQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}
	q.front = (q.front + 1) % len(q.elements)
	return true
}

func (q CircularQueue) Front() int {
	if q.IsEmpty() {
		return -1
	}
	return q.elements[q.front]
}

func (q CircularQueue) Rear() int {
	if q.IsEmpty() {
		return -1
	}
	return q.elements[(q.tail-1+len(q.elements))%len(q.elements)]
}

func (q CircularQueue) IsEmpty() bool {
	return q.tail == q.front
}

func (q CircularQueue) IsFull() bool {
	return (q.tail+1)%len(q.elements) == q.front
}
