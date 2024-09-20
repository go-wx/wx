package wx

import "fmt"

// VelocityType represents a unit of velocity.
type velocityType int

// Velocity units.
const (
	fps velocityType = iota // feet per second
	kts                     // knots
	kph                     // kilometers per hour
	mph                     // miles per hour
	mps                     // meters per second
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
	Fps = VelocityUnit{fps}
	Kts = VelocityUnit{kts}
	Kph = VelocityUnit{kph}
	Mph = VelocityUnit{mph}
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
func NewVelocity(measurement float64, u VelocityUnit) Velocity {
	return Velocity{
		measurement: measurement,
		unit:        u,
		valid:       measurement >= 0}
}

// Fps returns the velocity in feet per second.
func (v Velocity) Fps() float64 {
	switch v.unit.velocityType {
	case fps:
		return v.measurement
	case kts:
		return v.measurement * ktsToFps
	case kph:
		return v.measurement * kphToFps
	case mph:
		return v.measurement * mphToFps
	case mps:
		return v.measurement * mpsToFps
	}
	return 0
}

// Kts returns the velocity in knots.
func (v Velocity) Kts() float64 {
	switch v.unit.velocityType {
	case fps:
		return v.measurement / ktsToFps
	case kts:
		return v.measurement
	case kph:
		return v.measurement * kphToKts
	case mph:
		return v.measurement * mphToKts
	case mps:
		return v.measurement * mpsToKts
	}
	return 0
}

// Kph returns the velocity in kilometers per hour.
func (v Velocity) Kph() float64 {
	switch v.unit.velocityType {
	case fps:
		return v.measurement / kphToFps
	case kts:
		return v.measurement / kphToKts
	case kph:
		return v.measurement
	case mph:
		return v.measurement * mphToKph
	case mps:
		return v.measurement * mpsToKph
	}
	return 0
}

// Mph returns the velocity in miles per hour.
func (v Velocity) Mph() float64 {
	switch v.unit.velocityType {
	case fps:
		return v.measurement / mphToFps
	case kts:
		return v.measurement / mphToKts
	case kph:
		return v.measurement / mphToKph
	case mph:
		return v.measurement
	case mps:
		return v.measurement * mpsToMph
	}
	return 0
}

// Mps returns the velocity in meters per second.
func (v Velocity) Mps() float64 {
	switch v.unit.velocityType {
	case fps:
		return v.measurement / mpsToFps
	case kts:
		return v.measurement / mpsToKts
	case kph:
		return v.measurement / mpsToKph
	case mph:
		return v.measurement / mpsToMph
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

// Constants for converting between velocity units.
const (
	ktsToFps = 1.68781
	kphToFps = 0.911344
	mphToFps = 1.46667
	mpsToFps = 3.28084

	kphToKts = 0.539957
	mphToKts = 0.868976
	mpsToKts = 1.94384

	mphToKph = 1.60934
	mpsToKph = 60 * 60 / 1000

	mpsToMph = 2.23694
)
