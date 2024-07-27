package wx

import "fmt"

// VelocityFps represents a velocity in feet per second.
// This is rarely used in aviation, but is included for completeness.
// A velocity must have a valid value of at least 0.0 fps or greater.
type VelocityFps struct {
	measurement float64
}

// Fps returns the speed in feet per second.
func (v *VelocityFps) Fps() float64 {
	return v.measurement
}

// Kts returns the speed in knots.
func (v *VelocityFps) Kts() float64 {
	return v.measurement / conversionFactorFpsToKnots
}

// Kph returns the speed in kilometers per hour.
func (v *VelocityFps) Kph() float64 {
	return v.measurement * conversionFactorFpsToMps * secondsInHour / 1000
}

// Mph returns the speed in miles per hour.
func (v *VelocityFps) Mph() float64 {
	return v.measurement * secondsInHour / feetPerStatuteMile
}

// Mps returns the speed in meters per second.
func (v *VelocityFps) Mps() float64 {
	return v.measurement * conversionFactorFpsToMps
}

// Set sets the velocity to the specified value.
func (v *VelocityFps) Set(measurement float64) error {
	// Velocity measurement needs to be zero or greater.
	if measurement < 0.0 {
		v.measurement = 0

		return fmt.Errorf(ErrVelocityMustBeZeroOrGreater)
	}

	v.measurement = measurement

	return nil
}

// ToFps converts to feet per second.
func (v *VelocityFps) ToFps() VelocityFps {
	return VelocityFps{
		measurement: v.Fps(),
	}
}

// ToKts converts to knots or nautical miles per hour.
func (v *VelocityFps) ToKts() VelocityKts {
	return VelocityKts{
		measurement: v.Kts(),
	}
}

// ToKph converts to kilometers per hour (kph).
func (v *VelocityFps) ToKph() VelocityKph {
	return VelocityKph{
		measurement: v.Kph(),
	}
}

// ToMph converts to miles per hour.
func (v *VelocityFps) ToMph() VelocityMph {
	return VelocityMph{
		measurement: v.Mph(),
	}
}

// ToMps converts to meters per second.
func (v *VelocityFps) ToMps() VelocityMps {
	return VelocityMps{
		measurement: v.Mps(),
	}
}

// Units return the velocity's units.
func (v *VelocityFps) Units() VelocityUnit {
	return FeetPerSecond
}

// Valid returns true if the magnitude of the velocity
// is zero or greater.
func (v *VelocityFps) Valid() bool {
	return v.measurement >= 0.0
}
