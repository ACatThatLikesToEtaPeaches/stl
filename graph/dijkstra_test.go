package graph

import (
	"math"
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
			want: []int{math.MaxInt32, 0, 1, 2},
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

func TestBiDijkstra(t *testing.T) {
	type args struct {
		weight [][]int
		n      int
		start  int
		end    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "happy",
			args: args{
				weight: [][]int{{1,0,1},{1,2,1},{2,3,1}},
				n: 4,
				start: 1,
				end: 2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BiDijkstra(tt.args.weight, tt.args.n, tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("BiDijkstra() = %v, want %v", got, tt.want)
			}
		})
	}
}