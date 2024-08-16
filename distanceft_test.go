package wx

import (
	"github.com/go-wx/wx/internal/tests"
	"testing"
)

func TestNewDistanceFT(t *testing.T) {
	t.Run("positive altitude", func(t *testing.T) {
		d := NewDistanceFT(1000)
		if d.feet != 1000 {
			t.Errorf("expected 1000; got %v", d.feet)
		}
		if !d.valid {
			t.Error("expected valid")
		}
	})

	t.Run("negative altitude", func(t *testing.T) {
		d := NewDistanceFT(-1000)
		if d.feet != -1000 {
			t.Errorf("expected -1000; got %v", d.feet)
		}
		if !d.valid {
			t.Error("expected valid")
		}
	})

	t.Run("zero altitude", func(t *testing.T) {
		d := NewDistanceFT(0)
		if d.feet != 0 {
			t.Errorf("expected 0; got %v", d.feet)
		}
		if !d.valid {
			t.Error("expected valid")
		}
	})
}

func TestDistanceFT_FT(t *testing.T) {
	tt := []struct {
		name     string
		input    float64
		expected float64
	}{
		{
			name:     "positive altitude",
			input:    1000,
			expected: 1000,
		},
		{
			name:     "negative altitude",
			input:    -1000,
			expected: -1000,
		},
		{
			name:     "zero altitude",
			input:    0,
			expected: 0,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			d := NewDistanceFT(tc.input)
			if d.FT() != tc.expected {
				t.Errorf("expected %v; got %v", tc.expected, d.FT())
			}
		})
	}
}

func TestDistanceFT_KM(t *testing.T) {
	tt := []struct {
		name     string
		input    float64
		expected float64
	}{
		{
			name:     "positive altitude",
			input:    1000,
			expected: 1000 / feetPerMeter / 1000,
		},
		{
			name:     "negative altitude",
			input:    -1000,
			expected: -1000 / feetPerMeter / 1000,
		},
		{
			name:     "zero altitude",
			input:    0,
			expected: 0,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			d := NewDistanceFT(tc.input)
			if d.KM() != tc.expected {
				t.Errorf("expected %v; got %v", tc.expected, d.KM())
			}
		})
	}
}

func TestDistanceFT_M(t *testing.T) {
	tt := []struct {
		name     string
		input    float64
		expected float64
	}{
		{
			name:     "positive altitude",
			input:    1000,
			expected: 1000 / feetPerMeter,
		},
		{
			name:     "negative altitude",
			input:    -1000,
			expected: -1000 / feetPerMeter,
		},
		{
			name:     "zero altitude",
			input:    0,
			expected: 0,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			d := NewDistanceFT(tc.input)
			if d.M() != tc.expected {
				t.Errorf("expected %v; got %v", tc.expected, d.M())
			}
		})
	}
}

func TestDistanceFT_NM(t *testing.T) {
	tt := []struct {
		name     string
		input    float64
		expected float64
	}{
		{
			name:     "positive altitude",
			input:    1000,
			expected: 1000 / feetPerNauticalMile,
		},
		{
			name:     "negative altitude",
			input:    -1000,
			expected: -1000 / feetPerNauticalMile,
		},
		{
			name:     "zero altitude",
			input:    0,
			expected: 0,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			d := NewDistanceFT(tc.input)
			if !tests.CloseEnough(d.NM(), tc.expected, 1e-6) {
				t.Errorf("expected %v; got %v", tc.expected, d.NM())
			}
		})
	}
}

func TestDistanceFT_SM(t *testing.T) {
	tt := []struct {
		name     string
		input    float64
		expected float64
	}{
		{
			name:     "positive altitude",
			input:    1000,
			expected: 1000 / feetPerStatuteMile,
		},
		{
			name:     "negative altitude",
			input:    -1000,
			expected: -1000 / feetPerStatuteMile,
		},
		{
			name:     "zero altitude",
			input:    0,
			expected: 0,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			d := NewDistanceFT(tc.input)
			if !tests.CloseEnough(d.SM(), tc.expected, 1e-6) {
				t.Errorf("expected %v; got %v", tc.expected, d.SM())
			}
		})
	}
}

func TestDistanceFT_Set(t *testing.T) {
	tt := []struct {
		name     string
		input    float64
		expected float64
		error    bool
	}{
		{
			name:     "positive altitude",
			input:    1000,
			expected: 1000,
			error:    false,
		},
		{
			name:     "negative altitude",
			input:    -1000,
			expected: -1000,
			error:    false,
		},
		{
			name:     "zero altitude",
			input:    0,
			expected: 0,
			error:    false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			d := NewDistanceFT(0)
			err := d.Set(tc.input)
			if err != nil && !tc.error {
				t.Errorf("unexpected error: %v", err)
			}
			if err == nil && tc.error {
				t.Error("expected error")
			}
			if !tests.CloseEnough(d.FT(), tc.expected, 1e-6) {
				t.Errorf("expected %v; got %v", tc.expected, d.FT())
			}
		})
	}
}

