package tests

// CloseEnough returns true if a is within tolerance of b.
func CloseEnough(a, b, tolerance float64) bool {
	return a >= b-tolerance && a <= b+tolerance
}
