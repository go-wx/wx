package wx

import (
	"github.com/go-wx/wx/internal/tests"
	"math"
	"math/rand"
	"testing"
)

func TestNewDegrees(t *testing.T) {
	tt := []struct {
		name     string
		input    float64
		expected float64
	}{
		{
			name:     "0 degrees",
			input:    0,
			expected: 0,
		},
		{
			name:     "360 degrees",
			input:    360,
			expected: 0,
		},
		{
			name:     "720 degrees",
			input:    720,
			expected: 0,
		},
		{
			name:     "90 degrees",
			input:    90,
			expected: 90,
		},
		{
			name:     "450 degrees",
			input:    450,
			expected: 90,
		},
		{
			name:     "-180 degrees",
			input:    -180,
			expected: 180,
		},
		{
			name:     "-450 degrees",
			input:    -450,
			expected: 270,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			d := NewDegrees(tc.input)
			if !tests.CloseEnough(d.degrees, tc.expected, 1e-6) {
				t.Errorf("expected %v; got %v", tc.expected, d.degrees)
			}
		})
	}
}

func TestDegrees_Add(t *testing.T) {
	tt := []struct {
		name     string
		degrees  float64
		add      float64
		expected float64
	}{
		{
			name:     "0 + 0",
			degrees:  0,
			add:      0,
			expected: 0,
		},
		{
			name:     "0 + 90",
			degrees:  0,
			add:      90,
			expected: 90,
		},
		{
			name:     "90 + 90",
			degrees:  90,
			add:      90,
			expected: 180,
		},
		{
			name:     "180 + 90",
			degrees:  180,
			add:      90,
			expected: 270,
		},
		{
			name:     "270 + 90",
			degrees:  270,
			add:      90,
			expected: 0,
		},
		{
			name:     "0 + 360",
			degrees:  0,
			add:      360,
			expected: 0,
		},
		{
			name:     "0 + 720",
			degrees:  0,
			add:      720,
			expected: 0,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			d := NewDegrees(tc.degrees)
			d2 := NewDegrees(tc.add)
			d = d.Add(d2)
			if !tests.CloseEnough(d.Degrees(), tc.expected, 1e-6) {
				t.Errorf("expected %v; got %v", tc.expected, d.Degrees())
			}
		})
	}
}

func TestDegrees_Degrees(t *testing.T) {
	tt := []struct {
		name     string
		input    float64
		expected float64
	}{
		{
			name:     "0 degrees",
			input:    0,
			expected: 0,
		},
		{
			name:     "360 degrees",
			input:    360,
			expected: 0,
		},
		{
			name:     "720 degrees",
			input:    720,
			expected: 0,
		},
		{
			name:     "90 degrees",
			input:    90,
			expected: 90,
		},
		{
			name:     "450 degrees",
			input:    450,
			expected: 90,
		},
		{
			name:     "-180 degrees",
			input:    -180,
			expected: 180,
		},
		{
			name:     "-450 degrees",
			input:    -450,
			expected: 270,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			d := NewDegrees(tc.input)
			if !tests.CloseEnough(d.Degrees(), tc.expected, 1e-6) {
				t.Errorf("expected %v; got %v", tc.expected, d.Degrees())
			}
		})
	}
}

func TestDegrees_Divide(t *testing.T) {
	tt := []struct {
		name     string
		degrees  float64
		scalar   float64
		expected float64
	}{
		{
			name:     "0 / 1",
			degrees:  0,
			scalar:   1,
			expected: 0,
		},
		{
			name:     "0 / 2",
			degrees:  0,
			scalar:   2,
			expected: 0,
		},
		{
			name:     "90 / 2",
			degrees:  90,
			scalar:   2,
			expected: 45,
		},
		{
			name:     "180 / 2",
			degrees:  180,
			scalar:   2,
			expected: 90,
		},
		{
			name:     "270 / 2",
			degrees:  270,
			scalar:   2,
			expected: 135,
		},
		{
			name:     "0 / 0",
			degrees:  0,
			scalar:   0,
			expected: 0,
		},
		{
			name:     "0 / -1",
			degrees:  0,
			scalar:   -1,
			expected: 0,
		},
		{
			name:     "0 / -2",
			degrees:  0,
			scalar:   -2,
			expected: 0,
		},
		{
			name:     "90 / -2",
			degrees:  90,
			scalar:   -2,
			expected: 315,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			d := NewDegrees(tc.degrees)
			d = d.Divide(tc.scalar)
			if !tests.CloseEnough(d.Degrees(), tc.expected, 1e-6) {
				t.Errorf("expected %v; got %v", tc.expected, d.Degrees())
			}
		})
	}
}

