package wx

import "fmt"

// VelocityMph measures velocity in miles per hour.
type VelocityMph struct {
	measurement float64
}

// Fps returns the speed in feet per second.
func (v *VelocityMph) Fps() float64 {
	return v.measurement * feetPerStatuteMile / secondsInHour
}

// Kts returns the speed in knots.
func (v *VelocityMph) Kts() float64 {
	return v.measurement * (feetPerStatuteMile / feetPerNauticalMile)
}

// Kph returns the speed in kilometers per hour.
func (v *VelocityMph) Kph() float64 {
	return v.Fps() / feetPerMeter / 1000 * secondsInHour
}

// Mph returns the speed in miles per hour.
func (v *VelocityMph) Mph() float64 {
	return v.measurement
}

// Mps returns the speed in meters per second.
func (v *VelocityMph) Mps() float64 {
	return v.Fps() / feetPerMeter
}

// Set sets the velocity to the specified value.
func (v *VelocityMph) Set(measurement float64) error {
	if measurement < 0 {
		v.measurement = 0

		return fmt.Errorf(ErrVelocityMustBeZeroOrGreater)
	}

	v.measurement = measurement

	return nil
}

// ToFps converts the speed to feet per second.
func (v *VelocityMph) ToFps() VelocityFps {
	return VelocityFps{
		measurement: v.Fps(),
	}
}

// ToKts converts the speed to knots.
func (v *VelocityMph) ToKts() VelocityKts {
	return VelocityKts{
		measurement: v.Kts(),
	}
}

// ToKph converts the speed to kilometers per hour.
func (v *VelocityMph) ToKph() VelocityKph {
	return VelocityKph{
		measurement: v.Kph(),
	}
}

// ToMph converts the speed to miles per hour.
func (v *VelocityMph) ToMph() VelocityMph {
	return *v
}

// ToMps converts the speed to meters per second.
func (v *VelocityMph) ToMps() VelocityMps {
	return VelocityMps{
		measurement: v.Mps(),
	}
}

// Units return the units of the speed.
func (v *VelocityMph) Units() VelocityUnit {
	return MilesPerHour
}

// Valid returns true if the velocity is zero or greater.
func (v *VelocityMph) Valid() bool {
	return v.measurement >= 0
}
