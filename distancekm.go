package wx

// DistanceKM represents a distance in kilometers.
type DistanceKM struct {
	kilometers float64
	valid      bool
}

// NewDistanceKM creates a new DistanceKM.
func NewDistanceKM(kilometers float64) DistanceKM {
	return DistanceKM{
		kilometers: kilometers,
		valid:      kilometers >= 0,
	}
}

// FT returns the distance in feet.
func (d *DistanceKM) FT() float64 {
	return d.kilometers * feetPerMeter * 1000
}

// KM returns the distance in kilometers.
func (d *DistanceKM) KM() float64 {
	return d.kilometers
}

// M returns the distance in meters.
func (d *DistanceKM) M() float64 {
	return d.kilometers * 1000
}

// NM returns the distance in nautical miles.
func (d *DistanceKM) NM() float64 {
	return d.kilometers * 1000 * feetPerMeter / feetPerNauticalMile
}

// SM returns the distance in statute miles.
func (d *DistanceKM) SM() float64 {
	return d.kilometers * 1000 * feetPerMeter / feetPerStatuteMile
}

// Set sets the distance in kilometers.
func (d *DistanceKM) Set(kilometers float64) error {
	if kilometers < 0 {
		d.valid = false
		d.kilometers = 0
		return NewWxErr("distance cannot be negative", "DistanceKM.Set")
	}

	d.kilometers = kilometers
	d.valid = true

	return nil
}

// ToFT returns the distance in feet.
func (d *DistanceKM) ToFT() DistanceFT {
	return DistanceFT{
		feet:  d.FT(),
		valid: d.valid,
	}
}

// ToKM returns the distance in kilometers.
func (d *DistanceKM) ToKM() DistanceKM {
	return DistanceKM{
		kilometers: d.kilometers,
		valid:      d.valid,
	}
}

// ToM returns the distance in meters.
func (d *DistanceKM) ToM() DistanceM {
	return DistanceM{
		meters: d.M(),
		valid:  d.valid,
	}
}

// ToNM returns the distance in nautical miles.
func (d *DistanceKM) ToNM() DistanceNM {
	return DistanceNM{
		nauticalMiles: d.NM(),
		valid:         d.valid,
	}
}

// ToSM returns the distance in statute miles.
func (d *DistanceKM) ToSM() DistanceSM {
	return DistanceSM{
		statuteMiles: d.SM(),
		valid:        d.valid,
	}
}

// Units returns the units of the distance.
func (d *DistanceKM) Units() DistanceUnit {
	return KM
}

// Valid returns true if the distance is zero or greater.
func (d *DistanceKM) Valid() bool {
	return d.valid
}
