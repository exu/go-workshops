package add

import "testing"

func TestAdd(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"dwa plus dwa = cztery ", args{[]int{2, 2}}, 4},
		{"dwa plus dwa = sto ", args{[]int{2, 2}}, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a...); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
