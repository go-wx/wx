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
		return NewWxErr(fmt.Sprintf("temperature %v°F is below absolute zero rankine", t.F()), "TempF.Set")
	}

	return nil
}

// ToC converts the temperature to Celsius.
func (t *TempF) ToC() TempC {
	return TempC{t.C()}
}

// ToF returns the temperature in Fahrenheit.
func (t *TempF) ToF() TempF {
	return TempF{t.measurement}
}

// ToK converts the temperature to Kelvin.
func (t *TempF) ToK() TempK {
	return TempK{t.K()}
}

// ToR converts the temperature to Rankine.
func (t *TempF) ToR() TempR {
	return TempR{t.R()}
}

// Units returns the temperature units.
func (t *TempF) Units() TempUnit {
	return Fahrenheit
}
