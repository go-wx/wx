package wx

const (
	TempAbsoluteZeroRF = -459.67
)

// TempR represents a temperature measurement in the Rankine scale.
// TempR implements the Temp interface.
type TempR struct {
	measurement float64
	valid       bool
}

// C returns the temperature in Celsius.
func (t *TempR) C() float64 {
	return (t.measurement - 491.67) * TempFactorFtoC
}

// F returns the temperature in Fahrenheit.
func (t *TempR) F() float64 {
	return t.measurement + TempAbsoluteZeroRF
}

// K returns the temperature in Kelvin.
func (t *TempR) K() float64 {
	return t.measurement * TempFactorFtoC
}

// R returns the temperature in Rankine.
func (t *TempR) R() float64 {
	return t.measurement
}

// Set creates a new temperature measurement in Rankine.
func (t *TempR) Set(measurement float64) error {
	t.measurement = measurement

	if t.R() < 0 {
		t.measurement = 0
		t.valid = false

		return NewWxErr("temperature must be zero or greater", "TempR.Set")
	}

	return nil
}

// ToC converts the temperature to Celsius.
func (t *TempR) ToC() TempC {
	return TempC{
		measurement: t.C(),
		valid:       t.valid,
	}
}

// ToF converts the temperature to Fahrenheit.
func (t *TempR) ToF() TempF {
	return TempF{
		measurement: t.F(),
		valid:       t.valid,
	}
}

// ToK converts the temperature to Kelvin.
func (t *TempR) ToK() TempK {
	return TempK{
		measurement: t.K(),
		valid:       t.valid,
	}
}

// ToR returns the temperature in Rankine.
func (t *TempR) ToR() TempR {
	return TempR{
		measurement: t.R(),
		valid:       t.valid,
	}
}

// Units returns Rankine.
func (t *TempR) Units() TempUnit {
	return Rankine
}

// Valid returns true if the temperature measurement is valid.
func (t *TempR) Valid() bool {
	return t.valid
}
