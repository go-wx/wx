package wx

type DistanceUnit struct {
	string
}

var (
	FT = DistanceUnit{"ft"}
	KM = DistanceUnit{"km"}
	NM = DistanceUnit{"nm"}
	M  = DistanceUnit{"m"}
	SM = DistanceUnit{"sm"}
)

// Distance represents a distance.
type Distance interface {
	// FT returns the distance in feet.
	FT() float64

	// KM returns the distance in kilometers.
	KM() float64

	// NM returns the distance in nautical miles.
	NM() float64

	// M returns the distance in meters.
	M() float64

	// SM returns the distance in statute miles.
	SM() float64

	// Set sets the distance in feet.
	Set(float64) error

	// ToFT converts the distance to feet.
	ToFT() DistanceFT

	// ToKM converts the distance to kilometers.
	ToKM() DistanceKM

	// ToNM converts the distance to nautical miles.
	ToNM() DistanceNM

	// ToM converts the distance to meters.
	ToM() DistanceM

	// ToSM converts the distance to statute miles.
	ToSM() DistanceSM
}
