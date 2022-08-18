package stl

import (
	"container/heap"
	"reflect"
	"testing"
)

func TestNewPriorityQueue(t *testing.T) {
	tests := []struct {
		name string
		want PriorityQueue
	}{
		{
			name: "happy",
			want: &PriorityQueueImpl{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPriorityQueue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPriorityQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPriorityQueueBySlice(t *testing.T) {
	data := []*Item{
		{priority: 5},
		{priority: 4},
		{priority: 3},
		{priority: 2},
		{priority: 1},
	}
	type args struct {
		items []int
	}
	tests := []struct {
		name string
		args args
		want PriorityQueue
	}{
		{
			name: "happy",
			args: args{
				items: []int{1,2,3,4,5},
			},
			want: &PriorityQueueImpl{
				data: data,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := NewPriorityQueueBySlice(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				got := NewPriorityQueueBySlice(tt.args.items)
				i := 0
				for got.Size() > 0 {
					i++
					gotOne := got.Top()
					if !got.Empty() { got.Pop() }
					wantOne := tt.want.Top()
					if !tt.want.Empty() { tt.want.Pop() }
					if !reflect.DeepEqual(gotOne, wantOne) {
						t.Errorf("NewPriorityQueueBySlice(), Index: %v, got = %v, want %v", i, gotOne, wantOne)
					}
				}
			//}
		})
	}
}

func TestPriorityQueueImpl_Push(t *testing.T) {
	type fields struct {
		data HeapImpl
	}
	type args struct {
		v int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "happy push one element to empty heap",
			fields: fields{
				data: HeapImpl{},
			},
			args: args{
				v: 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pq := &PriorityQueueImpl{
				data: tt.fields.data,
			}
			pq.Push(tt.args.v)
			if !reflect.DeepEqual(pq.Top(), tt.args.v) {
				t.Errorf("Push() = %v, want %v", pq.Top(), tt.args.v)
			}
		})
	}
}

func TestPriorityQueueImpl_Pop(t *testing.T) {
	data := []*Item{
		{priority: 1},
		{priority: 2},
		{priority: 3},
		{priority: 4},
		{priority: 5},
	}
	type fields struct {
		data HeapImpl
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "happy, pop one element",
			fields: fields{
				data: data,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pq := &PriorityQueueImpl{
				data: tt.fields.data,
			}
			pq.Pop()
			if !reflect.DeepEqual(pq.Top(), 5) {
				t.Errorf("Push() = %v, want %v", pq.Top(), 5)
			}
		})
	}
}

func TestHeapImpl_Len(t *testing.T) {
	data := []*Item{
		{priority: 1},
		{priority: 2},
		{priority: 3},
		{priority: 4},
		{priority: 5},
	}
	tests := []struct {
		name string
		hp   HeapImpl
		want int
	}{
		{
			name: "happy",
			hp: HeapImpl(data),
			want: 5,
		},
		{
			name: "empty",
			hp: HeapImpl{},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hp.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeapImpl_Less(t *testing.T) {
	data := []*Item{
		{priority: 1},
		{priority: 2},
		{priority: 3},
		{priority: 4},
		{priority: 5},
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		hp   HeapImpl
		args args
		want bool
	}{
		{
			name: "happy",
			hp: HeapImpl(data),
			args: args{
				i: 0,
				j: 4,
			},
			want: false,
		},
		{
			name: "empty",
			hp: HeapImpl(data),
			args: args{
				i: 0,
				j: 4,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hp.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}


func TestPriorityQueueImpl_Empty(t *testing.T) {
	type fields struct {
		data HeapImpl
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "empty",
			fields: fields{
				data: HeapImpl{},
			},
			want: true,
		},
		{
			name: "not empty",
			fields: fields{
				data: HeapImpl([]*Item{{priority: 1}}),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pq := &PriorityQueueImpl{
				data: tt.fields.data,
			}
			if got := pq.Empty(); got != tt.want {
				t.Errorf("Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPriorityQueueImpl_Size(t *testing.T) {
	data := []*Item{
		{priority: 1},
		{priority: 2},
		{priority: 3},
		{priority: 4},
		{priority: 5},
	}
	type fields struct {
		data HeapImpl
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "empty",
			fields: fields{
				data: HeapImpl{},
			},
			want: 0,
		},
		{
			name: "happy",
			fields: fields{
				data: data,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pq := &PriorityQueueImpl{
				data: tt.fields.data,
			}
			if got := pq.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPriorityQueueImpl_Top(t *testing.T) {
	data := []*Item{
		{priority: 1},
		{priority: 2},
		{priority: 3},
		{priority: 4},
		{priority: 5},
	}
	type fields struct {
		data HeapImpl
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "happy",
			fields: fields{
				data: data,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pq := &PriorityQueueImpl{
				data: tt.fields.data,
			}
			heap.Init(&pq.data)
			if got := pq.Top(); got != tt.want {
				t.Errorf("Top() = %v, want %v", got, tt.want)
			}
		})
	}
}