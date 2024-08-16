package wx

type DistanceFT struct {
	feet  float64
	valid bool
}

// NewDistanceFT creates a new DistanceFT.
func NewDistanceFT(feet float64) DistanceFT {
	return DistanceFT{
		feet:  feet,
		valid: true, // feet can be negative (e.g., altitude below sea level)
	}
}

// FT returns the distance in feet.
func (d *DistanceFT) FT() float64 {
	return d.feet
}

// KM returns the distance in kilometers.
func (d *DistanceFT) KM() float64 {
	return d.feet / feetPerMeter / 1000
}

// M returns the distance in meters.
func (d *DistanceFT) M() float64 {
	return d.feet / feetPerMeter
}

// NM returns the distance in nautical miles.
func (d *DistanceFT) NM() float64 {
	return d.feet / feetPerNauticalMile
}

// SM returns the distance in statute miles.
func (d *DistanceFT) SM() float64 {
	return d.feet / feetPerStatuteMile
}

// Set sets the distance in feet.
func (d *DistanceFT) Set(feet float64) error {
	d.feet = feet
	d.valid = true
	return nil
}

// ToFT returns the distance in feet.
func (d *DistanceFT) ToFT() DistanceFT {
	return DistanceFT{
		feet:  d.FT(),
		valid: d.valid,
	}
}

// ToKM returns the distance in kilometers.
func (d *DistanceFT) ToKM() DistanceKM {
	return DistanceKM{
		kilometers: d.KM(),
		valid:      d.valid,
	}
}

// ToM returns the distance in meters.
func (d *DistanceFT) ToM() DistanceM {
	return DistanceM{
		meters: d.M(),
		valid:  d.valid,
	}
}

// ToNM returns the distance in nautical miles.
func (d *DistanceFT) ToNM() DistanceNM {
	return DistanceNM{
		nauticalMiles: d.NM(),
		valid:         d.valid,
	}
}

// ToSM returns the distance in statute miles.
func (d *DistanceFT) ToSM() DistanceSM {
	return DistanceSM{
		statuteMiles: d.SM(),
		valid:        d.valid,
	}
}

// Units returns the units of the distance.
func (d *DistanceFT) Units() DistanceUnit {
	return FT
}

// Valid returns true if the distance is zero or greater.
func (d *DistanceFT) Valid() bool {
	return d.valid
}
