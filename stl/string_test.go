package stl

import "testing"

func TestReverseString(t *testing.T) {
	type args struct {
		origin string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "happy",
			args: args{
				origin: "helloworld",
			},
			want: "dlrowolleh",
		},
		{
			name: "empty",
			args: args{
				origin: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseString(tt.args.origin); got != tt.want {
				t.Errorf("ReverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}
