package wx

import "fmt"

// pressureType represents a unit of pressure.
type pressureType uint8

// Pressure units.
// The values start at 1 to avoid a zero value if
// a developer creates an invalid PressureUnit.
//
// For example, if a developer creates a PressureUnit
// with a value of zero, the zero value would be
// considered invalid.
// Example:
//
//	var p PressureUnit
//	if !p.Valid() {
//		// This will be true.
//	}
const (
	hPa  pressureType = iota + 1 // hectopascals
	inHg                         // inches of mercury
	kPa                          // kilopascals
	mb                           // millibars
	pa                           // pascals
	psi                          // pounds per square inch
)

// String returns the string representation of the pressure unit.
func (p pressureType) String() string {
	switch p {
	case hPa:
		return "hPa"
	case inHg:
		return "inHg"
	case kPa:
		return "kPa"
	case mb:
		return "mb"
	case pa:
		return "Pa"
	case psi:
		return "psi"
	}
	return ""
}

// PressureUnit is a type for pressure units.
type PressureUnit struct {
	pressureType
}

// String returns the string representation of the pressure unit.
func (u PressureUnit) String() string {
	return u.pressureType.String()
}

// Pressure units.
var (
	// HPa is hectopascals.
	HPa = PressureUnit{hPa}

	// InHg is inches of mercury.
	InHg = PressureUnit{inHg}

	// KPa is kilopascals.
	KPa = PressureUnit{kPa}

	// Mb is millibars.
	Mb = PressureUnit{mb}

	// Pa is pascals.
	Pa = PressureUnit{pa}

	// Psi is pounds per square inch.
	Psi = PressureUnit{psi}
)

const (
	// convertPsiToPascals is the conversion factor for psi to pascals.
	convertPsiToPascals = 6894.76

	// convertPascalsToPsi is the conversion factor for pascals to psi.
	convertPascalsToPsi = 1 / convertPsiToPascals

	// convertInHgToPascals is the conversion factor for inches of mercury to pascals.
	convertInHgToPascals = 3386.39

	// convertPascalsToInHg is the conversion factor for pascals to inches of mercury.
	convertPascalsToInHg = 1 / convertInHgToPascals
)

// Pressure represents a pressure measurement.
type Pressure struct {
	measurement float64
	unit        PressureUnit
	valid       bool
}

// String returns the string representation of the pressure.
func (p Pressure) String() string {
	if !p.valid {
		return ""
	}

	return fmt.Sprintf("%0.2f %s", p.measurement, p.unit)
}

// Valid returns true if the pressure is valid.
func (p Pressure) Valid() bool {
	return p.valid
}

// NewPressure creates a new Pressure value.
func NewPressure(measurement float64, unit PressureUnit) Pressure {
	switch unit.pressureType {
	case hPa:
		unit = HPa
	case inHg:
		unit = InHg
	case kPa:
		unit = KPa
	case mb:
		unit = Mb
	case pa:
		unit = Pa
	case psi:
		unit = Psi
	default:
		return Pressure{valid: false}
	}

	return Pressure{
		measurement: measurement,
		unit:        unit,
		valid:       measurement >= 0,
	}
}

// HPa returns the pressure in hectopascals.
func (p Pressure) HPa() float64 {
	switch p.unit.pressureType {
	case hPa:
		return p.measurement
	case inHg:
		return p.measurement * convertInHgToPascals / 100.0
	case kPa:
		return p.measurement * 10.0
	case mb:
		return p.measurement
	case pa:
		return p.measurement / 100.0
	case psi:
		return p.measurement * convertPsiToPascals / 100.0
	}

	return 0
}

// InHg returns the pressure in inches of mercury.
func (p Pressure) InHg() float64 {
	switch p.unit.pressureType {
	case hPa:
		return p.measurement * 100 * convertPascalsToInHg
	case inHg:
		return p.measurement
	case kPa:
		return p.measurement * 1000.0 * convertPascalsToInHg
	case mb:
		return p.measurement / convertInHgToPascals * 100.0
	case pa:
		return p.measurement / convertInHgToPascals
	case psi:
		return p.measurement * convertPsiToPascals * convertPascalsToInHg
	}

	return 0
}

// KPa returns the pressure in kilopascals.
func (p Pressure) KPa() float64 {
	switch p.unit.pressureType {
	case hPa:
		return p.measurement / 10.0
	case inHg:
		return p.measurement * convertInHgToPascals / 1000
	case kPa:
		return p.measurement
	case mb:
		return p.measurement / 10.0
	case pa:
		return p.measurement / 1000.0
	case psi:
		return p.measurement * convertPsiToPascals / 1000
	}

	return 0
}

// Mb returns the pressure in millibars.
func (p Pressure) Mb() float64 {
	switch p.unit.pressureType {
	case hPa:
		return p.measurement
	case inHg:
		return p.measurement * convertInHgToPascals / 100.0
	case kPa:
		return p.measurement * 10.0
	case mb:
		return p.measurement
	case pa:
		return p.measurement / 100.0
	case psi:
		return p.measurement * convertPsiToPascals / 100.0
	}

	return 0
}

// Pa returns the pressure in pascals.
func (p Pressure) Pa() float64 {
	switch p.unit.pressureType {
	case hPa:
		return p.measurement * 100.0
	case inHg:
		return p.measurement * convertInHgToPascals
	case kPa:
		return p.measurement * 1000.0
	case mb:
		return p.measurement * 100.0
	case pa:
		return p.measurement
	case psi:
		return p.measurement * convertPsiToPascals
	}

	return 0
}

// Psi returns the pressure in pounds per square inch.
func (p Pressure) Psi() float64 {
	switch p.unit.pressureType {
	case hPa:
		return p.measurement * 100.0 / convertPsiToPascals
	case inHg:
		return p.measurement * convertInHgToPascals / convertPsiToPascals
	case kPa:
		return p.measurement * 1000.0 / convertPsiToPascals
	case mb:
		return p.measurement * 100.0 / convertPsiToPascals
	case pa:
		return p.measurement / convertPsiToPascals
	case psi:
		return p.measurement
	}

	return 0
}

// ToHPa returns the pressure in hectopascals.
func (p Pressure) ToHPa() Pressure {
	return NewPressure(p.HPa(), HPa)
}

// ToInHg returns the pressure in inches of mercury.
func (p Pressure) ToInHg() Pressure {
	return NewPressure(p.InHg(), InHg)
}

// ToKPa returns the pressure in kilopascals.
func (p Pressure) ToKPa() Pressure {
	return NewPressure(p.KPa(), KPa)
}

// ToMb returns the pressure in millibars.
func (p Pressure) ToMb() Pressure {
	return NewPressure(p.Mb(), Mb)
}

// ToPa returns the pressure in pascals.
func (p Pressure) ToPa() Pressure {
	return NewPressure(p.Pa(), Pa)
}

// ToPsi returns the pressure in pounds per square inch.
func (p Pressure) ToPsi() Pressure {
	return NewPressure(p.Psi(), Psi)
}

// Add adds two pressures together and returns a new pressure.
func (p Pressure) Add(p2 Pressure) Pressure {
	return NewPressure(p.HPa()+p2.HPa(), HPa)
}

// Sub subtracts two pressures and returns a new pressure.
func (p Pressure) Sub(p2 Pressure) Pressure {
	return NewPressure(p.HPa()-p2.HPa(), HPa)
}