func TestDegrees_Multiply(t *testing.T) {
	tt := []struct {
		name     string
		degrees  float64
		scalar   float64
		expected float64
	}{
		{
			name:     "0 * 1",
			degrees:  0,
			scalar:   1,
			expected: 0,
		},
		{
			name:     "0 * 2",
			degrees:  0,
			scalar:   2,
			expected: 0,
		},
		{
			name:     "90 * 2",
			degrees:  90,
			scalar:   2,
			expected: 180,
		},
		{
			name:     "180 * 2",
			degrees:  180,
			scalar:   2,
			expected: 0,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			d := NewDegrees(tc.degrees)
			d = d.Multiply(tc.scalar)
			if !tests.CloseEnough(d.Degrees(), tc.expected, 1e-6) {
				t.Errorf("expected %v; got %v", tc.expected, d.Degrees())
			}
		})
	}
}

func TestDegrees_Sub(t *testing.T) {
	tt := []struct {
		name     string
		degrees  float64
		sub      float64
		expected float64
	}{
		{
			name:     "0 - 0",
			degrees:  0,
			sub:      0,
			expected: 0,
		},
		{
			name:     "0 - 90",
			degrees:  0,
			sub:      90,
			expected: 270,
		},
		{
			name:     "90 - 90",
			degrees:  90,
			sub:      90,
			expected: 0,
		},
		{
			name:     "180 - 90",
			degrees:  180,
			sub:      90,
			expected: 90,
		},
		{
			name:     "270 - 90",
			degrees:  270,
			sub:      90,
			expected: 180,
		},
		{
			name:     "720 - 90",
			degrees:  720,
			sub:      90,
			expected: 270,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			d := NewDegrees(tc.degrees)
			d2 := NewDegrees(tc.sub)
			d = d.Sub(d2)
			if !tests.CloseEnough(d.Degrees(), tc.expected, 1e-6) {
				t.Errorf("expected %v; got %v", tc.expected, d.Degrees())
			}
		})
	}

}

func TestDegrees_Radians(t *testing.T) {
	tt := []struct {
		name     string
		input    float64
		expected float64
	}{
		{
			name:     "0 degrees",
			input:    0,
			expected: 0,
		},
		{
			name:     "360 degrees",
			input:    360,
			expected: 0,
		},
		{
			name:     "720 degrees",
			input:    720,
			expected: 0,
		},
		{
			name:     "90 degrees",
			input:    90,
			expected: math.Pi / 2,
		},
		{
			name:     "450 degrees",
			input:    450,
			expected: math.Pi / 2,
		},
		{
			name:     "-180 degrees",
			input:    -180,
			expected: math.Pi,
		},
		{
			name:     "-450 degrees",
			input:    -450,
			expected: 3 * math.Pi / 2,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			deg := NewDegrees(tc.input)
			if !tests.CloseEnough(deg.Radians(), tc.expected, 1e-6) {
				t.Errorf("expected %v; got %v", tc.expected, deg.Radians())
			}
		})
	}
}

