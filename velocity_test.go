package wx

import "testing"

func TestVelocityUnit_String(t *testing.T) {
	tt := []struct {
		name     string
		unit     VelocityUnit
		expected string
	}{
		{"FeetPerSecond", FeetPerSecond, "fps"},
		{"Knots", Knots, "kts"},
		{"KilometersPerHour", KilometersPerHour, "kph"},
		{"MilesPerHour", MilesPerHour, "mph"},
		{"MetersPerSecond", MetersPerSecond, "mps"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if actual := tc.unit.String(); actual != tc.expected {
				t.Errorf("expected %s; got %s", tc.expected, actual)
			}
		})
	}
}
