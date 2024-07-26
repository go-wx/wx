package wx

// PressureUnit is a type for pressure units.
type PressureUnit struct {
	string
}

var (
	HPa  = PressureUnit{"hPa"}
	InHg = PressureUnit{"inHg"}
	KPa  = PressureUnit{"kPa"}
	Mb   = PressureUnit{"mb"}
	Pa   = PressureUnit{"Pa"}
	Psi  = PressureUnit{"psi"}
)

const (
	conversionFactorInHgToHPa = 33.8639
	conversionFactoryPsiToPa  = 6894.76
)

const (
	// ErrPressureNegative is returned when negative pressure is set.
	ErrPressureNegative = "pressure cannot be negative"
)

// Pressure is an interface for pressure types.
type Pressure interface {
	// HPa returns the pressure in hPa.
	HPa() float64

	// InHg returns the pressure in inches of mercury.
	InHg() float64

	// Mb returns the pressure in millibars.
	Mb() float64

	// Pa returns the pressure in pascals.
	Pa() float64

	// KPa returns the pressure in kilopascals.
	KPa() float64

	// Psi returns the pressure in pounds per square inch.
	Psi() float64

	// Set sets the pressure.
	Set(measurement float64) error

	// ToPa returns the pressure in pascals.
	ToPa() PressurePa

	// ToHPa returns the pressure in hPa.
	ToHPa() PressureHPa

	// ToInHg returns the pressure in inches of mercury.
	ToInHg() PressureInHg

	// ToMb returns the pressure in millibars.
	ToMb() PressureMb

	// ToKPa returns the pressure in kilopascals.
	ToKPa() PressureKPa

	// ToPsi returns the pressure in pounds per square inch.
	ToPsi() PressurePsi

	// Units return the units of the pressure.
	Units() PressureUnit

	// Valid returns true if the pressure is valid.
	Valid() bool
}

// String returns the string representation of the pressure unit.
func (u PressureUnit) String() string {
	return u.string
}
