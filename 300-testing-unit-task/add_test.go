package add

import "testing"

func assert(t *testing.T, want, got int) {
	if want != got {
		t.Errorf("Want %d but got %d", want, got)
	}
}

func TestAdd_AddsTwoNumbers(t *testing.T) {
	result := Add(1, 1)

	assert(t, 2, result)
}

func TestAdd(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		isErros bool
	}{
		{"dwa plus dwa = cztery ", args{[]int{2, 2}}, 4, false},
		{"dwa plus dwa = sto ", args{[]int{2, 2}}, 100, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a...); got != tt.want && !tt.isErros {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
