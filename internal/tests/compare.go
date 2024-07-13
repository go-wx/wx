package tests

const (
	Tolerance = 1e-9
)

// CloseEnough returns true if a is within tolerance of b.
func CloseEnough(a, b, tolerance float64) bool {
	return a >= b-tolerance && a <= b+tolerance
}
