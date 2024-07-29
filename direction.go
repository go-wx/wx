package wx

// WindDirection represents the direction of the wind.
// The value is between 0 and 359.9 degrees and is the
// direction the wind is coming from, not the direction
// the wind is blowing to in degrees from true north.
type WindDirection struct {
	degrees Degrees
}

// Degrees returns the wind direction in degrees from true north.
func (wd WindDirection) Degrees() Degrees {
	return wd.degrees
}

// From returns the direction the wind is coming from.
func (wd WindDirection) From() Degrees {
	return wd.degrees
}

// Radians returns the wind direction in radians from true north.
func (wd WindDirection) Radians() float64 {
	return wd.degrees.Radians()
}

// To returns the direction the wind is blowing to.
func (wd WindDirection) To() Degrees {
	return NewDegrees(wd.degrees.degrees + 180)
}

// UnitVector returns the unit vector of the wind direction.
//
// The x-component is the east-west component.
// The y-component is the north-south component.
//
// The wind direction is the direction the wind is
// coming from, the x component is -sin(wind direction)
// and the y component is -cos(wind direction).
func (wd WindDirection) UnitVector() (x, y float64) {
	return wd.degrees.UnitVector()
}

// NewWindDirection creates a new WindDirection value from a float64.
// The value is normalized to be between 0 and 359.9 degrees.
func NewWindDirection(deg float64) WindDirection {
	return WindDirection{degrees: NewDegrees(deg)}
}