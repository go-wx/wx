package wx

import (
	"github.com/go-wx/wx/internal/tests"
	"testing"
)

func TestTempK_C(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
	}{
		{name: "absolute zero", measurement: 0, expected: -273.15},
		{name: "freezing", measurement: 273.15, expected: 0},
		{name: "boiling", measurement: 373.15, expected: 100},
		{name: "body", measurement: 310.15, expected: 37},
		{name: "arbitrary", measurement: 100, expected: -173.15},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tk := TempK{}
			if err := tk.Set(tc.measurement); err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			actual := tk.C()
			if !tests.CloseEnough(actual, tc.expected, tests.Tolerance) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestTempK_F(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
	}{
		{name: "absolute zero", measurement: 0, expected: -459.67},
		{name: "freezing", measurement: 273.15, expected: 32},
		{name: "boiling", measurement: 373.15, expected: 212},
		{name: "body", measurement: 310.15, expected: 98.6},
		{name: "arbitrary", measurement: 100, expected: -279.67},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tk := TempK{}
			if err := tk.Set(tc.measurement); err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			actual := tk.F()
			if !tests.CloseEnough(actual, tc.expected, tests.Tolerance) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestTempK_K(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
	}{
		{name: "absolute zero", measurement: 0, expected: 0},
		{name: "freezing", measurement: 273.15, expected: 273.15},
		{name: "boiling", measurement: 373.15, expected: 373.15},
		{name: "body", measurement: 310.15, expected: 310.15},
		{name: "arbitrary", measurement: 100, expected: 100},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tk := TempK{}
			if err := tk.Set(tc.measurement); err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			actual := tk.K()
			if !tests.CloseEnough(actual, tc.expected, tests.Tolerance) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestTempK_R(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
	}{
		{name: "absolute zero", measurement: 0, expected: 0},
		{name: "freezing", measurement: 273.15, expected: 491.67},
		{name: "boiling", measurement: 373.15, expected: 671.67},
		{name: "body", measurement: 310.15, expected: 558.27},
		{name: "arbitrary", measurement: 100, expected: 180.0},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tk := TempK{}
			if err := tk.Set(tc.measurement); err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			actual := tk.R()
			if !tests.CloseEnough(actual, tc.expected, tests.Tolerance) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestTempK_Set(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		error       bool
	}{
		{name: "absolute zero", measurement: 0, expected: 0},
		{name: "freezing", measurement: 273.15, expected: 273.15},
		{name: "boiling", measurement: 373.15, expected: 373.15},
		{name: "body", measurement: 310.15, expected: 310.15},
		{name: "arbitrary", measurement: 100, expected: 100},
		{name: "negative", measurement: -100, expected: -100, error: true},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tk := TempK{}
			err := tk.Set(tc.measurement)

			if tc.error && err == nil {
				t.Fatalf("expected error, got nil")
			}

			if !tc.error && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !tc.error {
				actual := tk.K()
				if !tests.CloseEnough(actual, tc.expected, tests.Tolerance) {
					t.Errorf("expected %v, got %v", tc.expected, actual)
				}
			}
		})
	}
}

func TestTempK_ToC(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
	}{
		{name: "absolute zero", measurement: 0, expected: -273.15},
		{name: "freezing", measurement: 273.15, expected: 0},
		{name: "boiling", measurement: 373.15, expected: 100},
		{name: "body", measurement: 310.15, expected: 37},
		{name: "arbitrary", measurement: 100, expected: -173.15},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tk := TempK{}
			if err := tk.Set(tc.measurement); err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			tcC := tk.ToC()
			actual := tcC.C()
			if !tests.CloseEnough(actual, tc.expected, tests.Tolerance) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestTempK_ToF(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
	}{
		{name: "absolute zero", measurement: 0, expected: -459.67},
		{name: "freezing", measurement: 273.15, expected: 32},
		{name: "boiling", measurement: 373.15, expected: 212},
		{name: "body", measurement: 310.15, expected: 98.6},
		{name: "arbitrary", measurement: 100, expected: -279.67},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tk := TempK{}
			if err := tk.Set(tc.measurement); err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			tcF := tk.ToF()
			actual := tcF.F()
			if !tests.CloseEnough(actual, tc.expected, tests.Tolerance) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestTempK_ToK(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
	}{
		{name: "absolute zero", measurement: 0, expected: 0},
		{name: "freezing", measurement: 273.15, expected: 273.15},
		{name: "boiling", measurement: 373.15, expected: 373.15},
		{name: "body", measurement: 310.15, expected: 310.15},
		{name: "arbitrary", measurement: 100, expected: 100},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tk := TempK{}
			if err := tk.Set(tc.measurement); err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			tcK := tk.ToK()
			actual := tcK.K()
			if !tests.CloseEnough(actual, tc.expected, tests.Tolerance) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestTempK_ToR(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
	}{
		{name: "absolute zero", measurement: 0, expected: 0},
		{name: "freezing", measurement: 273.15, expected: 491.67},
		{name: "boiling", measurement: 373.15, expected: 671.67},
		{name: "body", measurement: 310.15, expected: 558.27},
		{name: "arbitrary", measurement: 100, expected: 180.0},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tk := TempK{}
			if err := tk.Set(tc.measurement); err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			tcR := tk.ToR()
			actual := tcR.R()
			if !tests.CloseEnough(actual, tc.expected, tests.Tolerance) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestTempK_Units(t *testing.T) {
	tk := TempK{}
	expected := Kelvin
	actual := tk.Units()
	if expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestTempK_Valid(t *testing.T) {
	tk := TempK{}
	if tk.Valid() {
		t.Errorf("expected false, got true")
	}

	if err := tk.Set(100); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !tk.Valid() {
		t.Errorf("expected true, got false")
	}

}

func TestTempK_Temp(t *testing.T) {
	var _ Temp = &TempK{}
}

func BenchmarkTempK_C(b *testing.B) {
	tk := TempK{}
	if err := tk.Set(100); err != nil {
		b.Fatalf("unexpected error: %v", err)
	}

	for i := 0; i < b.N; i++ {
		tk.C()
	}
}

func BenchmarkTempK_F(b *testing.B) {
	tk := TempK{}
	if err := tk.Set(100); err != nil {
		b.Fatalf("unexpected error: %v", err)
	}

	for i := 0; i < b.N; i++ {
		tk.F()
	}
}

func BenchmarkTempK_K(b *testing.B) {
	tk := TempK{}
	if err := tk.Set(100); err != nil {
		b.Fatalf("unexpected error: %v", err)
	}

	for i := 0; i < b.N; i++ {
		tk.K()
	}
}

func BenchmarkTempK_R(b *testing.B) {
	tk := TempK{}
	if err := tk.Set(100); err != nil {
		b.Fatalf("unexpected error: %v", err)
	}

	for i := 0; i < b.N; i++ {
		tk.R()
	}
}

func BenchmarkTempK_Set(b *testing.B) {
	tk := TempK{}
	for i := 0; i < b.N; i++ {
		if err := tk.Set(100); err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

func BenchmarkTempK_ToC(b *testing.B) {
	tk := TempK{}
	if err := tk.Set(100); err != nil {
		b.Fatalf("unexpected error: %v", err)
	}

	for i := 0; i < b.N; i++ {
		tk.ToC()
	}
}

func BenchmarkTempK_ToF(b *testing.B) {
	tk := TempK{}
	if err := tk.Set(100); err != nil {
		b.Fatalf("unexpected error: %v", err)
	}

	for i := 0; i < b.N; i++ {
		tk.ToF()
	}
}

func BenchmarkTempK_ToK(b *testing.B) {
	tk := TempK{}
	if err := tk.Set(100); err != nil {
		b.Fatalf("unexpected error: %v", err)
	}

	for i := 0; i < b.N; i++ {
		tk.ToK()
	}
}

func BenchmarkTempK_ToR(b *testing.B) {
	tk := TempK{}
	if err := tk.Set(100); err != nil {
		b.Fatalf("unexpected error: %v", err)
	}

	for i := 0; i < b.N; i++ {
		tk.ToR()
	}
}

func BenchmarkTempK_Units(b *testing.B) {
	tk := TempK{}
	for i := 0; i < b.N; i++ {
		tk.Units()
	}
}
