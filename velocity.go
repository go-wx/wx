package wx

import "fmt"

// VelocityType represents a unit of velocity.
type velocityType uint8

// Velocity units.
// The values start at 1 to avoid a zero value if
// a developer creates an invalid VelocityUnit.
//
// For example, if a developer creates a VelocityUnit
// with a value of zero, the zero value would be
// considered invalid.
// Example:
//
//	var v VelocityUnit
//	if !v.Valid() {
//		// This will be true.
//	}
const (
	fps velocityType = iota + 1 // feet per second
	kts                         // knots
	kph                         // kilometers per hour
	mph                         // miles per hour
	mps                         // meters per second
)

// String returns the string representation of the velocity unit.
func (v velocityType) String() string {
	switch v {
	case fps:
		return "fps"
	case kts:
		return "kts"
	case kph:
		return "kph"
	case mph:
		return "mph"
	case mps:
		return "mps"
	}
	return ""
}

// VelocityUnit is a unit of velocity (e.g., mph, kts, mps, fps, etc.)
type VelocityUnit struct {
	velocityType
}

// String returns the string representation of the velocity unit.
func (v VelocityUnit) String() string {
	return v.velocityType.String()
}

// Velocity units.
var (
	// Fps is feet per second.
	Fps = VelocityUnit{fps}
	// Kts is knots.
	Kts = VelocityUnit{kts}
	// Kph is kilometers per hour.
	Kph = VelocityUnit{kph}
	// Mph is miles per hour.
	Mph = VelocityUnit{mph}
	// Mps is meters per second.
	Mps = VelocityUnit{mps}
)

// Velocity represents a velocity measurement.
type Velocity struct {
	measurement float64
	unit        VelocityUnit
	valid       bool
}

// String returns the string representation of the velocity.
func (v Velocity) String() string {
	if !v.valid {
		return "invalid velocity"
	}
	return fmt.Sprintf("%.1f %s", v.measurement, v.unit)
}

// Valid returns true if the velocity is valid.
func (v Velocity) Valid() bool {
	return v.valid
}

// NewVelocity creates a new Velocity value.
func NewVelocity(measurement float64, unit VelocityUnit) Velocity {
	switch unit.velocityType {
	case fps:
		unit = Fps
	case kts:
		unit = Kts
	case kph:
		unit = Kph
	case mph:
		unit = Mph
	case mps:
		unit = Mps
	default:
		return Velocity{valid: false}
	}

	return Velocity{
		measurement: measurement,
		unit:        unit,
		valid:       measurement >= 0}
}

// Fps returns the velocity in feet per second.
func (v Velocity) Fps() float64 {
	switch v.unit.velocityType {
	case fps:
		return v.measurement
	case kts:
		return v.measurement * feetPerNauticalMile / 3600
	case kph:
		return v.measurement * feetPerMeter * 1000 / 3600
	case mph:
		return v.measurement * feetPerStatuteMile / 3600
	case mps:
		return v.measurement * feetPerMeter
	}
	return 0
}

// Kph returns the velocity in kilometers per hour.
func (v Velocity) Kph() float64 {
	switch v.unit.velocityType {
	case fps:
		return v.measurement * feetPerMeter * 1000 / 3600
	case kts:
		return v.measurement * feetPerNauticalMile / feetPerMeter / 1000
	case kph:
		return v.measurement
	case mph:
		return v.measurement * feetPerStatuteMile / 3600 / feetPerMeter * 1000
	case mps:
		return v.measurement * 3600 / 1000
	}
	return 0
}

// Kts returns the velocity in knots.
func (v Velocity) Kts() float64 {
	switch v.unit.velocityType {
	case fps:
		return v.measurement * 3600 / feetPerNauticalMile
	case kts:
		return v.measurement
	case kph:
		return v.measurement * feetPerMeter * 1000 / feetPerNauticalMile
	case mph:
		return v.measurement * feetPerStatuteMile / feetPerNauticalMile
	case mps:
		return v.measurement * feetPerMeter / feetPerNauticalMile
	}
	return 0
}

// Mph returns the velocity in miles per hour.
func (v Velocity) Mph() float64 {
	switch v.unit.velocityType {
	case fps:
		return v.measurement * 3600 / feetPerStatuteMile
	case kts:
		return v.measurement * feetPerNauticalMile / feetPerStatuteMile
	case kph:
		return v.measurement * feetPerMeter * 1000 / feetPerStatuteMile
	case mph:
		return v.measurement
	case mps:
		return v.measurement * feetPerMeter * 3600 / feetPerStatuteMile
	}
	return 0
}

// Mps returns the velocity in meters per second.
func (v Velocity) Mps() float64 {
	switch v.unit.velocityType {
	case fps:
		return v.measurement / feetPerMeter
	case kts:
		return v.measurement * feetPerNauticalMile / feetPerMeter / 3600
	case kph:
		return v.measurement * 1000 / 3600
	case mph:
		return v.measurement * feetPerStatuteMile / feetPerMeter / 3600
	case mps:
		return v.measurement
	}

	return 0
}

// ToFps converts the velocity to feet per second.
func (v Velocity) ToFps() Velocity {
	return NewVelocity(v.Fps(), Fps)
}

// ToKts converts the velocity to knots.
func (v Velocity) ToKts() Velocity {
	return NewVelocity(v.Kts(), Kts)
}

// ToKph converts the velocity to kilometers per hour.
func (v Velocity) ToKph() Velocity {
	return NewVelocity(v.Kph(), Kph)
}

// ToMph converts the velocity to miles per hour.
func (v Velocity) ToMph() Velocity {
	return NewVelocity(v.Mph(), Mph)
}

// ToMps converts the velocity to meters per second.
func (v Velocity) ToMps() Velocity {
	return NewVelocity(v.Mps(), Mps)
}
