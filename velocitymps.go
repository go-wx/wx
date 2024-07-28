package wx

import "fmt"

// VelocityMps represents a velocity in meters per second.
type VelocityMps struct {
	measurement float64
	valid       bool
}

// Fps returns the speed in feet per second.
func (v *VelocityMps) Fps() float64 {
	return v.measurement * feetPerMeter
}

// Kts returns the speed in knots.
func (v *VelocityMps) Kts() float64 {
	return v.measurement * feetPerMeter / feetPerNauticalMile * secondsInHour
}

// Kph returns the speed in kilometers per hour.
func (v *VelocityMps) Kph() float64 {
	return v.measurement * secondsInHour / 1000
}

// Mph returns the speed in miles per hour.
func (v *VelocityMps) Mph() float64 {
	return v.Fps() / feetPerStatuteMile * secondsInHour
}

// Mps returns the speed in meters per second.
func (v *VelocityMps) Mps() float64 {
	return v.measurement
}

// Set sets the velocity to the specified value.
func (v *VelocityMps) Set(measurement float64) error {
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
func (v *VelocityMps) ToFps() VelocityFps {
	return VelocityFps{
		measurement: v.Fps(),
		valid:       v.valid,
	}
}

// ToKts converts the speed to knots.
func (v *VelocityMps) ToKts() VelocityKts {
	return VelocityKts{
		measurement: v.Kts(),
		valid:       v.valid,
	}
}

// ToKph converts the speed to kilometers per hour.
func (v *VelocityMps) ToKph() VelocityKph {
	return VelocityKph{
		measurement: v.Kph(),
		valid:       v.valid,
	}
}

// ToMph converts the speed to miles per hour.
func (v *VelocityMps) ToMph() VelocityMph {
	return VelocityMph{
		measurement: v.Mph(),
		valid:       v.valid,
	}
}

// ToMps converts the speed to meters per second.
func (v *VelocityMps) ToMps() VelocityMps {
	return *v
}

// Units return the units of the speed.
func (v *VelocityMps) Units() VelocityUnit {
	return MetersPerSecond
}

// Valid returns true if the velocity is zero or greater.
func (v *VelocityMps) Valid() bool {
	return v.valid
}
