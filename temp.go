package wx

import "fmt"

// TempUnit represents a unit of temperature.
type tempType int

// Temperature units.
const (
	celsius tempType = iota
	fahrenheit
	kelvin
	rankine
)

// String returns the string representation of the temperature unit.
func (t tempType) String() string {
	switch t {
	case celsius:
		return "°C"
	case fahrenheit:
		return "°F"
	case kelvin:
		return "K"
	case rankine:
		return "R"
	}
	return ""
}

// TempUnit represents a unit of temperature.
type TempUnit struct {
	tempType
}

func (t TempUnit) String() string {
	return t.tempType.String()
}

// Temperature units.
var (
	Celsius    = TempUnit{celsius}
	Fahrenheit = TempUnit{fahrenheit}
	Kelvin     = TempUnit{kelvin}
	Rankine    = TempUnit{rankine}
)

const (
	// Absolute zero in Celsius.
	absoluteZeroC = -273.15
	// Absolute zero in Fahrenheit.
	absoluteZeroF = -459.67
	// Absolute zero in Kelvin.
	absoluteZeroK = 0
	// Absolute zero in Rankine.
	absoluteZeroR = 0
	// Freezing point of water in Celsius.
	freezingC = 0
	// Freezing point of water in Fahrenheit.
	freezingF = 32
	// Conversion factor from Celsius to Fahrenheit.
	cToF = 9.0 / 5.0
	// Conversion factor from Fahrenheit to Celsius.
	fToC = 5.0 / 9.0
)

// Temp is a temperature measurement.
type Temp struct {
	measurement float64  // All temperatures are stored in the same unit.
	unit        TempUnit // The unit of the temperature.
	valid       bool     // True if the temperature is valid.
}

func (t Temp) String() string {
	switch t.unit.tempType {
	case celsius:
		if t.valid {
			return fmt.Sprintf("%.1f°C", t.measurement)
		}
		return "invalid °C"
	case fahrenheit:
		if t.valid {
			return fmt.Sprintf("%.1f°F", t.measurement)
		}
		return "invalid °F"
	case kelvin:
		if t.valid {
			return fmt.Sprintf("%.1f K", t.measurement)
		}
		return "invalid K"
	case rankine:
		if t.valid {
			return fmt.Sprintf("%.1f R", t.measurement)
		}
		return "invalid R"
	}

	return "invalid temp unit"
}

// Valid returns true if the temperature is valid.
func (t Temp) Valid() bool {
	return t.valid
}

// NewTemp creates a new temperature measurement.
func NewTemp(measurement float64, unit TempUnit) Temp {
	return Temp{measurement,
		unit,
		validMeasurement(measurement, unit)}
}

// C returns the temperature in Celsius.
func (t Temp) C() float64 {
	switch t.unit.tempType {
	case celsius:
		return t.measurement
	case fahrenheit:
		return (t.measurement - freezingF) * fToC
	case kelvin:
		return t.measurement + absoluteZeroC
	case rankine:
		return (t.measurement / cToF) + absoluteZeroC
	}

	return 0
}

// ToC converts the temperature to Celsius.
func (t Temp) ToC() Temp {
	return NewTemp(t.C(), Celsius)
}

// F returns the temperature in Fahrenheit.
func (t Temp) F() float64 {
	switch t.unit.tempType {
	case celsius:
		return t.measurement*cToF + freezingF
	case fahrenheit:
		return t.measurement
	case kelvin:
		return t.measurement*cToF + absoluteZeroF
	case rankine:
		return t.measurement + absoluteZeroF
	}

	return 0
}

// ToF converts the temperature to Fahrenheit.
func (t Temp) ToF() Temp {
	return NewTemp(t.F(), Fahrenheit)
}

// K returns the temperature in Kelvin.
func (t Temp) K() float64 {
	switch t.unit.tempType {
	case celsius:
		return t.measurement - absoluteZeroC
	case fahrenheit:
		return (t.measurement - absoluteZeroF) * fToC
	case kelvin:
		return t.measurement
	case rankine:
		return t.measurement * fToC
	}

	return 0
}

// ToK converts the temperature to Kelvin.
func (t Temp) ToK() Temp {
	return NewTemp(t.K(), Kelvin)
}

// R returns the temperature in Rankine.
func (t Temp) R() float64 {
	switch t.unit.tempType {
	case celsius:
		return (t.measurement - absoluteZeroC) * cToF
	case fahrenheit:
		return t.measurement - absoluteZeroF
	case kelvin:
		return t.measurement * cToF
	case rankine:
		return t.measurement
	}

	return 0
}

// ToR converts the temperature to Rankine.
func (t Temp) ToR() Temp {
	return NewTemp(t.R(), Rankine)
}

// validMeasurement returns true if the measurement is valid
// for the given temperature unit.
func validMeasurement(measurement float64, unit TempUnit) bool {
	switch unit {
	case Celsius:
		return measurement >= absoluteZeroC
	case Fahrenheit:
		return measurement >= absoluteZeroF
	case Kelvin:
		return measurement >= absoluteZeroK
	case Rankine:
		return measurement >= absoluteZeroR
	}
	return false
}
