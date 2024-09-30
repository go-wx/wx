package wx

import (
	"math"
	"testing"
)

func TestPressure_InHg(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		measurement float64
		unit        PressureUnit
		want        float64
	}{
		{name: "valid inHg", measurement: 29.9212, unit: InHg, want: 29.9212},
		{name: "valid hPa", measurement: 1013.25, unit: HPa, want: 29.9212},
		{name: "valid kPa", measurement: 101.325, unit: KPa, want: 29.9212},
		{name: "valid mb", measurement: 1013.25, unit: Mb, want: 29.9212},
		{name: "valid pa", measurement: 101325, unit: Pa, want: 29.9212},
		{name: "valid psi", measurement: 14.696, unit: Psi, want: 29.9213},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewPressure(tt.measurement, tt.unit)
			if got := p.InHg(); got != tt.want {
				// Check for floating point differences.
				if math.Abs(got-tt.want) > 0.0001 {
					t.Errorf("expected %v, got %v", tt.want, got)
				}
			}
		})
	}
}

func TestPressure_HPa(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		measurement float64
		unit        PressureUnit
		want        float64
	}{
		{name: "valid inHg", measurement: 29.9212, unit: InHg, want: 1013.25},
		{name: "valid hPa", measurement: 1013.25, unit: HPa, want: 1013.25},
		{name: "valid kPa", measurement: 101.325, unit: KPa, want: 1013.25},
		{name: "valid mb", measurement: 1013.25, unit: Mb, want: 1013.25},
		{name: "valid pa", measurement: 101325, unit: Pa, want: 1013.25},
		{name: "valid psi", measurement: 14.696, unit: Psi, want: 1013.25},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewPressure(tt.measurement, tt.unit)
			if got := p.HPa(); got != tt.want {
				// Check for floating point differences.
				if math.Abs(got-tt.want) > 0.01 {
					t.Errorf("expected %v, got %v", tt.want, got)
				}
			}
		})
	}
}

func TestPressure_KPa(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		measurement float64
		unit        PressureUnit
		want        float64
	}{
		{name: "valid inHg", measurement: 29.9212, unit: InHg, want: 101.325},
		{name: "valid hPa", measurement: 1013.25, unit: HPa, want: 101.325},
		{name: "valid kPa", measurement: 101.325, unit: KPa, want: 101.325},
		{name: "valid mb", measurement: 1013.25, unit: Mb, want: 101.325},
		{name: "valid pa", measurement: 101325, unit: Pa, want: 101.325},
		{name: "valid psi", measurement: 14.696, unit: Psi, want: 101.325},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewPressure(tt.measurement, tt.unit)
			if got := p.KPa(); got != tt.want {
				// Check for floating point differences.
				if math.Abs(got-tt.want) > 0.01 {
					t.Errorf("expected %v, got %v", tt.want, got)
				}
			}
		})
	}
}

func TestPressure_Mb(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		measurement float64
		unit        PressureUnit
		want        float64
	}{
		{name: "valid inHg", measurement: 29.9212, unit: InHg, want: 1013.25},
		{name: "valid hPa", measurement: 1013.25, unit: HPa, want: 1013.25},
		{name: "valid kPa", measurement: 101.325, unit: KPa, want: 1013.25},
		{name: "valid mb", measurement: 1013.25, unit: Mb, want: 1013.25},
		{name: "valid pa", measurement: 101325, unit: Pa, want: 1013.25},
		{name: "valid psi", measurement: 14.696, unit: Psi, want: 1013.25},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewPressure(tt.measurement, tt.unit)
			if got := p.Mb(); got != tt.want {
				// Check for floating point differences.
				if math.Abs(got-tt.want) > 0.01 {
					t.Errorf("expected %v, got %v", tt.want, got)
				}
			}
		})
	}
}

func TestPressure_Pa(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		measurement float64
		unit        PressureUnit
		want        float64
	}{
		{name: "valid inHg", measurement: 29.9212, unit: InHg, want: 101325},
		{name: "valid hPa", measurement: 1013.25, unit: HPa, want: 101325},
		{name: "valid kPa", measurement: 101.325, unit: KPa, want: 101325},
		{name: "valid mb", measurement: 1013.25, unit: Mb, want: 101325},
		{name: "valid pa", measurement: 101325, unit: Pa, want: 101325},
		{name: "valid psi", measurement: 14.696, unit: Psi, want: 101325},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewPressure(tt.measurement, tt.unit)
			if got := p.Pa(); got != tt.want {
				// Check for floating point differences.
				if math.Abs(got-tt.want) > 1 {
					t.Errorf("expected %v, got %v", tt.want, got)
				}
			}
		})
	}
}

func TestPressure_Psi(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		measurement float64
		unit        PressureUnit
		want        float64
	}{
		{name: "valid inHg", measurement: 29.9212, unit: InHg, want: 14.696},
		{name: "valid hPa", measurement: 1013.25, unit: HPa, want: 14.696},
		{name: "valid kPa", measurement: 101.325, unit: KPa, want: 14.696},
		{name: "valid mb", measurement: 1013.25, unit: Mb, want: 14.696},
		{name: "valid pa", measurement: 101325, unit: Pa, want: 14.696},
		{name: "valid psi", measurement: 14.696, unit: Psi, want: 14.696},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewPressure(tt.measurement, tt.unit)
			if got := p.Psi(); got != tt.want {
				// Check for floating point differences.
				if math.Abs(got-tt.want) > 0.001 {
					t.Errorf("expected %v, got %v", tt.want, got)
				}
			}
		})
	}
}
