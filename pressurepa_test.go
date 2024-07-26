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

func TestPressureKPa_HPa(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid pressure",
			measurement: 101.325,
			expected:    1013.25,
			valid:       true,
		},
		{
			name:        "invalid pressure",
			measurement: -101.325,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressureKPa{}
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

func TestPressureKPa_InHg(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid pressure",
			measurement: 101.325,
			expected:    1013.25 / conversionFactorInHgToHPa,
			valid:       true,
		},
		{
			name:        "invalid pressure",
			measurement: -101.325,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressureKPa{}
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

func TestPressureKPa_KPa(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid pressure",
			measurement: 101.325,
			expected:    101.325,
			valid:       true,
		},
		{
			name:        "invalid pressure",
			measurement: -101.325,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressureKPa{}
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

func TestPressureKPa_Mb(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid pressure",
			measurement: 101.325,
			expected:    1013.25,
			valid:       true,
		},
		{
			name:        "invalid pressure",
			measurement: -101.325,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressureKPa{}
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

func TestPressureKPa_Pa(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid pressure",
			measurement: 101.325,
			expected:    101325,
			valid:       true,
		},
		{
			name:        "invalid pressure",
			measurement: -101.325,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressureKPa{}
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

func TestPressureKPa_Psi(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid pressure",
			measurement: 101.325,
			expected:    101.325 * 1000 / conversionFactorPsiToPa,
			valid:       true,
		},
		{
			name:        "invalid pressure",
			measurement: -101.325,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressureKPa{}
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

func TestPressureKPa_Set(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid pressure",
			measurement: 101.325,
			expected:    101.325,
			valid:       true,
		},
		{
			name:        "invalid pressure",
			measurement: -101.325,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressureKPa{}
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

func TestPressureKPa_ToHPa(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid pressure",
			measurement: 101.325,
			expected:    1013.25,
			valid:       true,
		},
		{
			name:        "invalid pressure",
			measurement: -101.325,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			hpa := PressureKPa{}
			if err := hpa.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			p := hpa.ToHPa()
			if !tests.CloseEnough(p.measurement, tc.expected, 1e-6) {
				t.Errorf("expected %f, got %f", tc.expected, p.measurement)
			}
		})
	}
}

func TestPressureKPa_ToInHg(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid pressure",
			measurement: 101.325,
			expected:    1013.25 / conversionFactorInHgToHPa,
			valid:       true,
		},
		{
			name:        "invalid pressure",
			measurement: -101.325,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			inhg := PressureKPa{}
			if err := inhg.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			p := inhg.ToInHg()
			if !tests.CloseEnough(p.measurement, tc.expected, 1e-6) {
				t.Errorf("expected %f, got %f", tc.expected, p.measurement)
			}
		})
	}
}

func TestPressureKPa_ToKPa(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid pressure",
			measurement: 101.325,
			expected:    101.325,
			valid:       true,
		},
		{
			name:        "invalid pressure",
			measurement: -101.325,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			kpa := PressureKPa{}
			if err := kpa.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			p := kpa.ToKPa()
			if !tests.CloseEnough(p.measurement, tc.expected, 1e-6) {
				t.Errorf("expected %f, got %f", tc.expected, p.measurement)
			}
		})
	}
}

func TestPressureKPa_ToMb(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid pressure",
			measurement: 101.325,
			expected:    1013.25,
			valid:       true,
		},
		{
			name:        "invalid pressure",
			measurement: -101.325,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			mb := PressureKPa{}
			if err := mb.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			p := mb.ToMb()
			if !tests.CloseEnough(p.measurement, tc.expected, 1e-6) {
				t.Errorf("expected %f, got %f", tc.expected, p.measurement)
			}
		})
	}
}

func TestPressureKPa_ToPa(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid pressure",
			measurement: 101.325,
			expected:    101325,
			valid:       true,
		},
		{
			name:        "invalid pressure",
			measurement: -101.325,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			pa := PressureKPa{}
			if err := pa.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			p := pa.ToPa()
			if !tests.CloseEnough(p.measurement, tc.expected, 1e-6) {
				t.Errorf("expected %f, got %f", tc.expected, p.measurement)
			}
		})
	}
}

