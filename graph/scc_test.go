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
		wantComps [][]int
		wantSccid []int
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
			wantComps: [][]int{{2, 7, 9}, {0, 1, 5, 8}, {3, 4}, {6}},
			wantSccid: []int{2, 2, 3, 1, 1, 2, 0, 3, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotComps, gotSccid := KosarajuCalcSCC(tt.args.bias, tt.args.n)
			if !reflect.DeepEqual(gotComps, tt.wantComps) {
				t.Errorf("KosarajuCalcSCC() gotComps = %v, want %v", gotComps, tt.wantComps)
			}
			if !reflect.DeepEqual(gotSccid, tt.wantSccid) {
				t.Errorf("KosarajuCalcSCC() gotSccid = %v, want %v", gotSccid, tt.wantSccid)
			}

			gotComps2, _ := TarjanCalcSCC(tt.args.bias, tt.args.n)
			if !reflect.DeepEqual(len(gotComps2), len(tt.wantComps)) {
				t.Errorf("TarjanCalcSCC() gotComps = %v, want %v", gotComps2, tt.wantComps)
			}
		})
	}
}
