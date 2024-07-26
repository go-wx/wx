package wx

import "fmt"

// PressurePsi is a type for pressure in pounds per square inch.
type PressurePsi struct {
	measurement float64
	valid       bool
}

// HPa returns the pressure in hPa.
func (p *PressurePsi) HPa() float64 {
	return p.measurement * 68.947572932
}

// InHg returns the pressure in inches of mercury.
func (p *PressurePsi) InHg() float64 {
	return p.measurement * 2.036020247
}

// Mb returns the pressure in millibars.
func (p *PressurePsi) Mb() float64 {
	return p.measurement * 68.947572932
}

// KPa returns the pressure in kilopascals.
func (p *PressurePsi) KPa() float64 {
	return p.measurement * 6.8947572932
}

// Psi returns the pressure in pounds per square inch.
func (p *PressurePsi) Psi() float64 {
	return p.measurement
}

// Set sets the pressure.
func (p *PressurePsi) Set(measurement float64) error {
	if measurement < 0 {
		p.valid = false
		return fmt.Errorf(ErrPressureNegative)
	}
	p.measurement = measurement
	p.valid = true

	return nil
}

// ToHPa returns the pressure in hPa.
func (p *PressurePsi) ToHPa() PressureHPa {
	return PressureHPa{
		measurement: p.HPa(),
		valid:       p.valid,
	}
}

// ToInHg returns the pressure in inches of mercury.
func (p *PressurePsi) ToInHg() PressureInHg {
	return PressureInHg{
		measurement: p.InHg(),
		valid:       p.valid,
	}
}

// ToMb returns the pressure in millibars.
func (p *PressurePsi) ToMb() PressureMb {
	return PressureMb{
		measurement: p.Mb(),
		valid:       p.valid,
	}
}

// ToKPa returns the pressure in kilopascals.
func (p *PressurePsi) ToKPa() PressureKPa {
	return PressureKPa{
		measurement: p.KPa(),
		valid:       p.valid,
	}
}

// ToPsi returns the pressure in pounds per square inch.
func (p *PressurePsi) ToPsi() PressurePsi {
	return PressurePsi{
		measurement: p.measurement,
		valid:       p.valid,
	}
}

// Units returns the units of the pressure.
func (p *PressurePsi) Units() PressureUnit {
	return Psi
}

// Valid returns true if the pressure is valid.
func (p *PressurePsi) Valid() bool {
	return p.valid
}
