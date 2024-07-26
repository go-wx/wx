package wx

// PressureHPa is a type for pressure in hectopascals.
type PressureHPa struct {
	measurement float64
	valid       bool
}

// PressureKPa is a type for pressure in kilopascals.
type PressureKPa struct {
	measurement float64
	valid       bool
}

// PressurePa is a type for pressure in pascals.
type PressurePa struct {
	measurement float64
	valid       bool
}

func (p PressureHPa) HPa() float64 {
	return p.measurement
}
