package wx

import (
	"fmt"
	"math"
)

type distanceType uint8

// Distance units.
const (
	feet distanceType = iota
	kilometers
	nauticalMiles
	meters
	statuteMiles
	parsec
)

// String returns the string representation of the distance unit.
func (d distanceType) String() string {
	switch d {
	case feet:
		return "ft"
	case kilometers:
		return "km"
	case nauticalMiles:
		return "NM"
	case meters:
		return "m"
	case statuteMiles:
		return "SM"
	case parsec:
		return "pc"
	}

	return ""
}

// DistanceUnit represents a unit of distance.
type DistanceUnit struct {
	distanceType
}

var (
	// Feet represent a distance in feet.
	Feet = DistanceUnit{feet}
	// Kilometers represents a distance in kilometers.
	Kilometers = DistanceUnit{kilometers}
	// NauticalMiles represents a distance in nautical miles.
	NauticalMiles = DistanceUnit{nauticalMiles}
	// Meters represents a distance in meters.
	Meters = DistanceUnit{meters}
	// StatuteMiles represents a distance in statute miles.
	StatuteMiles = DistanceUnit{statuteMiles}
	// Parsec represents a distance in parsecs.
	Parsec = DistanceUnit{parsec}
)

const (
	feetPerNauticalMile = 6_076.11549
	feetPerStatuteMile  = 5_280.0
	feetPerMeter        = 3.280839895
	metersPerParsec     = 3.08567758e16
)

// String returns the string representation of the distance unit.
func (d DistanceUnit) String() string {
	return d.distanceType.String()
}

// Distance represents a distance measurement.
type Distance struct {
	measurement float64
	unit        DistanceUnit
	valid       bool
}

// NewDistance creates a new distance measurement.
func NewDistance(measurement float64, unit DistanceUnit) Distance {
	return Distance{
		measurement: measurement,
		unit:        unit,
		// There are no lines in the opposite direction.
		valid: measurement >= 0.0,
	}
}

// Abs returns the absolute value of the distance.
func (d Distance) Abs() Distance {
	return NewDistance(math.Abs(d.measurement), d.unit)
}

// Add returns the sum of two distances.
func (d Distance) Add(d2 Distance) Distance {
	// Convert d2 to the same unit as d.
	d2 = d2.to(d.unit)

	// Add the measurements.
	return NewDistance(d.measurement+d2.measurement, d.unit)
}

// Sub returns the difference of two distances.
//
// It is recommended you check for negative values in
// your application.
//
// For example, if you are subtracting a distance
// from another distance, you may want to ensure
// the result is not negative.
//
// Example:
//
//	d := NewDistance(1.0, Feet)
//	d2 := NewDistance(2.0, Feet)
//	d3 := d.Sub(d2)
//	if !d3.Valid() {
//	  // Handle negative distance.
//	}
func (d Distance) Sub(d2 Distance) Distance {
	// Convert d2 to the same unit as d.
	d2 = d2.to(d.unit)

	// Subtract the measurements.

	return NewDistance(d.measurement-d2.measurement, d.unit)
}

// FT returns the distance in feet.
func (d Distance) FT() float64 {
	switch d.unit.distanceType {
	case feet:
		return d.measurement
	case kilometers:
		return d.measurement * feetPerMeter * 1000.0
	case nauticalMiles:
		return d.measurement * feetPerNauticalMile
	case meters:
		return d.measurement * feetPerMeter
	case statuteMiles:
		return d.measurement * feetPerStatuteMile
	case parsec:
		return d.measurement * metersPerParsec * feetPerMeter
	}

	return 0.0
}

// KM returns the distance in kilometers.
func (d Distance) KM() float64 {
	switch d.unit.distanceType {
	case feet:
		return d.measurement / feetPerMeter / 1000.0
	case kilometers:
		return d.measurement
	case nauticalMiles:
		return d.measurement * feetPerNauticalMile / feetPerMeter / 1000.0
	case meters:
		return d.measurement / 1000.0
	case statuteMiles:
		return d.measurement * feetPerStatuteMile / feetPerMeter / 1000.0
	case parsec:
		return d.measurement * metersPerParsec / 1000.0
	}

	return 0.0
}

