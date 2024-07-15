package wx

// Temp is an interface for temperature measurements.
type Temp interface {
	C() float64
	F() float64
	K() float64
	Set(measurement float64) error
	ToC() TempC
	ToF() TempF
	ToK() TempK
	ToR() TempR
	Units() TempUnit
	Valid() bool
}

// TempUnit is a type for temperature units.
type TempUnit struct {
	string
}

// String returns the string representation of the temperature unit.
func (u TempUnit) String() string {
	return u.string
}

// Temperature units.
var (
	Celsius    = TempUnit{"C"}
	Fahrenheit = TempUnit{"F"}
	Kelvin     = TempUnit{"K"}
	Rankine    = TempUnit{"R"}
)
