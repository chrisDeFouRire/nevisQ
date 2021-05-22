package main

import "testing"

// TestCountSubSequences tests a few values of CountSubSequences
func TestCountSubSequences(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test0", args{3, 1}, 3},
		{"test00", args{3, 3}, 1},
		{"test1", args{4, 2}, 6},
		{"test2", args{4, 1}, 4},
		{"test3", args{4, 3}, 4},
		{"test4", args{4, 4}, 1},
		{"test52", args{5, 2}, 10},
		{"test53", args{5, 3}, 10},
		{"test54", args{5, 4}, 5},
		{"test54", args{6, 2}, 15},
		{"test54", args{6, 3}, 20},
		{"test54", args{6, 4}, 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountSubSequences(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("CountSeq() = %v, want %v", got, tt.want)
			}
		})
	}
}

//TestItsCnk checks the result of question 1 is Cnk(m,n)
func TestItsCnk(t *testing.T) {
	for m := 15; m <= 15; m++ {
		for n := 1; n <= m; n++ {
			if Cnk(m, n) != CountSubSequences(m, n) {
				t.Errorf("No it's not Cnk")
			}
		}
	}
}
