package wx

// PressureUnit is a type for pressure units.
type PressureUnit struct {
	string
}

var (
	HPa  = PressureUnit{"hPa"}
	InHg = PressureUnit{"inHg"}
	Mb   = PressureUnit{"mb"}
	KPa  = PressureUnit{"kPa"}
	Psi  = PressureUnit{"psi"}
)

// Pressure is an interface for pressure types.
type Pressure interface {
	// HPa returns the pressure in hPa.
	HPa() float64

	// InHg returns the pressure in inches of mercury.
	InHg() float64

	// Mb returns the pressure in millibars.
	Mb() float64

	// KPa returns the pressure in kilopascals.
	KPa() float64

	// Psi returns the pressure in pounds per square inch.
	Psi() float64

	// Set sets the pressure.
	Set(measurement float64) error

	// ToHPa returns the pressure in hPa.
	ToHPa() float64

	// ToInHg returns the pressure in inches of mercury.
	ToInHg() float64

	// ToMb returns the pressure in millibars.
	ToMb() float64

	// ToKPa returns the pressure in kilopascals.
	ToKPa() float64

	// ToPsi returns the pressure in pounds per square inch.
	ToPsi() float64

	// Units return the units of the pressure.
	Units() string

	// Valid returns true if the pressure is valid.
	Valid() bool
}

// String returns the string representation of the pressure unit.
func (u PressureUnit) String() string {
	return u.string
}
