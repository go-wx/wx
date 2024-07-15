package wx

// PressureHPa is a type for pressure in hectopascals.
type PressureHPa struct {
	measurement float64
	valid       bool
}

// PressureKPa is a type for pressure in kilopascals.
type PressureKPa struct {
	measurement float64
}

// PressurePa is a type for pressure in pascals.
type PressurePa struct {
	measurement float64
}
