package wx

import (
	"fmt"
	"github.com/go-wx/wx/internal/tests"
	"runtime"
	"testing"
)

const tolerance = 1e-9

func TestTempF_C(t *testing.T) {
	tt := []struct {
		name     string
		value    float64
		expected float64
	}{
		{"0°F", 0, -17.77777777777778},
		{"32°F", 32, 0},
		{"100°F", 100, 37.77777777777778},
		{"212°F", 212, 100},
		{"-40°F", -40, -40},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tf := TempF{}
			if err := tf.Set(tc.value); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if tf.C() != tc.expected {
				t.Fatalf("expected %v, got %v", tc.expected, tf.C())
			}
		})
	}
}

func TestTempF_F(t *testing.T) {
	tt := []struct {
		name     string
		value    float64
		expected float64
	}{
		{"0°F", 0, 0},
		{"32°F", 32, 32},
		{"100°F", 100, 100},
		{"212°F", 212, 212},
		{"-40°F", -40, -40},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tf := TempF{}
			if err := tf.Set(tc.value); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !tests.CloseEnough(tf.F(), tc.expected, tolerance) {
				t.Errorf("expected %v, got %v", tc.expected, tf.F())
			}
		})
	}
}

func TestTempF_K(t *testing.T) {
	tt := []struct {
		name     string
		value    float64
		expected float64
	}{
		{"0°F", 0, 255.3722222222222},
		{"32°F", 32, 273.15},
		{"100°F", 100, 310.9277777777778},
		{"212°F", 212, 373.15},
		{"-40°F", -40, 233.15},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tf := TempF{}
			if err := tf.Set(tc.value); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !tests.CloseEnough(tf.K(), tc.expected, tolerance) {
				t.Fatalf("expected %v, got %v", tc.expected, tf.K())
			}
		})
	}
}

func TestTempF_R(t *testing.T) {
	tt := []struct {
		name     string
		value    float64
		expected float64
	}{
		{"0°F", 0, 459.67},
		{"32°F", 32, 491.67},
		{"100°F", 100, 559.67},
		{"212°F", 212, 671.67},
		{"-40°F", -40, 419.67},
		{"-459.67°F", TempAbsoluteZeroRF, 0},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tf := TempF{}
			if err := tf.Set(tc.value); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !tests.CloseEnough(tf.R(), tc.expected, tolerance) {
				t.Fatalf("expected %v, got %v", tc.expected, tf.R())
			}
		})
	}
}

func TestTempF_Set(t *testing.T) {
	tt := []struct {
		value    float64
		expected float64
		err      bool
	}{
		{0, 0, false},
		{32, 32, false},
		{100, 100, false},
		{-100, -100, false},
		{0.0, 0.0, false},
		{32.0, 32.0, false},
		{100.0, 100.0, false},
		{-500.0, -500.0, true},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%v", tc.value), func(t *testing.T) {
			t.Parallel()
			tf := TempF{}
			err := tf.Set(tc.value)
			if tc.err && err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !tc.err && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if tf.measurement != tc.expected {
				t.Fatalf("expected %v, got %v", tc.expected, tf.measurement)
			}
		})
	}
}

func TestTempF_ToC(t *testing.T) {
	tt := []struct {
		name     string
		value    float64
		expected float64
	}{
		{"0°F", 0, -17.77777777777778},
		{"32°F", 32, 0},
		{"100°F", 100, 37.77777777777778},
		{"212°F", 212, 100},
		{"-40°F", -40, -40},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tf := TempF{}
			if err := tf.Set(tc.value); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			tcC := TempC{tc.expected}
			if !tests.CloseEnough(tf.ToC().measurement, tcC.measurement, tolerance) {
				t.Fatalf("expected %v, got %v", tcC, tf.ToC())
			}
		})
	}
}

func TestTempF_ToF(t *testing.T) {
	tt := []struct {
		name     string
		value    float64
		expected float64
	}{
		{"0°F", 0, 0},
		{"32°F", 32, 32},
		{"100°F", 100, 100},
		{"212°F", 212, 212},
		{"-40°F", -40, -40},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tf := TempF{}
			if err := tf.Set(tc.value); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			tcF := TempF{tc.expected}
			if !tests.CloseEnough(tf.ToF().measurement, tcF.measurement, tolerance) {
				t.Fatalf("expected %v, got %v", tcF, tf.ToF())
			}
		})
	}
}

