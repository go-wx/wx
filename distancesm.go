package wx

type DistanceSM struct {
	statuteMiles float64
	valid        bool
}

// NewDistanceSM creates a new DistanceSM.
func NewDistanceSM(statuteMiles float64) DistanceSM {
	return DistanceSM{
		statuteMiles: statuteMiles,
		valid:        statuteMiles >= 0,
	}
}

// FT returns the distance in feet.
func (d *DistanceSM) FT() float64 {
	return d.statuteMiles * feetPerStatuteMile
}

// KM returns the distance in kilometers.
func (d *DistanceSM) KM() float64 {
	return d.statuteMiles * feetPerStatuteMile / feetPerMeter / 1000
}

// M returns the distance in meters.
func (d *DistanceSM) M() float64 {
	return d.statuteMiles * feetPerStatuteMile / feetPerMeter
}

// NM returns the distance in nautical miles.
func (d *DistanceSM) NM() float64 {
	return d.statuteMiles * feetPerStatuteMile / feetPerNauticalMile
}

// SM returns the distance in statute miles.
func (d *DistanceSM) SM() float64 {
	return d.statuteMiles
}

// Set sets the distance in statute miles.
func (d *DistanceSM) Set(statuteMiles float64) error {
	if statuteMiles < 0 {
		d.valid = false
		d.statuteMiles = 0
		return NewWxErr("distance cannot be negative", "DistanceSM.Set")
	}

	d.statuteMiles = statuteMiles
	d.valid = true

	return nil
}

// ToFT returns the distance in feet.
func (d *DistanceSM) ToFT() DistanceFT {
	return DistanceFT{
		feet:  d.FT(),
		valid: d.valid,
	}
}

// ToKM returns the distance in kilometers.
func (d *DistanceSM) ToKM() DistanceKM {
	return DistanceKM{
		kilometers: d.KM(),
		valid:      d.valid,
	}
}

// ToM returns the distance in meters.
func (d *DistanceSM) ToM() DistanceM {
	return DistanceM{
		meters: d.M(),
		valid:  d.valid,
	}
}

// ToNM returns the distance in nautical miles.
func (d *DistanceSM) ToNM() DistanceNM {
	return DistanceNM{
		nauticalMiles: d.NM(),
		valid:         d.valid,
	}
}

// ToSM returns the distance in statute miles.
func (d *DistanceSM) ToSM() DistanceSM {
	return DistanceSM{
		statuteMiles: d.SM(),
		valid:        d.valid,
	}
}

// Units returns the units of the distance.
func (d *DistanceSM) Units() DistanceUnit {
	return SM
}

// Valid returns true if the distance is zero or greater.
func (d *DistanceSM) Valid() bool {
	return d.valid
}
