package main

import "testing"

func TestCountLoop(t *testing.T) {
	var tests1 = []struct {
		inputA int
		want   int
	}{
		{3, 1},
		{7, 2},
		{9, 3},
	}
	for _, test := range tests1 {
		if got := countLoop(test.inputA); got != test.want {
			t.Errorf("countLoop(%v) = %v", test.inputA, got)
		}
	}
}

func BenchmarkCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countLoop(8)
	}
}
