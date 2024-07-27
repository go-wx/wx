package wx

import "fmt"

// VelocityKph is a velocity measurement in kilometers
// per hour (kph).
type VelocityKph struct {
	measurement float64
	valid       bool
}

// Fps returns the velocity in feet per second (fps).
func (v *VelocityKph) Fps() float64 {
	return v.measurement / secondsInHour * 1000 * feetPerMeter
}

// Kts returns the velocity in knots.
func (v *VelocityKph) Kts() float64 {
	return v.Fps() * conversionFactorFpsToKnots
}

// Kph returns the velocity in kilometers per hour.
func (v *VelocityKph) Kph() float64 {
	return v.measurement
}

// Mph returns the velocity in miles per hour.
func (v *VelocityKph) Mph() float64 {
	return v.Fps() * feetPerStatuteMile / secondsInHour
}

// Mps returns the velocity in meters per second.
func (v *VelocityKph) Mps() float64 {
	return v.measurement / secondsInHour * 1000
}

// Set sets the velocity to the specified value.
func (v *VelocityKph) Set(measurement float64) error {
	// Velocity measurement needs to be zero or greater.
	if measurement < 0.0 {
		v.measurement = 0
		v.valid = false

		return fmt.Errorf(ErrVelocityMustBeZeroOrGreater)
	}

	v.measurement = measurement
	v.valid = true

	return nil
}

// ToFps converts the velocity to feet per second.
func (v *VelocityKph) ToFps() VelocityFps {
	return VelocityFps{
		measurement: v.Fps(),
		valid:       v.valid,
	}
}

// ToKts converts the velocity to knots.
func (v *VelocityKph) ToKts() VelocityKts {
	return VelocityKts{
		measurement: v.Kts(),
		valid:       v.valid,
	}
}

// ToKph converts the velocity to kilometers per hour.
func (v *VelocityKph) ToKph() VelocityKph {
	return VelocityKph{
		measurement: v.Kph(),
		valid:       v.valid,
	}
}

// ToMph converts the velocity to miles per hour.
func (v *VelocityKph) ToMph() VelocityMph {
	return VelocityMph{
		measurement: v.Mph(),
		valid:       v.valid,
	}
}

// ToMps converts the velocity to meters per second.
func (v *VelocityKph) ToMps() VelocityMps {
	return VelocityMps{
		measurement: v.Mps(),
		valid:       v.valid,
	}
}

// Units return the units of the velocity.
func (v *VelocityKph) Units() VelocityUnit {
	return KilometersPerHour
}

// Valid returns true if the velocity is zero or greater.
func (v *VelocityKph) Valid() bool {
	return v.valid
}
