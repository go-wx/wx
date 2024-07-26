package wx

import (
	"github.com/go-wx/wx/internal/tests"
	"testing"
)

func TestPressureHPa_Pa(t *testing.T) {
	tt := []struct {
		name                string
		pressureMeasurement float64
		expected            float64
		valid               bool
	}{
		{
			name:                "valid pressure",
			pressureMeasurement: 1013.25,
			expected:            101325,
			valid:               true,
		},
		{
			name:                "invalid pressure",
			pressureMeasurement: -1013.25,
			expected:            0,
			valid:               false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressureHPa{}
			if err := p.Set(tc.pressureMeasurement); err != nil {
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

func TestPressureHPa_HPa(t *testing.T) {
	tt := []struct {
		name                string
		pressureMeasurement float64
		expected            float64
		valid               bool
	}{
		{
			name:                "valid pressure",
			pressureMeasurement: 1013.25,
			expected:            1013.25,
			valid:               true,
		},
		{
			name:                "invalid pressure",
			pressureMeasurement: -1013.25,
			expected:            0,
			valid:               false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			p := PressureHPa{}
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
