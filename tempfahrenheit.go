package wx

import "fmt"

// Fahrenheit
// Source: https://en.wikipedia.org/wiki/Fahrenheit

const (
	TempFreezingF  = 32.0
	TempFactorFtoC = 5.0 / 9.0
	TempFactorCtoF = 9.0 / 5.0
)

// TempF is a temperature measurement in Fahrenheit.
type TempF struct {
	measurement float64
	valid       bool
}

// C returns the temperature in Celsius.
func (t *TempF) C() float64 {
	return (t.measurement - TempFreezingF) * TempFactorFtoC
}

// F returns the temperature in Fahrenheit.
func (t *TempF) F() float64 {
	return t.measurement
}

// K returns the temperature in Kelvin.
func (t *TempF) K() float64 {
	return t.C() + TempZeroCKelvin
}

// R returns the temperature in Rankine.
func (t *TempF) R() float64 {
	return t.F() - TempAbsoluteZeroRF
}

// Set creates a new temperature measurement in Fahrenheit.
func (t *TempF) Set(measurement float64) error {
	t.measurement = measurement
	if t.R() < 0 {
		t.valid = false
		return NewWxErr(fmt.Sprintf("temperature %v°F is below absolute zero rankine", t.F()), "TempF.Set")
	}

	return nil
}

// ToC converts the temperature to Celsius.
func (t *TempF) ToC() TempC {
	return TempC{
		measurement: t.C(),
		valid:       t.valid,
	}
}

// ToF returns the temperature in Fahrenheit.
func (t *TempF) ToF() TempF {
	return TempF{
		measurement: t.measurement,
		valid:       t.valid,
	}
}

// ToK converts the temperature to Kelvin.
func (t *TempF) ToK() TempK {
	return TempK{
		measurement: t.K(),
		valid:       t.valid,
	}
}

// ToR converts the temperature to Rankine.
func (t *TempF) ToR() TempR {
	return TempR{
		measurement: t.R(),
		valid:       t.valid,
	}
}

// Units return the temperature units.
func (t *TempF) Units() TempUnit {
	return Fahrenheit
}

// Valid returns true if the temperature is valid.
func (t *TempF) Valid() bool {
	return t.valid
}
