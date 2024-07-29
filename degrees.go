package wx

import "math"

// Degrees represents a normalized angle in degrees
// between 0 and 359.9 degrees.
type Degrees struct {
	degrees float64
}

// Degrees returns the angle in degrees.
func (d Degrees) Degrees() float64 {
	return d.degrees
}

// Radians returns the angle in radians.
func (d Degrees) Radians() float64 {
	return d.degrees * math.Pi / 180
}

// Add adds two angles together and returns a new
// normalized angle.
func (d Degrees) Add(d2 Degrees) Degrees {
	return NewDegrees(d.degrees + d2.degrees)
}

// Divide divides an angle by a scalar and returns
// a new normalized angle.
func (d Degrees) Divide(scalar float64) Degrees {
	if scalar == 0 {
		return NewDegrees(0)
	}

	return NewDegrees(d.degrees / scalar)
}

// Multiply multiplies an angle by a scalar and returns
// a new normalized angle.
func (d Degrees) Multiply(scalar float64) Degrees {
	return NewDegrees(d.degrees * scalar)
}

// Sub subtracts two angles and returns a new
// normalized angle.
func (d Degrees) Sub(d2 Degrees) Degrees {
	return NewDegrees(d.degrees - d2.degrees)
}

// UnitVector returns the unit vector of the degrees.
func (d Degrees) UnitVector() (x, y float64) {
	// Convert the wind direction to radians.
	rad := d.degrees * math.Pi / 180

	// Calculate the x and y components of the unit vector.
	x = -math.Sin(rad)
	y = -math.Cos(rad)

	return x, y
}

// NewDegrees creates a new degrees value from a float64.
// The value is normalized to be between 0 and 359.9 degrees.
func NewDegrees(deg float64) Degrees {
	return Degrees{degrees: naiveDegrees(deg)}
}

// naiveDegrees takes a float64 and determines the degrees
// using a modulus operation.
func naiveDegrees(d float64) float64 {
	// Normalize the angle to be between 0 and 359.9 degrees.
	d = math.Mod(d, 360)

	// If the angle is negative, add 360 to make it positive.
	if d < 0 {
		d += 360
	}

	return d
}
