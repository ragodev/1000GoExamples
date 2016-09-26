package main

import "testing"

func TestCountLoop(t *testing.T) {
	var tests1 = []struct {
		inputA int
		want   float64
	}{
		{3, 3},
	}
	for _, test := range tests1 {
		if got := countLoop(test.inputA); got != test.want {
			t.Errorf("countLoop(%v) = %v", test.inputA, got)
		}
	}
}

func BenchmarkFoo1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countLoop(8)
	}
}
