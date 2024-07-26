package wx

// The kelvin, symbol K, is the base unit of measurement for temperature in the International System of Units (SI).
// The Kelvin scale is an absolute temperature scale that starts from 0 K,
// the lowest possible temperature (absolute zero), then rises by exactly 1 K for each 1 °C.
// Source: https://en.wikipedia.org/wiki/Kelvin

// TempK is a temperature measurement in Kelvin.
type TempK struct {
	measurement float64
	valid       bool
}

// C returns the temperature in Celsius.
func (t *TempK) C() float64 {
	return t.measurement - 273.15
}

// F returns the temperature in Fahrenheit.
func (t *TempK) F() float64 {
	return t.C()*TempFactorCtoF + TempFreezingF
}

// K returns the temperature in Kelvin.
func (t *TempK) K() float64 {
	return t.measurement
}

// R returns the temperature in Rankine.
func (t *TempK) R() float64 {
	return t.K() * TempFactorCtoF
}

// Set creates a new temperature measurement in Kelvin.
func (t *TempK) Set(measurement float64) error {
	t.measurement = measurement

	// Check for valid temperature.
	// Source: https://en.wikipedia.org/wiki/Kelvin
	if t.K() < 0 {
		t.measurement = 0
		t.valid = false
		return NewWxErr("temperature must be zero or greater", "TempK.Set")
	}

	t.valid = true
	return nil
}

// ToC converts the temperature to Celsius.
func (t *TempK) ToC() TempC {
	return TempC{
		measurement: t.C(),
		valid:       t.valid,
	}
}

// ToF converts the temperature to Fahrenheit.
func (t *TempK) ToF() TempF {
	return TempF{
		measurement: t.F(),
		valid:       t.valid,
	}
}

// ToK returns the temperature in Kelvin.
func (t *TempK) ToK() TempK {
	return TempK{
		measurement: t.measurement,
		valid:       t.valid,
	}
}

// ToR converts the temperature to Rankine.
func (t *TempK) ToR() TempR {
	return TempR{
		measurement: t.R(),
		valid:       t.valid,
	}
}

// Units returns kelvin.
func (t *TempK) Units() TempUnit {
	return Kelvin
}

// Valid returns true if the temperature is valid.
func (t *TempK) Valid() bool {
	return t.valid
}
