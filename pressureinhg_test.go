package wx

import (
	"github.com/go-wx/wx/internal/tests"
	"testing"
)

func TestPressureInHg_InHg(t *testing.T) {
	tt := []struct {
		name                string
		pressureMeasurement float64
		expected            float64
		valid               bool
	}{
		{name: "valid pressure", pressureMeasurement: 29.92, expected: 29.92, valid: true},
		{name: "invalid pressure", pressureMeasurement: -29.92, expected: 0, valid: false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressureInHg{}
			if err := p.Set(tc.pressureMeasurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			if p.InHg() != tc.expected {
				t.Errorf("expected %f, got %f", tc.expected, p.InHg())
			}
		})
	}
}

func TestPressureInHg_HPa(t *testing.T) {
	tt := []struct {
		name                string
		pressureMeasurement float64
		expected            float64
		valid               bool
	}{
		{
			name:                "valid pressure",
			pressureMeasurement: 29.92,
			expected:            29.92 * conversionFactorInHgToHPa,
			valid:               true,
		},
		{
			name:                "invalid pressure",
			pressureMeasurement: -29.92,
			expected:            0,
			valid:               false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressureInHg{}
			if err := p.Set(tc.pressureMeasurement); err != nil {
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

func TestPressureInHg_Mb(t *testing.T) {
	tt := []struct {
		name                string
		pressureMeasurement float64
		expected            float64
		valid               bool
	}{
		{
			name:                "valid pressure",
			pressureMeasurement: 29.92,
			expected:            29.92 * conversionFactorInHgToHPa,
			valid:               true,
		},
		{
			name:                "invalid pressure",
			pressureMeasurement: -29.92,
			expected:            0,
			valid:               false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressureInHg{}
			if err := p.Set(tc.pressureMeasurement); err != nil {
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

func TestPressureInHg_KPa(t *testing.T) {
	tt := []struct {
		name                string
		pressureMeasurement float64
		expected            float64
		valid               bool
	}{
		{
			name:                "valid pressure",
			pressureMeasurement: 29.92,
			expected:            29.92 * conversionFactorInHgToHPa / 10,
			valid:               true,
		},
		{
			name:                "invalid pressure",
			pressureMeasurement: -29.92,
			expected:            0,
			valid:               false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressureInHg{}
			if err := p.Set(tc.pressureMeasurement); err != nil {
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

func TestPressureInHg_Psi(t *testing.T) {
	tt := []struct {
		name                string
		pressureMeasurement float64
		expected            float64
		valid               bool
	}{
		{
			name:                "valid pressure",
			pressureMeasurement: 29.92,
			expected:            29.92 * conversionFactorInHgToHPa * 100 / conversionFactorPsiToPa,
			valid:               true,
		},
		{
			name:                "invalid pressure",
			pressureMeasurement: -29.92,
			expected:            0,
			valid:               false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressureInHg{}
			if err := p.Set(tc.pressureMeasurement); err != nil {
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

func TestPressureInHg_Set(t *testing.T) {
	tt := []struct {
		name                string
		pressureMeasurement float64
		expected            bool
	}{
		{name: "valid pressure", pressureMeasurement: 29.92, expected: true},
		{name: "invalid pressure", pressureMeasurement: -29.92, expected: false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressureInHg{}
			err := p.Set(tc.pressureMeasurement)
			if err != nil && tc.expected {
				t.Errorf("unexpected error: %v", err)
			}

			if p.valid != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, p.valid)
			}
		})
	}
}

func TestPressureInHg_ToHPa(t *testing.T) {
	tt := []struct {
		name                string
		pressureMeasurement float64
		expected            float64
		valid               bool
	}{
		{
			name:                "valid pressure",
			pressureMeasurement: 29.92,
			expected:            29.92 * conversionFactorInHgToHPa,
			valid:               true,
		},
		{
			name:                "invalid pressure",
			pressureMeasurement: -29.92,
			expected:            0,
			valid:               false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressureInHg{}
			if err := p.Set(tc.pressureMeasurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			hpa := p.ToHPa()
			if !tests.CloseEnough(hpa.HPa(), tc.expected, 1e-6) {
				t.Errorf("expected %f, got %f", tc.expected, hpa.HPa())
			}
		})
	}
}

func TestPressureInHg_ToInHg(t *testing.T) {
	tt := []struct {
		name                string
		pressureMeasurement float64
		expected            float64
		valid               bool
	}{
		{
			name:                "valid pressure",
			pressureMeasurement: 29.92,
			expected:            29.92,
			valid:               true,
		},
		{
			name:                "invalid pressure",
			pressureMeasurement: -29.92,
			expected:            0,
			valid:               false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressureInHg{}
			if err := p.Set(tc.pressureMeasurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			inhg := p.ToInHg()
			if !tests.CloseEnough(inhg.InHg(), tc.expected, 1e-6) {
				t.Errorf("expected %f, got %f", tc.expected, inhg.InHg())
			}
		})
	}
}