func TestDistanceFT_ToFT(t *testing.T) {
	tt := []struct {
		name     string
		input    float64
		expected float64
	}{
		{
			name:     "positive altitude",
			input:    1000,
			expected: 1000,
		},
		{
			name:     "negative altitude",
			input:    -1000,
			expected: -1000,
		},
		{
			name:     "zero altitude",
			input:    0,
			expected: 0,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			d := NewDistanceFT(tc.input)
			d2 := d.ToFT()
			if !tests.CloseEnough(d2.FT(), tc.expected, 1e-6) {
				t.Errorf("expected %v; got %v", tc.expected, d2.FT())
			}
		})
	}
}

func TestDistanceFT_ToKM(t *testing.T) {
	tt := []struct {
		name     string
		input    float64
		expected float64
	}{
		{
			name:     "positive altitude",
			input:    1000,
			expected: 1000 / feetPerMeter / 1000,
		},
		{
			name:     "negative altitude",
			input:    -1000,
			expected: -1000 / feetPerMeter / 1000,
		},
		{
			name:     "zero altitude",
			input:    0,
			expected: 0,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			d := NewDistanceFT(tc.input)
			d2 := d.ToKM()
			if !tests.CloseEnough(d2.KM(), tc.expected, 1e-6) {
				t.Errorf("expected %v; got %v", tc.expected, d2.KM())
			}
		})
	}
}

func TestDistanceFT_ToNM(t *testing.T) {
	tt := []struct {
		name     string
		input    float64
		expected float64
	}{
		{
			name:     "positive altitude",
			input:    1000,
			expected: 1000 / feetPerNauticalMile,
		},
		{
			name:     "negative altitude",
			input:    -1000,
			expected: -1000 / feetPerNauticalMile,
		},
		{
			name:     "zero altitude",
			input:    0,
			expected: 0,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			d := NewDistanceFT(tc.input)
			d2 := d.ToNM()
			if !tests.CloseEnough(d2.NM(), tc.expected, 1e-6) {
				t.Errorf("expected %v; got %v", tc.expected, d2.NM())
			}
		})
	}
}

func TestDistanceFT_ToM(t *testing.T) {
	tt := []struct {
		name     string
		input    float64
		expected float64
	}{
		{
			name:     "positive altitude",
			input:    1000,
			expected: 1000 / feetPerMeter,
		},
		{
			name:     "negative altitude",
			input:    -1000,
			expected: -1000 / feetPerMeter,
		},
		{
			name:     "zero altitude",
			input:    0,
			expected: 0,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			d := NewDistanceFT(tc.input)
			d2 := d.ToM()
			if !tests.CloseEnough(d2.M(), tc.expected, 1e-6) {
				t.Errorf("expected %v; got %v", tc.expected, d2.M())
			}
		})
	}
}

func TestDistanceFT_ToSM(t *testing.T) {
	tt := []struct {
		name     string
		input    float64
		expected float64
	}{
		{
			name:     "positive altitude",
			input:    1000,
			expected: 1000 / feetPerStatuteMile,
		},
		{
			name:     "negative altitude",
			input:    -1000,
			expected: -1000 / feetPerStatuteMile,
		},
		{
			name:     "zero altitude",
			input:    0,
			expected: 0,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			d := NewDistanceFT(tc.input)
			d2 := d.ToSM()
			if !tests.CloseEnough(d2.SM(), tc.expected, 1e-6) {
				t.Errorf("expected %v; got %v", tc.expected, d2.SM())
			}
		})
	}
}

func TestDistanceFT_Units(t *testing.T) {
	d := NewDistanceFT(0)
	if d.Units().string != FT.string {
		t.Errorf("expected ft; got %v", d.Units().string)
	}
}

func TestDistanceFT_Valid(t *testing.T) {
	tt := []struct {
		name     string
		input    float64
		expected bool
	}{
		{
			name:     "positive altitude",
			input:    1000,
			expected: true,
		},
		{
			name:     "negative altitude",
			input:    -1000,
			expected: true,
		},
		{
			name:     "zero altitude",
			input:    0,
			expected: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			d := NewDistanceFT(tc.input)
			if d.Valid() != tc.expected {
				t.Errorf("expected %v; got %v", tc.expected, d.Valid())
			}
		})
	}
}
