package wx

import "fmt"

// PressureInHg is a type for pressure in inches of mercury.
type PressureInHg struct {
	measurement float64
	valid       bool
}

// InHg returns the pressure in inches of mercury.
func (p *PressureInHg) InHg() float64 {
	return p.measurement
}

// HPa returns the pressure in hPa.
func (p *PressureInHg) HPa() float64 {
	return p.measurement * 33.8639
}

// Mb returns the pressure in millibars.
func (p *PressureInHg) Mb() float64 {
	return p.measurement * 33.8639
}

// KPa returns the pressure in kilopascals.
func (p *PressureInHg) KPa() float64 {
	return p.measurement * 3.38639
}

// Psi returns the pressure in pounds per square inch.
func (p *PressureInHg) Psi() float64 {
	return p.measurement * 0.491154
}

// Set sets the pressure.
func (p *PressureInHg) Set(measurement float64) error {
	if measurement < 0 {
		p.valid = false
		return fmt.Errorf(ErrPressureNegative)
	}

	p.measurement = measurement
	p.valid = true
	return nil
}

// ToHPa returns the pressure in hPa.
func (p *PressureInHg) ToHPa() PressureHPa {
	return PressureHPa{
		measurement: p.HPa(),
		valid:       p.valid,
	}
}

// ToInHg returns the pressure in inches of mercury.
func (p *PressureInHg) ToInHg() PressureInHg {
	return *p
}

// ToMb returns the pressure in millibars.
func (p *PressureInHg) ToMb() PressureMb {
	return PressureMb{
		measurement: p.Mb(),
		valid:       p.valid,
	}
}

// ToKPa returns the pressure in kilopascals.
func (p *PressureInHg) ToKPa() PressureKPa {
	return PressureKPa{
		measurement: p.KPa(),
		valid:       p.valid,
	}
}

// ToPsi returns the pressure in pounds per square inch.
func (p *PressureInHg) ToPsi() PressurePsi {
	return PressurePsi{
		measurement: p.Psi(),
		valid:       p.valid,
	}
}

// Units return the units of the pressure.
func (p *PressureInHg) Units() PressureUnit {
	return InHg
}

// Valid returns true if the pressure is valid.
func (p *PressureInHg) Valid() bool {
	return p.valid
}
