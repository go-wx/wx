package wx

const (
	TempZeroCKelvin = 273.15
)

// TempC is a temperature measurement in Celsius.
type TempC struct {
	measurement float64
}

// C returns the temperature in Celsius.
func (t TempC) C() float64 {
	return t.measurement
}

// F returns the temperature in Fahrenheit.
func (t TempC) F() float64 {
	return t.measurement*TempFactorCtoF + TempFreezingF
}

// K returns the temperature in Kelvin.
func (t TempC) K() float64 {
	return t.measurement + TempZeroCKelvin
}

// R returns the temperature in Rankine.
func (t TempC) R() float64 {
	return t.F() + 459.67
}

// Set creates a new temperature measurement in Celsius.
func (t TempC) Set(measurement float64) error {
	if measurement < -TempZeroCKelvin {
		return NewWxErr(errTempMustBeZeroOrGreater, "TempC.New")
	}

	t.measurement = measurement

	return nil
}

// Temperature conversion methods.

// ToC returns the temperature in Celsius.
func (t TempC) ToC() TempC {
	return t
}

// ToF converts the temperature to Fahrenheit.
func (t TempC) ToF() TempF {
	return TempF{t.F()}
}

// ToK converts the temperature to Kelvin.
func (t TempC) ToK() TempK {
	return TempK{t.K()}
}

// ToR converts the temperature to Rankine.
func (t TempC) ToR() TempR {
	return TempR{t.R()}
}

// Units returns celsius.
func (t TempC) Units() TempUnit {
	return Celsius
}
