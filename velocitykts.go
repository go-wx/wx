package wx

import "fmt"

// VelocityKts is a velocity measurement in knots or
// nautical miles per hour.
type VelocityKts struct {
	measurement float64
	valid       bool
}

// Fps returns the speed in feet per second.
func (v *VelocityKts) Fps() float64 {
	return v.measurement / secondsInHour * feetPerNauticalMile
}

// Kts returns the speed in knots.
func (v *VelocityKts) Kts() float64 {
	return v.measurement
}

// Kph returns the speed in kilometers per hour.
func (v *VelocityKts) Kph() float64 {
	return v.Fps() / feetPerMeter / 1000 * secondsInHour
}

// Mph returns the speed in miles per hour.
func (v *VelocityKts) Mph() float64 {
	return v.measurement * (feetPerStatuteMile / feetPerNauticalMile)
}

// Mps returns the speed in meters per second.
func (v *VelocityKts) Mps() float64 {
	return v.Fps() / feetPerMeter
}

// Set sets the velocity to the specified value.
func (v *VelocityKts) Set(measurement float64) error {
	if measurement < 0 {
		v.measurement = 0
		v.valid = false

		return fmt.Errorf(ErrVelocityMustBeZeroOrGreater)
	}

	v.measurement = measurement
	v.valid = true

	return nil
}

// ToFps converts the speed to feet per second.
func (v *VelocityKts) ToFps() VelocityFps {
	return VelocityFps{
		measurement: v.Fps(),
		valid:       v.valid,
	}
}

// ToKts converts the speed to knots.
func (v *VelocityKts) ToKts() VelocityKts {
	return *v
}

// ToKph converts the speed to kilometers per hour.
func (v *VelocityKts) ToKph() VelocityKph {
	return VelocityKph{
		measurement: v.Kph(),
		valid:       v.valid,
	}
}

// ToMph converts the speed to miles per hour.
func (v *VelocityKts) ToMph() VelocityMph {
	return VelocityMph{
		measurement: v.Mph(),
		valid:       v.valid,
	}
}

// ToMps converts the speed to meters per second.
func (v *VelocityKts) ToMps() VelocityMps {
	return VelocityMps{
		measurement: v.Mps(),
		valid:       v.valid,
	}
}

// Units return the units of the speed.
func (v *VelocityKts) Units() VelocityUnit {
	return Knots
}

// Valid returns true if the velocity is zero or greater.
func (v *VelocityKts) Valid() bool {
	return v.valid
}
