package wx

const (
	TempAbsoluteZeroRF = -459.67
)

// TempR represents a temperature measurement in the Rankine scale.
// TempR implements the Temp interface.
type TempR struct {
	measurement float64
}
