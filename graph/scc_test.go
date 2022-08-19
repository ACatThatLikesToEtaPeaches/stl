package graph

import (
	"reflect"
	"testing"
)

func TestKosarajuCalcSCC(t *testing.T) {
	type args struct {
		bias [][]int
		n    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "happy",
			args: args{
				bias: [][]int{
					{0,1},{0,4},
					{1,0},{1,8},
					{2,1},{2,4},{2,7},
					{3,4},
					{4,3},
					{5,0},{5,6},
					{7,9},{7,4},
					{8,5},
					{9,2},
				},
				n: 10,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KosarajuCalcSCC(tt.args.bias, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KosarajuCalcSCC() = %v, want %v", got, tt.want)
			}
		})
	}
}
