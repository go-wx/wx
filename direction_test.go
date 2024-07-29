package wx

import (
	"github.com/go-wx/wx/internal/tests"
	"math"
	"math/rand"
	"testing"
)

func TestNewWindDirection(t *testing.T) {
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
			wd := NewWindDirection(tc.input)
			if !tests.CloseEnough(wd.Degrees().Degrees(), tc.expected, 1e-6) {
				t.Errorf("expected %v; got %v", tc.expected, wd.Degrees())
			}
		})
	}

}

func TestWindDirection_From(t *testing.T) {
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
			wd := NewWindDirection(tc.input)
			if !tests.CloseEnough(wd.From().Degrees(), tc.expected, 1e-6) {
				t.Errorf("expected %v; got %v", tc.expected, wd.From())
			}
		})
	}
}

func TestWindDirection_To(t *testing.T) {
	tt := []struct {
		name     string
		input    float64
		expected float64
	}{
		{
			name:     "0 degrees",
			input:    0,
			expected: 180,
		},
		{
			name:     "360 degrees",
			input:    360,
			expected: 180,
		},
		{
			name:     "720 degrees",
			input:    720,
			expected: 180,
		},
		{
			name:     "90 degrees",
			input:    90,
			expected: 270,
		},
		{
			name:     "450 degrees",
			input:    450,
			expected: 270,
		},
		{
			name:     "-180 degrees",
			input:    -180,
			expected: 0,
		},
		{
			name:     "-450 degrees",
			input:    -450,
			expected: 90,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			wd := NewWindDirection(tc.input)
			if !tests.CloseEnough(wd.To().Degrees(), tc.expected, 1e-6) {
				t.Errorf("expected %v; got %v", tc.expected, wd.To())
			}
		})
	}
}

func TestWindDirection_Degrees(t *testing.T) {
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
			wd := NewWindDirection(tc.input)
			if !tests.CloseEnough(wd.Degrees().Degrees(), tc.expected, 1e-6) {
				t.Errorf("expected %v; got %v", tc.expected, wd.Degrees())
			}
		})
	}
}

func TestWindDirection_Radians(t *testing.T) {
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
		{
			name:     "180 degrees",
			input:    180,
			expected: math.Pi,
		},
		{
			name:     "270 degrees",
			input:    270,
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

func TestWindDirectionUnitVector(t *testing.T) {
	tt := []struct {
		name      string
		degrees   float64
		expectedX float64
		expectedY float64
	}{
		{
			name:      "0 degrees",
			degrees:   0,
			expectedX: 0,
			expectedY: -1,
		},
		{
			name:      "90 degrees",
			degrees:   90,
			expectedX: -1,
			expectedY: 0,
		},
		{
			name:      "180 degrees",
			degrees:   180,
			expectedX: 0,
			expectedY: 1,
		},
		{
			name:      "270 degrees",
			degrees:   270,
			expectedX: 1,
			expectedY: 0,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			wd := NewWindDirection(tc.degrees)
			x, y := wd.UnitVector()
			if !tests.CloseEnough(x, tc.expectedX, 1e-6) {
				t.Errorf("expected x to be %v; got %v", tc.expectedX, x)
			}
			if !tests.CloseEnough(y, tc.expectedY, 1e-6) {
				t.Errorf("expected y to be %v; got %v", tc.expectedY, y)
			}
		})
	}
}

func BenchmarkNewWindDirection(b *testing.B) {
	deg := rand.Float64() * 1080
	for i := 0; i < b.N; i++ {
		NewWindDirection(deg)
	}
}
