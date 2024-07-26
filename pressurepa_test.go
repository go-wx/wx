package wx

import (
	"github.com/go-wx/wx/internal/tests"
	"testing"
)

func TestPressureHPa_HPa(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid pressure",
			measurement: 1013.25,
			expected:    1013.25,
			valid:       true,
		},
		{
			name:        "invalid pressure",
			measurement: -1013.25,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressureHPa{}
			if err := p.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			if !tests.CloseEnough(p.HPa(), tc.expected, 1e-6) {
				t.Errorf("expected %f, got %f", tc.expected, p.HPa())
			}
		})
	}
}

func TestPressureHPa_InHg(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid pressure",
			measurement: 1013.25,
			expected:    1013.25 / conversionFactorInHgToHPa,
			valid:       true,
		},
		{
			name:        "invalid pressure",
			measurement: -1013.25,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressureHPa{}
			if err := p.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			if !tests.CloseEnough(p.InHg(), tc.expected, 1e-6) {
				t.Errorf("expected %f, got %f", tc.expected, p.InHg())
			}
		})
	}
}

func TestPressureHPa_Mb(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid pressure",
			measurement: 1013.25,
			expected:    1013.25,
			valid:       true,
		},
		{
			name:        "invalid pressure",
			measurement: -1013.25,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressureHPa{}
			if err := p.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			if !tests.CloseEnough(p.Mb(), tc.expected, 1e-6) {
				t.Errorf("expected %f, got %f", tc.expected, p.Mb())
			}
		})
	}
}

func TestPressureHPa_KPa(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid pressure",
			measurement: 1013.25,
			expected:    101.325,
			valid:       true,
		},
		{
			name:        "invalid pressure",
			measurement: -1013.25,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressureHPa{}
			if err := p.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			if !tests.CloseEnough(p.KPa(), tc.expected, 1e-6) {
				t.Errorf("expected %f, got %f", tc.expected, p.KPa())
			}
		})
	}
}

func TestPressureHPa_Pa(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid pressure",
			measurement: 1013.25,
			expected:    101325,
			valid:       true,
		},
		{
			name:        "invalid pressure",
			measurement: -1013.25,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressureHPa{}
			if err := p.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			if !tests.CloseEnough(p.Pa(), tc.expected, 1e-6) {
				t.Errorf("expected %f, got %f", tc.expected, p.Pa())
			}
		})
	}
}

func TestPressureHPa_Psi(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid pressure",
			measurement: 1013.25,
			expected:    1013.25 * 100 / conversionFactorPsiToPa,
			valid:       true,
		},
		{
			name:        "invalid pressure",
			measurement: -1013.25,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressureHPa{}
			if err := p.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			if !tests.CloseEnough(p.Psi(), tc.expected, 1e-6) {
				t.Errorf("expected %f, got %f", tc.expected, p.Psi())
			}
		})
	}
}

func TestPressureHPa_Set(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid pressure",
			measurement: 1013.25,
			expected:    1013.25,
			valid:       true,
		},
		{
			name:        "invalid pressure",
			measurement: -1013.25,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressureHPa{}
			if err := p.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			if p.measurement != tc.expected {
				t.Errorf("expected %f, got %f", tc.expected, p.measurement)
			}
		})
	}
}

func TestPressureHPa_ToHPa(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid pressure",
			measurement: 1013.25,
			expected:    1013.25,
			valid:       true,
		},
		{
			name:        "invalid pressure",
			measurement: -1013.25,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressureHPa{}
			if err := p.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			hpa := p.ToHPa()
			if !tests.CloseEnough(hpa.measurement, tc.expected, 1e-6) {
				t.Errorf("expected %f, got %f", tc.expected, hpa.measurement)
			}
		})
	}
}

func TestPressureHPa_ToInHg(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid pressure",
			measurement: 1013.25,
			expected:    1013.25 / conversionFactorInHgToHPa,
			valid:       true,
		},
		{
			name:        "invalid pressure",
			measurement: -1013.25,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressureHPa{}
			if err := p.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			inhg := p.ToInHg()
			if !tests.CloseEnough(inhg.measurement, tc.expected, 1e-6) {
				t.Errorf("expected %f, got %f", tc.expected, inhg.measurement)
			}
		})
	}
}

func TestPressureHPa_ToKPa(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid pressure",
			measurement: 1013.25,
			expected:    101.325,
			valid:       true,
		},
		{
			name:        "invalid pressure",
			measurement: -1013.25,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressureHPa{}
			if err := p.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			kpa := p.ToKPa()
			if !tests.CloseEnough(kpa.measurement, tc.expected, 1e-6) {
				t.Errorf("expected %f, got %f", tc.expected, kpa.measurement)
			}
		})
	}
}

func TestPressureHPa_ToMb(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid pressure",
			measurement: 1013.25,
			expected:    1013.25,
			valid:       true,
		},
		{
			name:        "invalid pressure",
			measurement: -1013.25,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressureHPa{}
			if err := p.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			mb := p.ToMb()
			if !tests.CloseEnough(mb.measurement, tc.expected, 1e-6) {
				t.Errorf("expected %f, got %f", tc.expected, mb.measurement)
			}
		})
	}
}

func TestPressureHPa_ToPa(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid pressure",
			measurement: 1013.25,
			expected:    101325,
			valid:       true,
		},
		{
			name:        "invalid pressure",
			measurement: -1013.25,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressureHPa{}
			if err := p.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			pa := p.ToPa()
			if !tests.CloseEnough(pa.measurement, tc.expected, 1e-6) {
				t.Errorf("expected %f, got %f", tc.expected, pa.measurement)
			}
		})
	}
}

func TestPressureHPa_ToPsi(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid pressure",
			measurement: 1013.25,
			expected:    1013.25 * 100 / conversionFactorPsiToPa,
			valid:       true,
		},
		{
			name:        "invalid pressure",
			measurement: -1013.25,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressureHPa{}
			if err := p.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			psi := p.ToPsi()
			if !tests.CloseEnough(psi.measurement, tc.expected, 1e-6) {
				t.Errorf("expected %f, got %f", tc.expected, psi.measurement)
			}
		})
	}
}

func TestPressureHPa_Units(t *testing.T) {
	p := PressureHPa{}
	if p.Units() != HPa {
		t.Errorf("expected hPa, got %s", p.Units())
	}
}

func TestPressureHPa_Valid(t *testing.T) {
	p := PressureHPa{}
	if p.Valid() {
		t.Errorf("expected false, got true")
	}

	if err := p.Set(1013.25); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !p.Valid() {
		t.Errorf("expected true, got false")
	}
}
