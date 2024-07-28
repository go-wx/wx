package wx

import (
	"github.com/go-wx/wx/internal/tests"
	"math"
	"math/rand/v2"
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

}

func TestDegrees_Multiply(t *testing.T) {

}

func TestDegrees_Sub(t *testing.T) {

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
			wd := NewWindDirection(tc.input)
			if !tests.CloseEnough(wd.Radians(), tc.expected, 1e-6) {
				t.Errorf("expected %v; got %v", tc.expected, wd.Radians())
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
			wd := NewWindDirection(tc.input)
			x, y := wd.UnitVector()
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