func TestTempF_ToK(t *testing.T) {
	tt := []struct {
		name     string
		value    float64
		expected float64
	}{
		{"0°F", 0, 255.3722222222222},
		{"32°F", 32, 273.15},
		{"100°F", 100, 310.9277777777778},
		{"212°F", 212, 373.15},
		{"-40°F", -40, 233.15},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tf := TempF{}
			if err := tf.Set(tc.value); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			tcK := TempK{tc.expected}
			if !tests.CloseEnough(tf.ToK().measurement, tcK.measurement, tolerance) {
				t.Fatalf("expected %v, got %v", tcK, tf.ToK())
			}
		})
	}
}

func TestTempF_ToR(t *testing.T) {
	tt := []struct {
		name     string
		value    float64
		expected float64
	}{
		{"0°F", 0, 459.67},
		{"32°F", 32, 491.67},
		{"100°F", 100, 559.67},
		{"212°F", 212, 671.67},
		{"-40°F", -40, 419.67},
		{"-459.67°F", TempAbsoluteZeroRF, 0},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tf := TempF{}
			if err := tf.Set(tc.value); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			tcR := TempR{tc.expected}
			if !tests.CloseEnough(tf.ToR().measurement, tcR.measurement, tolerance) {
				t.Fatalf("expected %v, got %v", tcR, tf.ToR())
			}
		})
	}
}

func TestTempF_Units(t *testing.T) {
	tf := TempF{}
	if tf.Units() != Fahrenheit {
		t.Fatalf("expected %v, got %v", Fahrenheit, tf.Units())
	}
}

func BenchmarkTempF_C(b *testing.B) {
	b.SetParallelism(runtime.NumCPU())
	tf := TempF{}
	err := tf.Set(100)
	if err != nil {
		b.Fatalf("unexpected error: %v", err)
	}

	for i := 0; i < b.N; i++ {
		_ = tf.C()
	}
}

func BenchmarkTempF_F(b *testing.B) {
	b.SetParallelism(runtime.NumCPU())
	tf := TempF{}
	err := tf.Set(100)
	if err != nil {
		b.Fatalf("unexpected error: %v", err)
	}

	for i := 0; i < b.N; i++ {
		_ = tf.F()
	}
}

func BenchmarkTempF_K(b *testing.B) {
	b.SetParallelism(runtime.NumCPU())
	tf := TempF{}
	err := tf.Set(100)
	if err != nil {
		b.Fatalf("unexpected error: %v", err)
	}

	for i := 0; i < b.N; i++ {
		_ = tf.K()
	}
}

func BenchmarkTempF_R(b *testing.B) {
	b.SetParallelism(runtime.NumCPU())
	tf := TempF{}
	err := tf.Set(100)
	if err != nil {
		b.Fatalf("unexpected error: %v", err)
	}

	for i := 0; i < b.N; i++ {
		_ = tf.R()
	}
}

func BenchmarkTempF_Set(b *testing.B) {
	b.SetParallelism(runtime.NumCPU())
	tf := TempF{}
	for i := 0; i < b.N; i++ {
		err := tf.Set(100)
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

func BenchmarkTempF_ToC(b *testing.B) {
	b.SetParallelism(runtime.NumCPU())
	tf := TempF{}
	err := tf.Set(100)
	if err != nil {
		b.Fatalf("unexpected error: %v", err)
	}

	for i := 0; i < b.N; i++ {
		_ = tf.ToC()
	}
}

func BenchmarkTempF_ToF(b *testing.B) {
	b.SetParallelism(runtime.NumCPU())
	tf := TempF{}
	err := tf.Set(100)
	if err != nil {
		b.Fatalf("unexpected error: %v", err)
	}

	for i := 0; i < b.N; i++ {
		_ = tf.ToF()
	}
}

func BenchmarkTempF_ToK(b *testing.B) {
	b.SetParallelism(runtime.NumCPU())
	tf := TempF{}
	err := tf.Set(100)
	if err != nil {
		b.Fatalf("unexpected error: %v", err)
	}

	for i := 0; i < b.N; i++ {
		_ = tf.ToK()
	}
}

func BenchmarkTempF_ToR(b *testing.B) {
	b.SetParallelism(runtime.NumCPU())
	tf := TempF{}
	err := tf.Set(100)
	if err != nil {
		b.Fatalf("unexpected error: %v", err)
	}

	for i := 0; i < b.N; i++ {
		_ = tf.ToR()
	}
}

func BenchmarkTempF_Units(b *testing.B) {
	b.SetParallelism(runtime.NumCPU())
	tf := TempF{}
	for i := 0; i < b.N; i++ {
		_ = tf.Units()
	}
}
