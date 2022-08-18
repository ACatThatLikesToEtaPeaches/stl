package stl

import (
	"reflect"
	"testing"
)

func Test_buildTreeByPreorderAndInorder(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{
				preorder: []int{3,9,20,15,7},
				inorder: []int{9,3,15,20,7},
			},
			want: []int{9,3,15,20,7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildTreeByPreorderAndInorder(tt.args.preorder, tt.args.inorder); !reflect.DeepEqual(InorderTraverse(got), tt.want) {
				t.Errorf("buildTreeByPreorderAndInorder() = %v, want %v", InorderTraverse(got), tt.want)
			}
		})
	}
}
