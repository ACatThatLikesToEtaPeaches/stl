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
				weight: [][]int{
					{1,0,1},
					{1,2,1},
					{2,3,1},
				},
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
		{
			name: "luogu p3371",
			args: args{
				weight: [][]int{{0, 1, 2}, {1, 2, 2}, {1, 3, 1}, {0, 2, 5}, {2, 3, 3}, {0, 3, 4}},
				n: 4,
				start: 0,
				end: 3,
			},
			want: 3,
		},
		{
			name: "test case from blog: https://www.cnblogs.com/lang5230/p/5510908.html",
			args: args{
				weight: [][]int{
					{0, 3, 10},
					{0, 2, 50},
					{3, 1, 20},
					{1, 2, 10},
					{2, 4, 40},
					{4, 0, 20},
					{4, 1, 30},
				},
				n: 5,
				start: 0,
				end: 4,
			},
			want: 80,
		},
		{
			name: "test case0->2 from blog: https://www.cnblogs.com/lang5230/p/5510908.html",
			args: args{
				weight: [][]int{
					{0, 3, 10},
					{0, 2, 50},
					{3, 1, 20},
					{1, 2, 10},
					{2, 4, 40},
					{4, 0, 20},
					{4, 1, 30},
				},
				n: 5,
				start: 0,
				end: 2,
			},
			want: 40,
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