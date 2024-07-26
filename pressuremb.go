package wx

import "fmt"

type PressureMb struct {
	measurement float64
	valid       bool
}

// HPa returns the pressure in hPa (hectopascal).
func (p *PressureMb) HPa() float64 {
	return p.measurement
}

// InHg returns the pressure in inches of mercury.
func (p *PressureMb) InHg() float64 {
	return p.measurement / conversionFactorInHgToHPa
}

// KPa returns the pressure in kilopascals.
func (p *PressureMb) KPa() float64 {
	return p.measurement / 10
}

// Mb returns the pressure in millibars.
func (p *PressureMb) Mb() float64 {
	return p.measurement
}

// Pa returns the pressure in pascals.
func (p *PressureMb) Pa() float64 {
	return p.measurement * 100
}

// Psi returns the pressure in pounds per square inch.
func (p *PressureMb) Psi() float64 {
	return p.measurement * 100 / conversionFactoryPsiToPa
}

// Set sets the pressure.
func (p *PressureMb) Set(measurement float64) error {
	if measurement < 0 {
		p.valid = false
		return fmt.Errorf(ErrPressureNegative)
	}
	p.measurement = measurement
	p.valid = true

	return nil
}

// ToHPa returns the pressure in hPa.
func (p *PressureMb) ToHPa() PressureHPa {
	return PressureHPa{
		measurement: p.HPa(),
		valid:       p.valid,
	}
}

// ToInHg returns the pressure in inches of mercury.
func (p *PressureMb) ToInHg() PressureInHg {
	return PressureInHg{
		measurement: p.InHg(),
		valid:       p.valid,
	}
}

// ToKPa returns the pressure in kilopascals.
func (p *PressureMb) ToKPa() PressureKPa {
	return PressureKPa{
		measurement: p.KPa(),
		valid:       p.valid,
	}
}

// ToMb returns the pressure in millibars.
func (p *PressureMb) ToMb() PressureMb {
	return PressureMb{
		measurement: p.measurement,
		valid:       p.valid,
	}
}

// ToPa returns the pressure in pascals.
func (p *PressureMb) ToPa() PressurePa {
	return PressurePa{
		measurement: p.Pa(),
		valid:       p.valid,
	}
}

// ToPsi returns the pressure in pounds per square inch.
func (p *PressureMb) ToPsi() PressurePsi {
	return PressurePsi{
		measurement: p.Psi(),
		valid:       p.valid,
	}
}

// Units return the units of the pressure.
func (p *PressureMb) Units() PressureUnit {
	return Mb
}

// Valid returns true if the pressure is valid.
func (p *PressureMb) Valid() bool {
	return p.valid
}
