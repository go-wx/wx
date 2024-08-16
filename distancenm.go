package wx

// DistanceNM represents a distance in nautical miles.
type DistanceNM struct {
	nauticalMiles float64
	valid         bool
}

// NewDistanceNM creates a new DistanceNM.
func NewDistanceNM(nauticalMiles float64) DistanceNM {
	return DistanceNM{
		nauticalMiles: nauticalMiles,
		valid:         nauticalMiles >= 0,
	}
}

// FT returns the distance in feet.
func (d *DistanceNM) FT() float64 {
	return d.nauticalMiles * feetPerNauticalMile
}

// KM returns the distance in kilometers.
func (d *DistanceNM) KM() float64 {
	return d.nauticalMiles * feetPerNauticalMile / feetPerMeter / 1000
}

// M returns the distance in meters.
func (d *DistanceNM) M() float64 {
	return d.nauticalMiles * feetPerNauticalMile / feetPerMeter
}

// NM returns the distance in nautical miles.
func (d *DistanceNM) NM() float64 {
	return d.nauticalMiles
}

// SM returns the distance in statute miles.
func (d *DistanceNM) SM() float64 {
	return d.nauticalMiles * feetPerNauticalMile / feetPerStatuteMile
}

// Set sets the distance in nautical miles.
func (d *DistanceNM) Set(nauticalMiles float64) error {
	if nauticalMiles < 0 {
		d.valid = false
		d.nauticalMiles = 0
		return NewWxErr("distance cannot be negative", "DistanceNM.Set")
	}

	d.nauticalMiles = nauticalMiles
	d.valid = true

	return nil
}

// ToFT returns the distance in feet.
func (d *DistanceNM) ToFT() DistanceFT {
	return DistanceFT{
		feet:  d.FT(),
		valid: d.valid,
	}
}

// ToKM returns the distance in kilometers.
func (d *DistanceNM) ToKM() DistanceKM {
	return DistanceKM{
		kilometers: d.KM(),
		valid:      d.valid,
	}
}

// ToM returns the distance in meters.
func (d *DistanceNM) ToM() DistanceM {
	return DistanceM{
		meters: d.M(),
		valid:  d.valid,
	}
}

// ToNM returns the distance in nautical miles.
func (d *DistanceNM) ToNM() DistanceNM {
	return DistanceNM{
		nauticalMiles: d.NM(),
		valid:         d.valid,
	}
}

// ToSM returns the distance in statute miles.
func (d *DistanceNM) ToSM() DistanceSM {
	return DistanceSM{
		statuteMiles: d.SM(),
		valid:        d.valid,
	}
}

// Units returns the units of the distance.
func (d *DistanceNM) Units() DistanceUnit {
	return NM
}

// Valid returns true if the distance is zero or greater.
func (d *DistanceNM) Valid() bool {
	return d.valid
}
