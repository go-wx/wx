package wx

import "fmt"

const (
	TempZeroCKelvin = 273.15
)

// TempC is a temperature measurement in Celsius.
type TempC struct {
	measurement float64
	valid       bool
}

// C returns the temperature in Celsius.
func (t *TempC) C() float64 {
	return t.measurement
}

// F returns the temperature in Fahrenheit.
func (t *TempC) F() float64 {
	return t.measurement*TempFactorCtoF + TempFreezingF
}

// K returns the temperature in Kelvin.
func (t *TempC) K() float64 {
	return t.measurement + TempZeroCKelvin
}

// R returns the temperature in Rankine.
func (t *TempC) R() float64 {
	return t.F() + 459.67
}

// Set creates a new temperature measurement in Celsius.
func (t *TempC) Set(measurement float64) error {
	t.measurement = measurement
	if t.K() < 0 {
		t.valid = false
		t.measurement = 0
		return NewWxErr(fmt.Sprintf("temperature %v°C is below absolute zero kelvin", t.C()), "TempC.Set")
	}

	t.valid = true

	return nil
}

// Temperature conversion methods.

// ToC returns the temperature in Celsius.
func (t *TempC) ToC() TempC {
	return TempC{
		measurement: t.measurement,
		valid:       t.valid,
	}
}

// ToF converts the temperature to Fahrenheit.
func (t *TempC) ToF() TempF {
	return TempF{
		measurement: t.F(),
		valid:       t.valid,
	}
}

// ToK converts the temperature to Kelvin.
func (t *TempC) ToK() TempK {
	return TempK{
		measurement: t.K(),
		valid:       t.valid,
	}
}

// ToR converts the temperature to Rankine.
func (t *TempC) ToR() TempR {
	return TempR{
		measurement: t.R(),
		valid:       t.valid,
	}
}

// Units returns celsius.
func (t *TempC) Units() TempUnit {
	return Celsius
}

// Valid returns true if the temperature is valid.
func (t *TempC) Valid() bool {
	return t.valid
}
