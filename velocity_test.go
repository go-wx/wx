package wx

import (
	"math"
	"testing"
)

func TestVelocityType_String(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name string
		unit velocityType
		want string
	}{
		{"fps", fps, "fps"},
		{"kts", kts, "kts"},
		{"kph", kph, "kph"},
		{"mph", mph, "mph"},
		{"mps", mps, "mps"},
		{"invalid", velocityType(0), ""},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if tc.unit.String() != tc.want {
				t.Errorf("got %v, want %v", tc.unit.String(), tc.want)
			}
		})
	}
}

func TestVelocityUnit_String(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name string
		unit VelocityUnit
		want string
	}{
		{"fps", Fps, "fps"},
		{"kts", Kts, "kts"},
		{"kph", Kph, "kph"},
		{"mph", Mph, "mph"},
		{"mps", Mps, "mps"},
		{"invalid", VelocityUnit{velocityType(0)}, ""},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if tc.unit.String() != tc.want {
				t.Errorf("got %v, want %v", tc.unit.String(), tc.want)
			}
		})
	}
}

func TestNewVelocity(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name        string
		measurement float64
		unit        VelocityUnit
		expected    Velocity
	}{
		{"fps", 1, Fps,
			Velocity{measurement: 1, unit: Fps, valid: true}},
		{"kts", 1, Kts,
			Velocity{measurement: 1, unit: Kts, valid: true}},
		{"kph", 1, Kph,
			Velocity{measurement: 1, unit: Kph, valid: true}},
		{"mph", 1, Mph,
			Velocity{measurement: 1, unit: Mph, valid: true}},
		{"mps", 1, Mps,
			Velocity{measurement: 1, unit: Mps, valid: true}},
		{"invalid fps", -1, Fps,
			Velocity{measurement: -1, unit: Fps, valid: false}},
		{"invalid", -1,
			VelocityUnit{velocityType(0)}, Velocity{valid: false}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			velocity := NewVelocity(tc.measurement, tc.unit)
			if velocity != tc.expected {
				t.Errorf("got %v, want %v", velocity, tc.expected)
			}
		})
	}
}

func TestVelocity_Fps(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name        string
		measurement float64
		unit        VelocityUnit
		want        float64
	}{
		{"fps", 1, Fps, 1},
		{"kts", 1, Kts, feetPerNauticalMile / 3600},
		{"kph", 1, Kph, feetPerMeter * 1000 / 3600},
		{"mph", 1, Mph, feetPerStatuteMile / 3600},
		{"mps", 1, Mps, feetPerMeter},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			velocity := NewVelocity(tc.measurement, tc.unit)
			if got := velocity.Fps(); got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestVelocity_Kph(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name        string
		measurement float64
		unit        VelocityUnit
		want        float64
	}{
		{"fps", 1, Fps, feetPerMeter * 1000 / 3600},
		{"kts", 1, Kts, feetPerNauticalMile / feetPerMeter / 1000},
		{"kph", 1, Kph, 1},
		{"mph", 1, Mph, feetPerStatuteMile / 3600 / feetPerMeter * 1000},
		{"mps", 1, Mps, 3600.0 / 1000.0},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			velocity := NewVelocity(tc.measurement, tc.unit)
			if got := velocity.Kph(); got != tc.want {
				// Check for floating point error.
				if math.Abs(got-tc.want) > 0.0000001 {
					t.Errorf("got %v, want %v", got, tc.want)
				}
			}
		})
	}
}

func TestVelocity_Kts(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name        string
		measurement float64
		unit        VelocityUnit
		want        float64
	}{
		{"fps", 1, Fps, 3600 / feetPerNauticalMile},
		{"kts", 1, Kts, 1},
		{"kph", 1, Kph, feetPerMeter * 1000 / feetPerNauticalMile},
		{"mph", 1, Mph, feetPerStatuteMile / feetPerNauticalMile},
		{"mps", 1, Mps, feetPerMeter / feetPerNauticalMile},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			velocity := NewVelocity(tc.measurement, tc.unit)
			if got := velocity.Kts(); got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestVelocity_Mph(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name        string
		measurement float64
		unit        VelocityUnit
		want        float64
	}{
		{"fps", 1, Fps, 3600 / feetPerStatuteMile},
		{"kts", 1, Kts, feetPerNauticalMile / feetPerStatuteMile},
		{"kph", 1, Kph, feetPerMeter * 1000 / feetPerStatuteMile},
		{"mph", 1, Mph, 1},
		{"mps", 1, Mps, feetPerMeter * 3600 / feetPerStatuteMile},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			velocity := NewVelocity(tc.measurement, tc.unit)
			if got := velocity.Mph(); got != tc.want {
				// Check for floating point error.
				if math.Abs(got-tc.want) > 0.0000001 {
					t.Errorf("got %v, want %v", got, tc.want)
				}
			}
		})
	}
}

func TestVelocity_Mps(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name        string
		measurement float64
		unit        VelocityUnit
		want        float64
	}{
		{"fps", 1, Fps, 1 / feetPerMeter},
		{"kts", 1, Kts, feetPerNauticalMile / feetPerMeter / 3600},
		{"kph", 1, Kph, 1000.0 / 3600.0},
		{"mph", 1, Mph, feetPerStatuteMile / 3600.0 / feetPerMeter},
		{"mps", 1, Mps, 1},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			velocity := NewVelocity(tc.measurement, tc.unit)
			if got := velocity.Mps(); got != tc.want {
				// Check for floating point error.
				if math.Abs(got-tc.want) > 0.0000001 {
					t.Errorf("got %v, want %v", got, tc.want)
				}
			}
		})
	}
}