// NM returns the distance in nautical miles.
func (d Distance) NM() float64 {
	switch d.unit.distanceType {
	case feet:
		return d.measurement / feetPerNauticalMile
	case kilometers:
		return d.measurement * feetPerMeter * 1000.0 / feetPerNauticalMile
	case nauticalMiles:
		return d.measurement
	case meters:
		return d.measurement / feetPerMeter / feetPerNauticalMile
	case statuteMiles:
		return d.measurement * feetPerStatuteMile / feetPerMeter / feetPerNauticalMile
	case parsec:
		return d.measurement * metersPerParsec / feetPerMeter / feetPerNauticalMile
	}

	return 0.0
}

// M returns the distance in meters.
func (d Distance) M() float64 {
	switch d.unit.distanceType {
	case feet:
		return d.measurement / feetPerMeter
	case kilometers:
		return d.measurement * 1000.0
	case nauticalMiles:
		return d.measurement * feetPerNauticalMile / feetPerMeter
	case meters:
		return d.measurement
	case statuteMiles:
		return d.measurement * feetPerStatuteMile / feetPerMeter
	case parsec:
		return d.measurement * metersPerParsec
	}

	return 0.0
}

// SM returns the distance in statute miles.
func (d Distance) SM() float64 {
	switch d.unit.distanceType {
	case feet:
		return d.measurement / feetPerStatuteMile
	case kilometers:
		return d.measurement * feetPerMeter * 1000.0 / feetPerStatuteMile
	case nauticalMiles:
		return d.measurement * feetPerNauticalMile / feetPerMeter / feetPerStatuteMile
	case meters:
		return d.measurement / feetPerMeter / feetPerStatuteMile
	case statuteMiles:
		return d.measurement
	case parsec:
		return d.measurement * metersPerParsec / feetPerMeter / feetPerStatuteMile
	}

	return 0.0
}

// Parsec returns the distance in parsecs.
func (d Distance) Parsec() float64 {
	switch d.unit.distanceType {
	case feet:
		return d.measurement / feetPerMeter / metersPerParsec
	case kilometers:
		return d.measurement / metersPerParsec * 1000.0
	case nauticalMiles:
		return d.measurement / feetPerNauticalMile * feetPerMeter / metersPerParsec
	case meters:
		return d.measurement / metersPerParsec
	case statuteMiles:
		return d.measurement / feetPerStatuteMile / metersPerParsec
	case parsec:
		return d.measurement
	}

	return 0.0
}

// String returns the string representation of the distance.
// The distance is formatted to two decimal places and the
// measurement must be zero or greater.
func (d Distance) String() string {
	if !d.valid {
		return fmt.Sprintf("invalid %s", d.unit.String())
	}

	return fmt.Sprintf("%0.2f %s", d.measurement,
		d.unit.String())
}

// to converts a distance to the specified unit.
func (d Distance) to(unit DistanceUnit) Distance {
	switch unit.distanceType {
	case feet:
		return d.ToFt()
	case kilometers:
		return d.ToKM()
	case nauticalMiles:
		return d.ToNM()
	case meters:
		return d.ToM()
	case statuteMiles:
		return d.ToSM()
	case parsec:
		return d.ToPc()
	}

	return Distance{valid: false}
}

// ToFt converts the distance to feet.
func (d Distance) ToFt() Distance {
	return NewDistance(d.FT(), Feet)
}

// ToKM converts the distance to kilometers.
func (d Distance) ToKM() Distance {
	return NewDistance(d.KM(), Kilometers)
}

// ToNM converts the distance to nautical miles.
func (d Distance) ToNM() Distance {
	return NewDistance(d.NM(), NauticalMiles)
}

// ToM converts the distance to meters.
func (d Distance) ToM() Distance {
	return NewDistance(d.M(), Meters)
}

// ToSM converts the distance to statute miles.
func (d Distance) ToSM() Distance {
	return NewDistance(d.SM(), StatuteMiles)
}

// ToPc converts the distance to parsecs.
func (d Distance) ToPc() Distance {
	return NewDistance(d.Parsec(), Parsec)
}

// Valid returns true if the distance is valid.
func (d Distance) Valid() bool {
	return d.valid
}
