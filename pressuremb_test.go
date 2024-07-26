package wx

import (
	"github.com/go-wx/wx/internal/tests"
	"testing"
)

func TestPressureMb_HPa(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        float64
		valid       bool
	}{
		{
			name:        "valid",
			measurement: 1013.25,
			want:        1013.25,
			valid:       true,
		},
		{
			name:        "invalid",
			measurement: -1,
			want:        0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressureMb{}
			if err := p.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			if !tests.CloseEnough(p.HPa(), tc.want, 1e-6) {
				t.Errorf("expected %v, got %v", tc.want, p.HPa())
			}
		})
	}
}

func TestPressureMb_InHg(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        float64
		valid       bool
	}{
		{
			name:        "valid",
			measurement: 1013.25,
			want:        1013.25 / conversionFactorInHgToHPa,
			valid:       true,
		},
		{
			name:        "invalid",
			measurement: -1,
			want:        0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressureMb{}
			if err := p.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			if !tests.CloseEnough(p.InHg(), tc.want, 1e-5) {
				t.Errorf("expected %v, got %v", tc.want, p.InHg())
			}
		})
	}
}

func TestPressureMb_KPa(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        float64
		valid       bool
	}{
		{
			name:        "valid",
			measurement: 1013.25,
			want:        1013.25 / 10,
			valid:       true,
		},
		{
			name:        "invalid",
			measurement: -1,
			want:        0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressureMb{}
			if err := p.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			if !tests.CloseEnough(p.KPa(), tc.want, 1e-6) {
				t.Errorf("expected %v, got %v", tc.want, p.KPa())
			}
		})
	}
}

func TestPressureMb_Mb(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        float64
		valid       bool
	}{
		{
			name:        "valid",
			measurement: 1013.25,
			want:        1013.25,
			valid:       true,
		},
		{
			name:        "invalid",
			measurement: -1,
			want:        0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressureMb{}
			if err := p.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			if !tests.CloseEnough(p.Mb(), tc.want, 1e-6) {
				t.Errorf("expected %v, got %v", tc.want, p.Mb())
			}
		})
	}
}

func TestPressureMb_Pa(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        float64
		valid       bool
	}{
		{
			name:        "valid",
			measurement: 1013.25,
			want:        1013.25 * 100,
			valid:       true,
		},
		{
			name:        "invalid",
			measurement: -1,
			want:        0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressureMb{}
			if err := p.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			if !tests.CloseEnough(p.Pa(), tc.want, 1e-6) {
				t.Errorf("expected %v, got %v", tc.want, p.Pa())
			}
		})
	}
}

func TestPressureMb_Psi(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        float64
		valid       bool
	}{
		{
			name:        "valid",
			measurement: 1013.25,
			want:        1013.25 * 100 / conversionFactorPsiToPa,
			valid:       true,
		},
		{
			name:        "invalid",
			measurement: -1,
			want:        0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressureMb{}
			if err := p.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			if !tests.CloseEnough(p.Psi(), tc.want, 1e-6) {
				t.Errorf("expected %v, got %v", tc.want, p.Psi())
			}
		})
	}
}

func TestPressureMb_Set(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		valid       bool
	}{
		{
			name:        "valid",
			measurement: 1013.25,
			valid:       true,
		},
		{
			name:        "invalid",
			measurement: -1,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressureMb{}
			if err := p.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			if p.valid != tc.valid {
				t.Errorf("expected %v, got %v", tc.valid, p.valid)
			}
		})
	}
}

func TestPressureMb_ToHPa(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid",
			measurement: 1013.25,
			expected:    1013.25,
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
			p := PressureMb{}
			if err := p.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			hpa := p.ToHPa()
			if hpa.valid != tc.valid {
				t.Errorf("expected %v, got %v", tc.valid, hpa.valid)
			}

			if !tests.CloseEnough(hpa.HPa(), tc.expected, 1e-6) {
				t.Errorf("expected %v, got %v", tc.measurement, hpa.HPa())
			}
		})
	}
}

func TestPressureMb_ToInHg(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid",
			measurement: 1013.25,
			expected:    1013.25 / conversionFactorInHgToHPa,
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
			p := PressureMb{}
			if err := p.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			inhg := p.ToInHg()
			if inhg.valid != tc.valid {
				t.Errorf("expected %v, got %v", tc.valid, inhg.valid)
			}

			if !tests.CloseEnough(inhg.InHg(), tc.expected, 1e-5) {
				t.Errorf("expected %v, got %v", tc.measurement, inhg.InHg())
			}
		})
	}
}

func TestPressureMb_ToKPa(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid",
			measurement: 1013.25,
			expected:    1013.25 / 10,
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
			p := PressureMb{}
			if err := p.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			kpa := p.ToKPa()
			if kpa.valid != tc.valid {
				t.Errorf("expected %v, got %v", tc.valid, kpa.valid)
			}

			if !tests.CloseEnough(kpa.KPa(), tc.expected, 1e-6) {
				t.Errorf("expected %v, got %v", tc.measurement, kpa.KPa())
			}
		})
	}
}

func TestPressureMb_ToMb(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid",
			measurement: 1013.25,
			expected:    1013.25,
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
			p := PressureMb{}
			if err := p.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			mb := p.ToMb()
			if mb.valid != tc.valid {
				t.Errorf("expected %v, got %v", tc.valid, mb.valid)
			}

			if !tests.CloseEnough(mb.Mb(), tc.expected, 1e-6) {
				t.Errorf("expected %v, got %v", tc.measurement, mb.Mb())
			}
		})
	}
}

func TestPressureMb_ToPa(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid",
			measurement: 1013.25,
			expected:    1013.25 * 100,
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
			p := PressureMb{}
			if err := p.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			pa := p.ToPa()
			if pa.valid != tc.valid {
				t.Errorf("expected %v, got %v", tc.valid, pa.valid)
			}

			if !tests.CloseEnough(pa.Pa(), tc.expected, 1e-6) {
				t.Errorf("expected %v, got %v", tc.measurement, pa.Pa())
			}
		})
	}
}

func TestPressureMb_ToPsi(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		expected    float64
		valid       bool
	}{
		{
			name:        "valid",
			measurement: 1013.25,
			expected:    1013.25 * 100 / conversionFactorPsiToPa,
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
			p := PressureMb{}
			if err := p.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			psi := p.ToPsi()
			if psi.valid != tc.valid {
				t.Errorf("expected %v, got %v", tc.valid, psi.valid)
			}

			if !tests.CloseEnough(psi.Psi(), tc.expected, 1e-6) {
				t.Errorf("expected %v, got %v", tc.measurement, psi.Psi())
			}
		})
	}
}

func TestPressureMb_Units(t *testing.T) {
	p := PressureMb{}
	units := p.Units()
	if units != Mb {
		t.Errorf("expected mb, got %s", units)
	}
}

func TestPressureMb_Valid(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		valid       bool
	}{
		{
			name:        "valid",
			measurement: 1013.25,
			valid:       true,
		},
		{
			name:        "invalid",
			measurement: -1,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressureMb{}
			if err := p.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			if p.Valid() != tc.valid {
				t.Errorf("expected %v, got %v", tc.valid, p.Valid())
			}
		})
	}
}
