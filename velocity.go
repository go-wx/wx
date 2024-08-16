package wx

// VelocityUnit is a unit of velocity (e.g., mph, kts, mps, fps, etc.)
type VelocityUnit struct {
	string
}

// String returns the units of the velocity measurement.
func (v VelocityUnit) String() string {
	return v.string
}

var (
	// FeetPerSecond measures the velocity in feet per second.
	// This is rarely used in aviation, but is included for completeness.
	FeetPerSecond = VelocityUnit{"fps"}

	// Knots is the standard unit of velocity and speed measurement
	// for aviation and marine navigation.
	// One knot = 1 nautical mile per hour = 1.15078 mph = 1.852 kph
	Knots = VelocityUnit{"kts"}

	// KilometersPerHour measures the velocity in kilometers per hour.
	// A kilometer is 1,000 meters.
	KilometersPerHour = VelocityUnit{"kph"}

	// MilesPerHour measures the velocity in miles per hour.
	// A statute mile is 5,280 feet.
	MilesPerHour = VelocityUnit{"mph"}

	// MetersPerSecond measures the velocity in meters per second.
	MetersPerSecond = VelocityUnit{"mps"}
)

const (
	// Conversion factors
	conversionFactorFpsToKnots = secondsInHour / feetPerNauticalMile
	conversionFactorFpsToMps   = 0.3048
	secondsInHour              = 60.0 * 60.0
)

const (
	ErrVelocityMustBeZeroOrGreater = "velocity must be zero or greater"
)

// Velocity is an interface for all velocity or speed measurements.
type Velocity interface {
	// Fps returns the speed in feet per second
	Fps() float64

	// Kts returns the speed in knots
	Kts() float64

	// Kph returns the speed in kilometers per hour
	Kph() float64

	// Mph returns the speed in miles per hour
	Mph() float64

	// Mps returns the speed in meters per second
	Mps() float64

	// Set sets the velocity to the specified value.
	Set(float64) error

	// ToFps converts the speed to feet per second
	ToFps() VelocityFps

	// ToKts converts the speed to knots
	ToKts() VelocityKts

	// ToKph converts the speed to kilometers per hour
	ToKph() VelocityKph

	// ToMph converts the speed to miles per hour
	ToMph() VelocityMph

	// ToMps converts the speed to meters per second
	ToMps() VelocityMps

	// Units return the units of the speed
	Units() VelocityUnit

	// Valid returns true if the velocity is zero or
	// greater.
	Valid() bool
}
