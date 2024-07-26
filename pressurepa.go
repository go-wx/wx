package wx

import "fmt"

// PressureHPa is a type for pressure in hectopascals.
type PressureHPa struct {
	measurement float64
	valid       bool
}

// HPa returns the pressure in hectopascals.
func (p *PressureHPa) HPa() float64 {
	return p.measurement
}

// InHg returns the pressure in inches of mercury.
func (p *PressureHPa) InHg() float64 {
	return p.measurement * 0.0295299830714
}

// Mb returns the pressure in millibars.
func (p *PressureHPa) Mb() float64 {
	return p.measurement
}

// KPa returns the pressure in kilopascals.
func (p *PressureHPa) KPa() float64 {
	return p.measurement * 0.1
}

// Psi returns the pressure in pounds per square inch.
func (p *PressureHPa) Psi() float64 {
	return p.measurement * 0.014503773773
}

// Set sets the pressure.
func (p *PressureHPa) Set(measurement float64) error {
	if measurement < 0 {
		p.valid = false
		return fmt.Errorf(ErrPressureNegative)
	}
	p.measurement = measurement
	p.valid = true
	return nil
}

// ToHPa returns the pressure in hectopascals.
func (p *PressureHPa) ToHPa() PressureHPa {
	return PressureHPa{
		measurement: p.measurement,
		valid:       p.valid,
	}
}

// ToInHg returns the pressure in inches of mercury.
func (p *PressureHPa) ToInHg() PressureInHg {
	return PressureInHg{
		measurement: p.InHg(),
		valid:       p.valid,
	}
}

// ToMb returns the pressure in millibars.
func (p *PressureHPa) ToMb() PressureMb {
	return PressureMb{
		measurement: p.measurement,
		valid:       p.valid,
	}
}

// ToKPa returns the pressure in kilopascals.
func (p *PressureHPa) ToKPa() PressureKPa {
	return PressureKPa{
		measurement: p.KPa(),
		valid:       p.valid,
	}
}

// ToPsi returns the pressure in pounds per square inch.
func (p *PressureHPa) ToPsi() PressurePsi {
	return PressurePsi{
		measurement: p.Psi(),
		valid:       p.valid,
	}
}

// Units return the units of the pressure.
func (p *PressureHPa) Units() PressureUnit {
	return HPa
}

// Valid returns true if the pressure is valid.
func (p *PressureHPa) Valid() bool {
	return p.valid
}

// PressureKPa is a type for pressure in kilopascals.
type PressureKPa struct {
	measurement float64
	valid       bool
}

// HPa returns the pressure in hectopascals.
func (p *PressureKPa) HPa() float64 {
	return p.measurement * 10
}

// InHg returns the pressure in inches of mercury.
func (p *PressureKPa) InHg() float64 {
	return p.measurement * 0.295299830714
}

// Mb returns the pressure in millibars.
func (p *PressureKPa) Mb() float64 {
	return p.measurement * 10
}

// KPa returns the pressure in kilopascals.
func (p *PressureKPa) KPa() float64 {
	return p.measurement
}

// Psi returns the pressure in pounds per square inch.
func (p *PressureKPa) Psi() float64 {
	return p.measurement * 0.14503773773
}

// Set sets the pressure.
func (p *PressureKPa) Set(measurement float64) error {
	if measurement < 0 {
		p.valid = false
		return fmt.Errorf(ErrPressureNegative)
	}
	p.measurement = measurement
	p.valid = true

	return nil

}

// ToHPa returns the pressure in hectopascals.
func (p *PressureKPa) ToHPa() PressureHPa {
	return PressureHPa{
		measurement: p.HPa(),
		valid:       p.valid,
	}
}

// ToInHg returns the pressure in inches of mercury.
func (p *PressureKPa) ToInHg() PressureInHg {
	return PressureInHg{
		measurement: p.InHg(),
		valid:       p.valid,
	}
}

// ToMb returns the pressure in millibars.
func (p *PressureKPa) ToMb() PressureMb {
	return PressureMb{
		measurement: p.Mb(),
		valid:       p.valid,
	}
}

// ToKPa returns the pressure in kilopascals.
func (p *PressureKPa) ToKPa() PressureKPa {
	return PressureKPa{
		measurement: p.measurement,
		valid:       p.valid,
	}
}

// ToPsi returns the pressure in pounds per square inch.
func (p *PressureKPa) ToPsi() PressurePsi {
	return PressurePsi{
		measurement: p.Psi(),
		valid:       p.valid,
	}
}

// Units return the units of the pressure.
func (p *PressureKPa) Units() PressureUnit {
	return KPa
}

// Valid returns true if the pressure is valid.
func (p *PressureKPa) Valid() bool {
	return p.valid
}

// PressurePa is a type for pressure in pascals.
type PressurePa struct {
	measurement float64
	valid       bool
}

// HPa returns the pressure in hectopascals.
func (p *PressurePa) HPa() float64 {
	return p.measurement * 0.01
}

// InHg returns the pressure in inches of mercury.
func (p *PressurePa) InHg() float64 {
	return p.measurement * 0.000295299830714
}

// Mb returns the pressure in millibars.
func (p *PressurePa) Mb() float64 {
	return p.measurement * 0.01
}

// KPa returns the pressure in kilopascals.
func (p *PressurePa) KPa() float64 {
	return p.measurement * 0.001
}

// Psi returns the pressure in pounds per square inch.
func (p *PressurePa) Psi() float64 {
	return p.measurement * 0.00014503773773
}

// Set sets the pressure.
func (p *PressurePa) Set(measurement float64) error {
	if measurement < 0 {
		p.valid = false
		return fmt.Errorf(ErrPressureNegative)
	}
	p.measurement = measurement
	p.valid = true

	return nil
}

// ToHPa returns the pressure in hectopascals.
func (p *PressurePa) ToHPa() PressureHPa {
	return PressureHPa{
		measurement: p.HPa(),
		valid:       p.valid,
	}
}

// ToInHg returns the pressure in inches of mercury.
func (p *PressurePa) ToInHg() PressureInHg {
	return PressureInHg{
		measurement: p.InHg(),
		valid:       p.valid,
	}
}

// ToMb returns the pressure in millibars.
func (p *PressurePa) ToMb() PressureMb {
	return PressureMb{
		measurement: p.Mb(),
		valid:       p.valid,
	}
}

// ToKPa returns the pressure in kilopascals.
func (p *PressurePa) ToKPa() PressureKPa {
	return PressureKPa{
		measurement: p.KPa(),
		valid:       p.valid,
	}
}

// ToPsi returns the pressure in pounds per square inch.
func (p *PressurePa) ToPsi() PressurePsi {
	return PressurePsi{
		measurement: p.Psi(),
		valid:       p.valid,
	}
}

// Units return the units of the pressure.
func (p *PressurePa) Units() PressureUnit {
	return HPa
}

// Valid returns true if the pressure is valid.
func (p *PressurePa) Valid() bool {
	return p.valid
}
