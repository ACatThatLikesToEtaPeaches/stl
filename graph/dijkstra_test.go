package graph

import (
	"datastruct_algorithm/stl"
	"reflect"
	"testing"
)

func TestDijkstra(t *testing.T) {
	type args struct {
		weight [][]int
		n      int
		start  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "happy",
			args: args{
				weight: [][]int{{1,0,1},{1,2,1},{2,3,1}},
				n: 4,
				start: 1,
			},
			want: []int{1, 0, 1, 2},
		},
		{
			name: "not all point can reach",
			args: args{
				weight: [][]int{{0,1,1},{1,2,1},{2,3,1}},
				n: 4,
				start: 1,
			},
			want: []int{stl.INF, 0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Dijkstra(tt.args.weight, tt.args.n, tt.args.start); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dijkstra() = %v, want %v", got, tt.want)
			}
		})
	}
}
