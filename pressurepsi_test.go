package wx

import (
	"github.com/go-wx/wx/internal/tests"
	"testing"
)

func TestPressurePsi_HPa(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid",
			measurement: 1,
			expected:    conversionFactorPsiToPa / 100,
			valid:       true,
		},
		{
			name:        "invalid",
			measurement: -1,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressurePsi{}
			if err := p.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error %v", err)
				}
			}

			if !tests.CloseEnough(p.HPa(), tc.expected, 1e-6) {
				t.Errorf("expected %v, got %v", tc.expected, p.HPa())
			}
		})
	}
}

func TestPressurePsi_InHg(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid",
			measurement: 1,
			expected:    conversionFactorPsiToPa / (conversionFactorInHgToHPa * 100),
			valid:       true,
		},
		{
			name:        "invalid",
			measurement: -1,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressurePsi{}
			if err := p.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error %v", err)
				}
			}

			if !tests.CloseEnough(p.InHg(), tc.expected, 1e-6) {
				t.Errorf("expected %v, got %v", tc.expected, p.InHg())
			}
		})
	}
}

func TestPressurePsi_KPa(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid",
			measurement: 1,
			expected:    conversionFactorPsiToPa / 1000,
			valid:       true,
		},
		{
			name:        "invalid",
			measurement: -1,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressurePsi{}
			if err := p.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error %v", err)
				}
			}

			if !tests.CloseEnough(p.KPa(), tc.expected, 1e-6) {
				t.Errorf("expected %v, got %v", tc.expected, p.KPa())
			}
		})
	}
}

func TestPressurePsi_Mb(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid",
			measurement: 1,
			expected:    conversionFactorPsiToPa / 100,
			valid:       true,
		},
		{
			name:        "invalid",
			measurement: -1,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressurePsi{}
			if err := p.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error %v", err)
				}
			}

			if !tests.CloseEnough(p.Mb(), tc.expected, 1e-6) {
				t.Errorf("expected %v, got %v", tc.expected, p.Mb())
			}
		})
	}
}

func TestPressurePsi_Pa(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid",
			measurement: 1,
			expected:    conversionFactorPsiToPa,
			valid:       true,
		},
		{
			name:        "invalid",
			measurement: -1,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressurePsi{}
			if err := p.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error %v", err)
				}
			}

			if !tests.CloseEnough(p.Pa(), tc.expected, 1e-6) {
				t.Errorf("expected %v, got %v", tc.expected, p.Pa())
			}
		})
	}
}

func TestPressurePsi_Psi(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid",
			measurement: 1,
			expected:    1,
			valid:       true,
		},
		{
			name:        "invalid",
			measurement: -1,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressurePsi{}
			if err := p.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error %v", err)
				}
			}

			if !tests.CloseEnough(p.Psi(), tc.expected, 1e-6) {
				t.Errorf("expected %v, got %v", tc.expected, p.Psi())
			}
		})
	}
}

func TestPressurePsi_Set(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid",
			measurement: 1,
			expected:    1,
			valid:       true,
		},
		{
			name:        "invalid",
			measurement: -1,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressurePsi{}
			if err := p.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error %v", err)
				}
			}

			if p.measurement != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, p.measurement)
			}
		})
	}
}

func TestPressurePsi_ToHPa(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid",
			measurement: 1,
			expected:    conversionFactorPsiToPa / 100,
			valid:       true,
		},
		{
			name:        "invalid",
			measurement: -1,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressurePsi{}
			if err := p.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error %v", err)
				}
			}

			hpa := p.ToHPa()
			if !tests.CloseEnough(hpa.HPa(), tc.expected, 1e-6) {
				t.Errorf("expected %v, got %v", tc.expected, hpa.HPa())
			}
		})
	}
}

func TestPressurePsi_ToInHg(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid",
			measurement: 1,
			expected:    conversionFactorPsiToPa / (conversionFactorInHgToHPa * 100),
			valid:       true,
		},
		{
			name:        "invalid",
			measurement: -1,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressurePsi{}
			if err := p.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error %v", err)
				}
			}

			inhg := p.ToInHg()
			if !tests.CloseEnough(inhg.InHg(), tc.expected, 1e-6) {
				t.Errorf("expected %v, got %v", tc.expected, inhg.InHg())
			}
		})
	}
}

func TestPressurePsi_ToKPa(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid",
			measurement: 1,
			expected:    conversionFactorPsiToPa / 1000,
			valid:       true,
		},
		{
			name:        "invalid",
			measurement: -1,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressurePsi{}
			if err := p.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error %v", err)
				}
			}

			kpa := p.ToKPa()
			if !tests.CloseEnough(kpa.KPa(), tc.expected, 1e-6) {
				t.Errorf("expected %v, got %v", tc.expected, kpa.KPa())
			}
		})
	}
}

func TestPressurePsi_ToMb(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid",
			measurement: 1,
			expected:    conversionFactorPsiToPa / 100,
			valid:       true,
		},
		{
			name:        "invalid",
			measurement: -1,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressurePsi{}
			if err := p.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error %v", err)
				}
			}

			mb := p.ToMb()
			if !tests.CloseEnough(mb.Mb(), tc.expected, 1e-6) {
				t.Errorf("expected %v, got %v", tc.expected, mb.Mb())
			}
		})
	}
}

func TestPressurePsi_ToPa(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid",
			measurement: 1,
			expected:    conversionFactorPsiToPa,
			valid:       true,
		},
		{
			name:        "invalid",
			measurement: -1,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressurePsi{}
			if err := p.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error %v", err)
				}
			}

			pa := p.ToPa()
			if !tests.CloseEnough(pa.Pa(), tc.expected, 1e-6) {
				t.Errorf("expected %v, got %v", tc.expected, pa.Pa())
			}
		})
	}
}

func TestPressurePsi_ToPsi(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid",
			measurement: 1,
			expected:    1,
			valid:       true,
		},
		{
			name:        "invalid",
			measurement: -1,
			expected:    0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressurePsi{}
			if err := p.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error %v", err)
				}
			}

			psi := p.ToPsi()
			if !tests.CloseEnough(psi.Psi(), tc.expected, 1e-6) {
				t.Errorf("expected %v, got %v", tc.expected, psi.Psi())
			}
		})
	}
}

func TestPressurePsi_Units(t *testing.T) {
	p := PressurePsi{}
	units := p.Units()
	if units != Psi {
		t.Errorf("expected psi, got %v", units)
	}
}

func TestPressurePsi_Valid(t *testing.T) {
	p := PressurePsi{}
	if p.Valid() {
		t.Errorf("expected false, got true")
	}

	if err := p.Set(1); err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if !p.Valid() {
		t.Errorf("expected true, got false")
	}
}
