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
}

const (
	errTempMustBeZeroOrGreater = "temperature must be zero or greater"
)

// TempUnit is a type for temperature units.
type TempUnit struct {
	string
}

// Temperature units.
var (
	Celsius    = TempUnit{"C"}
	Fahrenheit = TempUnit{"F"}
	Kelvin     = TempUnit{"K"}
	Rankine    = TempUnit{"R"}
)
