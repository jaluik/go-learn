package main

import "testing"

func TestTriangle(t *testing.T) {
	test := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{30000, 40000, 50000},
	}

	for _, tt := range test {
		if actual := calcTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calcTriangle(%d, %d);  got: %d; expected: %d", tt.a, tt.b, actual, tt.c)
		}
	}
}

func BenchmarkCalcTriangle(t *testing.B) {
	a, b := 3, 4
	for i := 0; i < t.N; i++ {
		actual := calcTriangle(a, b)
		if actual != 5 {
			t.Errorf("Error result is %d, expected %d", actual, 5)
		}
	}
}
