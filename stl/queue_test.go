package stl

import (
	"container/list"
	"reflect"
	"testing"
)

func TestNewQueue(t *testing.T) {
	tests := []struct {
		name string
		want *Queue
	}{
		{
			name: "happy",
			want: &Queue{list: list.New().Init()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewQueue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Push(t *testing.T) {
	listWhoCare := list.New().Init()
	listWhoCare.PushBack("firstElement")
	listWhoCare.PushBack("secondElement")
	type fields struct {
		list *list.List
	}
	type args struct {
		v interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		{
			name: "happy",
			fields: fields{
				list: listWhoCare,
			},
			args: args{
				v: "thirdElement",
			},
			want: "thirdElement",
		},
		{
			name: "empty",
			fields: fields{
				list: list.New().Init(),
			},
			args: args{
				v: "firstElement",
			},
			want: "firstElement",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue{
				list: tt.fields.list,
			}
			q.Push(tt.args.v)
		})
	}
}

func TestQueue_Pop(t *testing.T) {
	listWhoCare := list.New().Init()
	listWhoCare.PushBack("WhoCare")

	listWhoCareTwoElements := list.New().Init()
	listWhoCareTwoElements.PushBack("WhoCare")
	listWhoCareTwoElements.PushBack("WhoCare2")
	type fields struct {
		list *list.List
	}
	tests := []struct {
		name        string
		fields      fields
		expectEmpty bool
	}{
		{
			name: "one element",
			fields: fields{
				list: listWhoCare,
			},
			expectEmpty: true,
		},
		{
			name: "two element",
			fields: fields{
				list: listWhoCareTwoElements,
			},
			expectEmpty: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue{
				list: tt.fields.list,
			}
			q.Pop()
			if got := q.Empty(); got != tt.expectEmpty {
				t.Errorf("AfterPop() = %v, want %v", got, tt.expectEmpty)
			}
		})
	}
}

func TestQueue_Empty(t *testing.T) {
	listWhoCare := list.New().Init()
	listWhoCare.PushBack("WhoCare")
	type fields struct {
		list *list.List
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "empty",
			fields: fields{
				list: &list.List{},
			},
			want: true,
		},
		{
			name: "not empty",
			fields: fields{
				list: listWhoCare,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue{
				list: tt.fields.list,
			}
			if got := q.Empty(); got != tt.want {
				t.Errorf("Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Front(t *testing.T) {
	listWhoCare := list.New().Init()
	listWhoCare.PushBack("firstElement")
	listWhoCare.PushBack("secondElement")

	type fields struct {
		list *list.List
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name: "empty",
			fields: fields{
				list: list.New().Init(),
			},
			want: nil,
		},
		{
			name: "happy",
			fields: fields{
				list: listWhoCare,
			},
			want: "firstElement",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue{
				list: tt.fields.list,
			}
			if got := q.Front(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Front() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Size(t *testing.T) {
	listWhoCare := list.New().Init()
	listWhoCare.PushBack("firstElement")
	listWhoCare.PushBack("secondElement")

	listNumber := list.New().Init()
	listNumber.PushBack(1)
	listNumber.PushBack(2)
	listNumber.PushBack(3)

	type fields struct {
		list *list.List
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "queue of number",
			fields: fields{
				list: listNumber,
			},
			want: 3,
		},
		{
			name: "queue of string",
			fields: fields{
				list: listWhoCare,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue{
				list: tt.fields.list,
			}
			if got := q.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}
