package wx

import (
	"github.com/go-wx/wx/internal/tests"
	"testing"
)

func TestTempC_C(t *testing.T) {
	tt := []struct {
		name     string
		value    float64
		expected float64
	}{
		{"0°C", 0, 0},
		{"100°C", 100, 100},
		{"-40°C", -40, -40},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tc := TempC{tc.value}
			if !tests.CloseEnough(tc.C(), tc.measurement, tolerance) {
				t.Fatalf("expected %v, got %v", tc.measurement, tc.C())
			}
		})
	}
}

func TestTempC_F(t *testing.T) {
	tt := []struct {
		name     string
		value    float64
		expected float64
	}{
		{"0°C", 0, 32},
		{"100°C", 100, 212},
		{"-40°C", -40, -40},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c := TempC{}
			if err := c.Set(tc.value); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !tests.CloseEnough(c.F(), tc.expected, tolerance) {
				t.Fatalf("expected %v, got %v", tc.expected, c.F())
			}
		})
	}
}

func TestTempC_K(t *testing.T) {
	tt := []struct {
		name     string
		value    float64
		expected float64
	}{
		{"0°C", 0, 273.15},
		{"100°C", 100, 373.15},
		{"-40°C", -40, 233.15},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c := TempC{}
			if err := c.Set(tc.value); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !tests.CloseEnough(c.K(), tc.expected, tolerance) {
				t.Fatalf("expected %v, got %v", tc.expected, c.K())
			}
		})
	}
}

func TestTempC_R(t *testing.T) {
	tt := []struct {
		name     string
		value    float64
		expected float64
	}{
		{"0°C", 0, 491.67},
		{"100°C", 100, 671.67},
		{"-40°C", -40, 419.67},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c := TempC{}
			if err := c.Set(tc.value); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !tests.CloseEnough(c.R(), tc.expected, tolerance) {
				t.Fatalf("expected %v, got %v", tc.expected, c.R())
			}
		})
	}
}

func TestTempC_Set(t *testing.T) {
	tt := []struct {
		name     string
		value    float64
		expected float64
		err      bool
	}{
		{"0°C", 0, 0, false},
		{"100°C", 100, 100, false},
		{"-40°C", -40, -40, false},
		{"-100°C", -100, -100, false},
		{"0.0°C", 0.0, 0.0, false},
		{"-459.67°C", -459.67, -459.67, true},
		{"-273.16°C", -273.16, -273.16, true},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c := TempC{}
			err := c.Set(tc.value)
			if tc.err && err == nil {
				t.Fatalf("expected error, got nil")
			}

			if !tc.err && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !tests.CloseEnough(c.measurement, tc.expected, tolerance) {
				t.Fatalf("expected %v, got %v", tc.expected, c.measurement)
			}
		})
	}
}

func TestTempC_ToC(t *testing.T) {
	tt := []struct {
		name     string
		value    float64
		expected float64
	}{
		{"0°C", 0, 0},
		{"100°C", 100, 100},
		{"-40°C", -40, -40},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c := TempC{tc.value}
			if !tests.CloseEnough(c.ToC().measurement, tc.expected, tolerance) {
				t.Fatalf("expected %v, got %v", tc.expected, c.ToC().measurement)
			}
		})
	}
}

func TestTempC_ToF(t *testing.T) {
	tt := []struct {
		name     string
		value    float64
		expected float64
	}{
		{"0°C", 0, 32},
		{"100°C", 100, 212},
		{"-40°C", -40, -40},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c := TempC{tc.value}
			if !tests.CloseEnough(c.ToF().measurement, tc.expected, tolerance) {
				t.Fatalf("expected %v, got %v", tc.expected, c.ToF().measurement)
			}
		})
	}
}

func TestTempC_ToK(t *testing.T) {
	tt := []struct {
		name     string
		value    float64
		expected float64
	}{
		{"0°C", 0, 273.15},
		{"100°C", 100, 373.15},
		{"-40°C", -40, 233.15},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c := TempC{tc.value}
			if !tests.CloseEnough(c.ToK().measurement, tc.expected, tolerance) {
				t.Fatalf("expected %v, got %v", tc.expected, c.ToK().measurement)
			}
		})
	}
}

func TestTempC_ToR(t *testing.T) {
	tt := []struct {
		name     string
		value    float64
		expected float64
	}{
		{"0°C", 0, 491.67},
		{"100°C", 100, 671.67},
		{"-40°C", -40, 419.67},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c := TempC{tc.value}
			if !tests.CloseEnough(c.ToR().measurement, tc.expected, tolerance) {
				t.Fatalf("expected %v, got %v", tc.expected, c.ToR().measurement)
			}
		})
	}
}

func TestTempC_Units(t *testing.T) {
	c := TempC{}
	if c.Units() != Celsius {
		t.Fatalf("expected %v, got %v", Celsius, c.Units())
	}
}

func BenchmarkTempC_C(b *testing.B) {
	c := TempC{
		measurement: 100,
	}

	for i := 0; i < b.N; i++ {
		c.C()
	}
}

func BenchmarkTempC_F(b *testing.B) {
	c := TempC{
		measurement: 100,
	}

	for i := 0; i < b.N; i++ {
		c.F()
	}
}

func BenchmarkTempC_K(b *testing.B) {
	c := TempC{
		measurement: 100,
	}

	for i := 0; i < b.N; i++ {
		c.K()
	}
}

func BenchmarkTempC_R(b *testing.B) {
	c := TempC{
		measurement: 100,
	}

	for i := 0; i < b.N; i++ {
		c.R()
	}
}

func BenchmarkTempC_Set(b *testing.B) {
	c := TempC{}
	for i := 0; i < b.N; i++ {
		_ = c.Set(100)
	}
}

func BenchmarkTempC_ToC(b *testing.B) {
	c := TempC{
		measurement: 100,
	}

	for i := 0; i < b.N; i++ {
		c.ToC()
	}
}

func BenchmarkTempC_ToF(b *testing.B) {
	c := TempC{
		measurement: 100,
	}

	for i := 0; i < b.N; i++ {
		c.ToF()
	}
}

func BenchmarkTempC_ToK(b *testing.B) {
	c := TempC{
		measurement: 100,
	}

	for i := 0; i < b.N; i++ {
		c.ToK()
	}
}

func BenchmarkTempC_ToR(b *testing.B) {
	c := TempC{
		measurement: 100,
	}

	for i := 0; i < b.N; i++ {
		c.ToR()
	}
}
