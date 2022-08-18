package stl

import (
	"container/list"
	"reflect"
	"testing"
)

func TestNewStack(t *testing.T) {
	tests := []struct {
		name string
		want *Stack
	}{
		{
			name: "happy",
			want: &Stack{list: list.New().Init()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStack(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Empty(t *testing.T) {
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
			s := &Stack{
				list: tt.fields.list,
			}
			if got := s.Empty(); got != tt.want {
				t.Errorf("Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	listWhoCare := list.New().Init()
	listWhoCare.PushBack("WhoCare")

	listWhoCareTwoElements := list.New().Init()
	listWhoCareTwoElements.PushBack("WhoCare")
	listWhoCareTwoElements.PushBack("WhoCare2")

	type fields struct {
		list *list.List
	}
	tests := []struct {
		name   string
		fields fields
		expectEmpty    bool
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
			s := &Stack{
				list: tt.fields.list,
			}
			s.Pop()
			if got := s.Empty(); got != tt.expectEmpty {
				t.Errorf("AfterPop() = %v, want %v", got, tt.expectEmpty)
			}
		})
	}
}

func TestStack_Push(t *testing.T) {
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				list: tt.fields.list,
			}
			if got := s.Push(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Push() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Size(t *testing.T) {
	type fields struct {
		list *list.List
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				list: tt.fields.list,
			}
			if got := s.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Top(t *testing.T) {
	type fields struct {
		list *list.List
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				list: tt.fields.list,
			}
			if got := s.Top(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Top() = %v, want %v", got, tt.want)
			}
		})
	}
}