func TestPressureKPa_ToPsi(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid pressure",
			measurement: 101.325,
			expected:    101.325 * 1000 / conversionFactorPsiToPa,
			valid:       true,
		},
		{
			name:        "invalid pressure",
			measurement: -101.325,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			psi := PressureKPa{}
			if err := psi.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			p := psi.ToPsi()
			if !tests.CloseEnough(p.measurement, tc.expected, 1e-6) {
				t.Errorf("expected %f, got %f", tc.expected, p.measurement)
			}
		})
	}
}

func TestPressureKPa_Units(t *testing.T) {
	p := PressureKPa{}
	if p.Units() != KPa {
		t.Errorf("expected kPa, got %s", p.Units())
	}
}

func TestPressureKPa_Valid(t *testing.T) {
	p := PressureKPa{}
	if p.Valid() {
		t.Errorf("expected false, got true")
	}

	if err := p.Set(101.325); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !p.Valid() {
		t.Errorf("expected true, got false")
	}
}

func TestPressurePa_HPa(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid pressure",
			measurement: 101325,
			expected:    1013.25,
			valid:       true,
		},
		{
			name:        "invalid pressure",
			measurement: -101325,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressurePa{}
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

func TestPressurePa_InHg(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid pressure",
			measurement: 101325,
			expected:    1013.25 / conversionFactorInHgToHPa,
			valid:       true,
		},
		{
			name:        "invalid pressure",
			measurement: -101325,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressurePa{}
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

func TestPressurePa_KPa(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid pressure",
			measurement: 101325,
			expected:    101.325,
			valid:       true,
		},
		{
			name:        "invalid pressure",
			measurement: -101325,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressurePa{}
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

func TestPressurePa_Mb(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid pressure",
			measurement: 101325,
			expected:    1013.25,
			valid:       true,
		},
		{
			name:        "invalid pressure",
			measurement: -101325,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressurePa{}
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

func TestPressurePa_Pa(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid pressure",
			measurement: 101325,
			expected:    101325,
			valid:       true,
		},
		{
			name:        "invalid pressure",
			measurement: -101325,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressurePa{}
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

func TestPressurePa_Psi(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid pressure",
			measurement: 101325,
			expected:    101325 / conversionFactorPsiToPa,
			valid:       true,
		},
		{
			name:        "invalid pressure",
			measurement: -101325,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			p := PressurePa{}
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

func TestPressurePa_Set(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid pressure",
			measurement: 101325,
			expected:    101325,
			valid:       true,
		},
		{
			name:        "invalid pressure",
			measurement: -101325,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			p := PressurePa{}
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

func TestPressurePa_ToHPa(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid pressure",
			measurement: 101325,
			expected:    1013.25,
			valid:       true,
		},
		{
			name:        "invalid pressure",
			measurement: -101325,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			p := PressurePa{}
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

func TestPressurePa_ToInHg(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid pressure",
			measurement: 101325,
			expected:    1013.25 / conversionFactorInHgToHPa,
			valid:       true,
		},
		{
			name:        "invalid pressure",
			measurement: -101325,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			p := PressurePa{}
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

func TestPressurePa_ToKPa(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{

		{
			name:        "valid pressure",
			measurement: 101325,
			expected:    101.325,
			valid:       true,
		},
		{
			name:        "invalid pressure",
			measurement: -101325,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			p := PressurePa{}
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

func TestPressurePa_ToMb(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{

		{
			name:        "valid pressure",
			measurement: 101325,
			expected:    1013.25,
			valid:       true,
		},
		{
			name:        "invalid pressure",
			measurement: -101325,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			p := PressurePa{}
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

func TestPressurePa_ToPa(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{

		{
			name:        "valid pressure",
			measurement: 101325,
			expected:    101325,
			valid:       true,
		},
		{
			name:        "invalid pressure",
			measurement: -101325,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			p := PressurePa{}
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

func TestPressurePa_ToPsi(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{

		{
			name:        "valid pressure",
			measurement: 101325,
			expected:    101325 / conversionFactorPsiToPa,
			valid:       true,
		},
		{
			name:        "invalid pressure",
			measurement: -101325,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			p := PressurePa{}
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

func TestPressurePa_Units(t *testing.T) {
	p := PressurePa{}
	if p.Units() != Pa {
		t.Errorf("expected Pa, got %s", p.Units())
	}
}

func TestPressurePa_Valid(t *testing.T) {
	p := PressurePa{}
	if p.Valid() {
		t.Errorf("expected false, got true")
	}

	if err := p.Set(101325); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !p.Valid() {
		t.Errorf("expected true, got false")
	}
}
