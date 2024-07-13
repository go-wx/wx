package tests

import "testing"

func TestCloseEnough(t *testing.T) {
	tt := []struct {
		name      string
		a         float64
		b         float64
		tolerance float64
		expected  bool
	}{
		{name: "equal", a: 1, b: 1, tolerance: 0, expected: true},
		{name: "within", a: 1, b: 1.1, tolerance: 0.1, expected: true},
		{name: "outside", a: 1, b: 1.1, tolerance: 0.05, expected: false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := CloseEnough(tc.a, tc.b, tc.tolerance)
			if actual != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