func TestDegrees_UnitVector(t *testing.T) {
	tt := []struct {
		name      string
		input     float64
		expectedX float64
		expectedY float64
	}{
		{
			name:      "0 degrees",
			input:     0,
			expectedX: 0,
			expectedY: -1,
		},
		{
			name:      "90 degrees",
			input:     90,
			expectedX: -1,
			expectedY: 0,
		},
		{
			name:      "180 degrees",
			input:     180,
			expectedX: 0,
			expectedY: 1,
		},
		{
			name:      "270 degrees",
			input:     270,
			expectedX: 1,
			expectedY: 0,
		},
		{
			name:      "360 degrees",
			input:     360,
			expectedX: 0,
			expectedY: -1,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			deg := NewDegrees(tc.input)
			x, y := deg.UnitVector()
			if !tests.CloseEnough(x, tc.expectedX, 1e-6) {
				t.Errorf("expected %v; got %v", tc.expectedX, x)
			}
			if !tests.CloseEnough(y, tc.expectedY, 1e-6) {
				t.Errorf("expected %v; got %v", tc.expectedY, y)
			}
		})
	}
}

func TestNaiveDegrees(t *testing.T) {
	tt := []struct {
		name     string
		input    float64
		expected float64
	}{
		{
			name:     "0 degrees",
			input:    0,
			expected: 0,
		},
		{
			name:     "360 degrees",
			input:    360,
			expected: 0,
		},
		{
			name:     "720 degrees",
			input:    720,
			expected: 0,
		},
		{
			name:     "90 degrees",
			input:    90,
			expected: 90,
		},
		{
			name:     "450 degrees",
			input:    450,
			expected: 90,
		},
		{
			name:     "-180 degrees",
			input:    -180,
			expected: 180,
		},
		{
			name:     "-450 degrees",
			input:    -450,
			expected: 270,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			d := naiveDegrees(tc.input)
			if !tests.CloseEnough(d, tc.expected, 1e-6) {
				t.Errorf("expected %v; got %v", tc.expected, d)
			}
		})
	}
}

func BenchmarkNaiveDegrees(b *testing.B) {
	deg := rand.Float64() * 1080
	for i := 0; i < b.N; i++ {
		naiveDegrees(deg)
	}
}

func BenchmarkDegrees_Add(b *testing.B) {
	deg := rand.Float64() * 1080
	deg2 := rand.Float64() * 1080
	d := NewDegrees(deg)
	d2 := NewDegrees(deg2)
	for i := 0; i < b.N; i++ {
		d.Add(d2)
	}
}

func BenchmarkDegrees_Degrees(b *testing.B) {
	deg := rand.Float64() * 1080
	d := NewDegrees(deg)
	for i := 0; i < b.N; i++ {
		d.Degrees()
	}
}

func BenchmarkDegrees_Divide(b *testing.B) {
	deg := rand.Float64() * 1080
	scalar := rand.Float64() * 1080
	d := NewDegrees(deg)
	for i := 0; i < b.N; i++ {
		d.Divide(scalar)
	}
}

func BenchmarkDegrees_Multiply(b *testing.B) {
	deg := rand.Float64() * 1080
	scalar := rand.Float64() * 1080
	d := NewDegrees(deg)
	for i := 0; i < b.N; i++ {
		d.Multiply(scalar)
	}
}

func BenchmarkDegrees_Sub(b *testing.B) {
	deg := rand.Float64() * 1080
	deg2 := rand.Float64() * 1080
	d := NewDegrees(deg)
	d2 := NewDegrees(deg2)
	for i := 0; i < b.N; i++ {
		d.Sub(d2)
	}
}

func BenchmarkDegrees_Radians(b *testing.B) {
	deg := rand.Float64() * 1080
	d := NewDegrees(deg)
	for i := 0; i < b.N; i++ {
		d.Radians()
	}
}

func BenchmarkDegrees_UnitVector(b *testing.B) {
	deg := rand.Float64() * 1080
	d := NewDegrees(deg)
	for i := 0; i < b.N; i++ {
		d.UnitVector()
	}
}

func BenchmarkNewDegrees(b *testing.B) {
	deg := rand.Float64() * 1080
	for i := 0; i < b.N; i++ {
		NewDegrees(deg)
	}
}
