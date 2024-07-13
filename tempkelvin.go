package wx

import "fmt"

// The kelvin, symbol K, is the base unit of measurement for temperature in the International System of Units (SI).
// The Kelvin scale is an absolute temperature scale that starts from 0 K,
// the lowest possible temperature (absolute zero), then rises by exactly 1 K for each 1 °C.
// Source: https://en.wikipedia.org/wiki/Kelvin

// TempK is a temperature measurement in Kelvin.
type TempK struct {
	measurement float64
}

// C returns the temperature in Celsius.
func (t TempK) C() float64 {
	return t.measurement - 273.15
}

// F returns the temperature in Fahrenheit.
func (t TempK) F() float64 {
	return t.C()*9/5 + 32
}

// K returns the temperature in Kelvin.
func (t TempK) K() float64 {
	return t.measurement
}

// R returns the temperature in Rankine.
func (t TempK) R() float64 {
	return t.F() + 459.67
}

// Set creates a new temperature measurement in Kelvin.
func (t TempK) Set(measurement float64) error {
	if measurement < 0 {
		return NewWxErr(fmt.Sprintf("%s: %0.2f", errTempMustBeZeroOrGreater, measurement), "TempK.New")
	}

	t.measurement = measurement

	return nil
}

// ToC converts the temperature to Celsius.
func (t TempK) ToC() TempC {
	return TempC{t.C()}
}

// ToF converts the temperature to Fahrenheit.
func (t TempK) ToF() TempF {
	return TempF{t.F()}
}

// ToK returns the temperature in Kelvin.
func (t TempK) ToK() TempK {
	return t
}

// ToR converts the temperature to Rankine.
func (t TempK) ToR() TempR {
	return TempR{t.R()}
}

// Units returns kelvin.
func (t TempK) Units() TempUnit {
	return Kelvin
}
