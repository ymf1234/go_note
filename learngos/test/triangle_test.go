package main

import "testing"
import "learngos/test/calcTriangle"

func TestTriangle(t *testing.T) {
	tests := []struct{a, b, c int} {
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 37},
		{30000, 40000, 50000},
	}

	for _, tt := range tests {
		if actual := calcTriangle.CalcTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calcTriangle(%d, %d); got %d; expected %d", tt.a, tt.b, actual, tt.c)
		}
	}

}

func BenchmarkSubstr(b *testing.B) {
	a := 30000
	d := 40000
	c := 50000
	for i := 0; i < b.N; i++ {
		if actual := calcTriangle.CalcTriangle(a, d); actual != c {
			b.Errorf("calcTriangle(%d, %d); got %d; expected %d", a, d, actual, c)
		}
	}

}
