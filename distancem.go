package wx

type DistanceM struct {
	meters float64
	valid  bool
}

// NewDistanceM creates a new DistanceM.
func NewDistanceM(meters float64) DistanceM {
	return DistanceM{
		meters: meters,
		valid:  true,
	}
}

// FT returns the distance in feet.
func (d *DistanceM) FT() float64 {
	return d.meters * feetPerMeter
}

// KM returns the distance in kilometers.
func (d *DistanceM) KM() float64 {
	return d.meters / 1000
}

// M returns the distance in meters.
func (d *DistanceM) M() float64 {
	return d.meters
}

// NM returns the distance in nautical miles.
func (d *DistanceM) NM() float64 {
	return d.meters / feetPerNauticalMile
}

// SM returns the distance in statute miles.
func (d *DistanceM) SM() float64 {
	return d.meters / feetPerStatuteMile
}

// Set sets the distance in meters.
func (d *DistanceM) Set(meters float64) error {
	d.meters = meters
	d.valid = true
	return nil
}

// ToFT returns the distance in feet.
func (d *DistanceM) ToFT() DistanceFT {
	return DistanceFT{
		feet:  d.FT(),
		valid: d.valid,
	}
}

// ToKM returns the distance in kilometers.
func (d *DistanceM) ToKM() DistanceKM {
	return DistanceKM{
		kilometers: d.KM(),
		valid:      d.valid,
	}
}

// ToNM returns the distance in nautical miles.
func (d *DistanceM) ToNM() DistanceNM {
	return DistanceNM{
		nauticalMiles: d.NM(),
		valid:         d.valid,
	}
}

// ToSM returns the distance in statute miles.
func (d *DistanceM) ToSM() DistanceSM {
	return DistanceSM{
		statuteMiles: d.SM(),
		valid:        d.valid,
	}
}

// Units returns the units of the distance.
func (d *DistanceM) Units() DistanceUnit {
	return M
}

// Valid returns true if the distance is zero or greater.
func (d *DistanceM) Valid() bool {
	return d.valid
}
