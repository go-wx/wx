package wx

import (
	"github.com/go-wx/wx/internal/tests"
	"testing"
)

func TestTempR_C(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
	}{
		{name: "absolute zero", measurement: 0, expected: -273.15},
		{name: "freezing", measurement: 491.67, expected: 0},
		{name: "boiling", measurement: 671.67, expected: 100},
		{name: "body", measurement: 558.87, expected: 37.33333333333334},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tr := TempR{}
			if err := tr.Set(tc.measurement); err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			actual := tr.C()
			if !tests.CloseEnough(actual, tc.expected, tests.Tolerance) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestTempR_F(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
	}{
		{name: "absolute zero", measurement: 0, expected: -459.67},
		{name: "freezing", measurement: 491.67, expected: 32},
		{name: "boiling", measurement: 671.67, expected: 212},
		{name: "body", measurement: 558.27, expected: 98.6},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tr := TempR{}
			if err := tr.Set(tc.measurement); err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			actual := tr.F()
			if !tests.CloseEnough(actual, tc.expected, tests.Tolerance) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestTempR_K(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
	}{
		{name: "absolute zero", measurement: 0, expected: 0},
		{name: "freezing", measurement: 491.67, expected: 273.15},
		{name: "boiling", measurement: 671.67, expected: 373.15},
		{name: "body", measurement: 558.87, expected: 310.48333333333335},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tr := TempR{}
			if err := tr.Set(tc.measurement); err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			actual := tr.K()
			if !tests.CloseEnough(actual, tc.expected, tests.Tolerance) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestTempR_R(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
	}{
		{name: "absolute zero", measurement: 0, expected: 0},
		{name: "freezing", measurement: 491.67, expected: 491.67},
		{name: "boiling", measurement: 671.67, expected: 671.67},
		{name: "body", measurement: 558.87, expected: 558.87},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tr := TempR{}
			if err := tr.Set(tc.measurement); err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			actual := tr.R()
			if !tests.CloseEnough(actual, tc.expected, tests.Tolerance) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestTempR_Set(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		error       bool
	}{
		{name: "absolute zero", measurement: 0, expected: 0},
		{name: "freezing", measurement: 491.67, expected: 491.67},
		{name: "boiling", measurement: 671.67, expected: 671.67},
		{name: "body", measurement: 558.87, expected: 558.87},
		{name: "below zero", measurement: -1, expected: 0, error: true},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tr := TempR{}
			err := tr.Set(tc.measurement)
			if tc.error && err == nil {
				t.Errorf("expected error, got nil")
			}

			if !tc.error && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !tc.error {
				actual := tr.R()
				if !tests.CloseEnough(actual, tc.expected, tests.Tolerance) {
					t.Errorf("expected %v, got %v", tc.expected, actual)
				}

			}
		})
	}
}

func TestTempR_ToC(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
	}{
		{name: "absolute zero", measurement: 0, expected: -273.15},
		{name: "freezing", measurement: 491.67, expected: 0},
		{name: "boiling", measurement: 671.67, expected: 100},
		{name: "body", measurement: 558.87, expected: 37.33333333333334},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tr := TempR{}
			if err := tr.Set(tc.measurement); err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			tempC := tr.ToC()
			actual := tempC.C()
			if !tests.CloseEnough(actual, tc.expected, tests.Tolerance) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestTempR_ToF(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
	}{
		{name: "absolute zero", measurement: 0, expected: -459.67},
		{name: "freezing", measurement: 491.67, expected: 32},
		{name: "boiling", measurement: 671.67, expected: 212},
		{name: "body", measurement: 558.27, expected: 98.6},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tr := TempR{}
			if err := tr.Set(tc.measurement); err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			tempF := tr.ToF()
			actual := tempF.F()
			if !tests.CloseEnough(actual, tc.expected, tests.Tolerance) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestTempR_ToK(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
	}{
		{name: "absolute zero", measurement: 0, expected: 0},
		{name: "freezing", measurement: 491.67, expected: 273.15},
		{name: "boiling", measurement: 671.67, expected: 373.15},
		{name: "body", measurement: 558.87, expected: 310.48333333333335},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tr := TempR{}
			if err := tr.Set(tc.measurement); err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			tempK := tr.ToK()
			actual := tempK.K()
			if !tests.CloseEnough(actual, tc.expected, tests.Tolerance) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestTempR_ToR(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
	}{
		{name: "absolute zero", measurement: 0, expected: 0},
		{name: "freezing", measurement: 491.67, expected: 491.67},
		{name: "boiling", measurement: 671.67, expected: 671.67},
		{name: "body", measurement: 558.87, expected: 558.87},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tr := TempR{}
			if err := tr.Set(tc.measurement); err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			tempR := tr.ToR()
			actual := tempR.R()
			if !tests.CloseEnough(actual, tc.expected, tests.Tolerance) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestTempR_Units(t *testing.T) {
	tr := TempR{}
	actual := tr.Units()
	if actual != Rankine {
		t.Errorf("expected %v, got %v", Rankine, actual)
	}
}

func TestTempR_Valid(t *testing.T) {
	tr := TempR{}
	if tr.Valid() {
		t.Errorf("expected false, got true")
	}

	if err := tr.Set(558.87); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !tr.Valid() {
		t.Errorf("expected true, got false")
	}
}

func TestTempR_Temp(t *testing.T) {
	var _ Temp = &TempR{}
}

func BenchmarkTempR_C(b *testing.B) {
	tr := TempR{}
	if err := tr.Set(558.87); err != nil {
		b.Errorf("unexpected error: %v", err)
	}

	for i := 0; i < b.N; i++ {
		tr.C()
	}
}

func BenchmarkTempR_F(b *testing.B) {
	tr := TempR{}
	if err := tr.Set(558.27); err != nil {
		b.Errorf("unexpected error: %v", err)
	}

	for i := 0; i < b.N; i++ {
		tr.F()
	}
}

func BenchmarkTempR_K(b *testing.B) {
	tr := TempR{}
	if err := tr.Set(558.87); err != nil {
		b.Errorf("unexpected error: %v", err)
	}

	for i := 0; i < b.N; i++ {
		tr.K()
	}
}

func BenchmarkTempR_R(b *testing.B) {
	tr := TempR{}
	if err := tr.Set(558.87); err != nil {
		b.Errorf("unexpected error: %v", err)
	}

	for i := 0; i < b.N; i++ {
		tr.R()
	}
}

func BenchmarkTempR_Set(b *testing.B) {
	tr := TempR{}
	for i := 0; i < b.N; i++ {
		if err := tr.Set(558.87); err != nil {
			b.Errorf("unexpected error: %v", err)
		}
	}
}

func BenchmarkTempR_ToC(b *testing.B) {
	tr := TempR{}
	if err := tr.Set(558.87); err != nil {
		b.Errorf("unexpected error: %v", err)
	}

	for i := 0; i < b.N; i++ {
		tr.ToC()
	}
}

func BenchmarkTempR_ToF(b *testing.B) {
	tr := TempR{}
	if err := tr.Set(558.27); err != nil {
		b.Errorf("unexpected error: %v", err)
	}

	for i := 0; i < b.N; i++ {
		tr.ToF()
	}
}

func BenchmarkTempR_ToK(b *testing.B) {
	tr := TempR{}
	if err := tr.Set(558.87); err != nil {
		b.Errorf("unexpected error: %v", err)
	}

	for i := 0; i < b.N; i++ {
		tr.ToK()
	}
}

func BenchmarkTempR_ToR(b *testing.B) {
	tr := TempR{}
	if err := tr.Set(558.87); err != nil {
		b.Errorf("unexpected error: %v", err)
	}

	for i := 0; i < b.N; i++ {
		tr.ToR()
	}
}

func BenchmarkTempR_Units(b *testing.B) {
	tr := TempR{}
	for i := 0; i < b.N; i++ {
		tr.Units()
	}
}
