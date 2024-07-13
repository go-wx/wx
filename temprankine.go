package wx

const (
	TempAbsoluteZeroRF = -459.67
)

// TempR represents a temperature measurement in the Rankine scale.
// TempR implements the Temp interface.
type TempR struct {
	measurement float64
}

// C returns the temperature in Celsius.
func (t *TempR) C() float64 {
	return (t.measurement - TempAbsoluteZeroRF) * TempFactorFtoC
}

// F returns the temperature in Fahrenheit.
func (t *TempR) F() float64 {
	return t.measurement - TempAbsoluteZeroRF
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
		return NewWxErr("temperature must be zero or greater", "TempR.Set")
	}

	return nil
}

// ToC converts the temperature to Celsius.
func (t *TempR) ToC() TempC {
	return TempC{t.C()}
}

// ToF converts the temperature to Fahrenheit.
func (t *TempR) ToF() TempF {
	return TempF{t.F()}
}

// ToK converts the temperature to Kelvin.
func (t *TempR) ToK() TempK {
	return TempK{t.K()}
}

// ToR returns the temperature in Rankine.
func (t *TempR) ToR() TempR {
	return TempR{t.R()}
}

// Units returns Rankine.
func (t *TempR) Units() TempUnit {
	return Rankine
}
