package wx

import "testing"

func TestDistanceType_String(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		distance distanceType
		expected string
	}{
		{"feet", feet, "ft"},
		{"kilometers", kilometers, "km"},
		{"nautical miles", nauticalMiles, "NM"},
		{"meters", meters, "m"},
		{"statute miles", statuteMiles, "SM"},
		{"parsec", parsec, "pc"},
		{"random", 99, ""},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.distance.String() != test.expected {
				t.Errorf("expected %s, got %s", test.expected, test.distance.String())
			}
		})
	}
}

func TestNewDistance(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		measurement float64
		unit        DistanceUnit
		expected    Distance
	}{
		{
			name:        "valid feet",
			measurement: 1.0,
			unit:        Feet,
			expected: Distance{
				measurement: 1.0,
				unit:        DistanceUnit{distanceType: feet},
				valid:       true,
			},
		},
		{
			name:        "valid kilometers",
			measurement: 1.0,
			unit:        Kilometers,
			expected: Distance{
				measurement: 1.0,
				unit:        DistanceUnit{distanceType: kilometers},
				valid:       true,
			},
		},
		{
			name:        "valid nautical miles",
			measurement: 1.0,
			unit:        NauticalMiles,
			expected: Distance{
				measurement: 1.0,
				unit:        DistanceUnit{distanceType: nauticalMiles},
				valid:       true,
			},
		},
		{
			name:        "valid meters",
			measurement: 1.0,
			unit:        Meters,
			expected: Distance{
				measurement: 1.0,
				unit:        DistanceUnit{distanceType: meters},
				valid:       true,
			},
		},
		{
			name:        "valid statute miles",
			measurement: 1.0,
			unit:        StatuteMiles,
			expected: Distance{
				measurement: 1.0,
				unit:        DistanceUnit{distanceType: statuteMiles},
				valid:       true,
			},
		},
		{
			name:        "valid parsec",
			measurement: 1.0,
			unit:        Parsec,
			expected: Distance{
				measurement: 1.0,
				unit:        DistanceUnit{distanceType: parsec},
				valid:       true,
			},
		},
		{
			name:        "invalid feet",
			measurement: -1.0,
			unit:        Feet,
			expected: Distance{
				measurement: -1.0,
				unit:        DistanceUnit{distanceType: feet},
				valid:       false,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			d := NewDistance(test.measurement, test.unit)
			if d != test.expected {
				t.Errorf("expected %v, got %v", test.expected, d)
			}
		})
	}
}

func TestDistance_String(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		measurement float64
		unit        DistanceUnit
		expected    string
	}{
		{"valid feet", 1.0, Feet, "1.00 ft"},
		{"valid kilometers", 1.0, Kilometers, "1.00 km"},
		{"valid nautical miles", 1.0, NauticalMiles, "1.00 NM"},
		{"valid meters", 1.0, Meters, "1.00 m"},
		{"valid statute miles", 1.0, StatuteMiles, "1.00 SM"},
		{"valid parsec", 1.0, Parsec, "1.00 pc"},
		{"invalid feet", -1.0, Feet, "invalid ft"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			d := NewDistance(test.measurement, test.unit)
			if d.String() != test.expected {
				t.Errorf("expected %s, got %s", test.expected, d.String())
			}
		})
	}
}

func TestDistance_Valid(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		measurement float64
		unit        DistanceUnit
		expected    bool
	}{
		{"valid feet", 1.0, Feet, true},
		{"valid kilometers", 1.0, Kilometers, true},
		{"valid nautical miles", 1.0, NauticalMiles, true},
		{"valid meters", 1.0, Meters, true},
		{"valid statute miles", 1.0, StatuteMiles, true},
		{"valid parsec", 1.0, Parsec, true},
		{"invalid", -1.0, Feet, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			d := NewDistance(test.measurement, test.unit)
			if d.Valid() != test.expected {
				t.Errorf("expected %t, got %t", test.expected, d.Valid())
			}
		})
	}
}

func TestDistanceFeet(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		measurement float64
		unit        DistanceUnit
		expected    float64
	}{
		{"valid feet", 1.0, Feet, 1.0},
		{"valid kilometers", 1.0, Kilometers, feetPerMeter * 1000},
		{"valid nautical miles", 1.0, NauticalMiles, feetPerNauticalMile},
		{"valid meters", 1.0, Meters, feetPerMeter},
		{"valid statute miles", 1.0, StatuteMiles, feetPerStatuteMile},
		{"valid parsec", 1.0, Parsec, metersPerParsec * feetPerMeter},
		{"invalid", -1.0, DistanceUnit{99}, 0.0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			d := NewDistance(test.measurement, test.unit)
			if d.FT() != test.expected {
				// Check if the expected value is within 0.0001 of the actual value
				if d.FT()-test.expected > 0.000001 {
					t.Errorf("expected %f, got %f", test.expected, d.FT())
				}
			}
		})
	}
}

func TestDistanceKilometers(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		measurement float64
		unit        DistanceUnit
		expected    float64
	}{
		{"valid feet", 1.0, Feet, 1.0 / feetPerMeter / 1000},
		{"valid kilometers", 1.0, Kilometers, 1.0},
		{"valid nautical miles", 1.0, NauticalMiles,
			feetPerNauticalMile / feetPerMeter / 1000},
		{"valid meters", 1.0, Meters, 1.0 / 1000},
		{"valid statute miles", 1.0, StatuteMiles,
			feetPerStatuteMile / feetPerMeter / 1000},
		{"valid parsec", 1.0, Parsec, metersPerParsec / 1000},
		{"invalid", -1.0, DistanceUnit{99}, 0.0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			d := NewDistance(test.measurement, test.unit)
			if d.KM() != test.expected {
				// Check if the expected value is within 0.0001 of the actual value
				if d.KM()-test.expected > 0.000001 {
					t.Errorf("expected %f, got %f", test.expected, d.KM())
				}
			}
		})
	}
}

func TestDistanceNauticalMiles(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		measurement float64
		unit        DistanceUnit
		expected    float64
	}{
		{"valid feet", 1.0, Feet, 1.0 / feetPerNauticalMile},
		{"valid kilometers", 1.0, Kilometers, 1.0 / (feetPerNauticalMile / feetPerMeter / 1000)},
		{"valid nautical miles", 1.0, NauticalMiles, 1.0},
		{"valid meters", 1.0, Meters, 1.0 / (feetPerMeter / 1000)},
		{"valid statute miles", 1.0, StatuteMiles, feetPerStatuteMile / feetPerNauticalMile},
		{"valid parsec", 1.0, Parsec, metersPerParsec / (feetPerMeter / 1000)},
		{"invalid", -1.0, DistanceUnit{99}, 0.0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			d := NewDistance(test.measurement, test.unit)
			if d.NM() != test.expected {
				// Check if the expected value is within 0.0001 of the actual value
				if d.NM()-test.expected > 0.000001 {
					t.Errorf("expected %f, got %f", test.expected, d.NM())
				}
			}
		})
	}
}

func TestDistanceMeters(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		measurement float64
		unit        DistanceUnit
		expected    float64
	}{
		{"valid feet", 1.0, Feet, 1.0 / feetPerMeter},
		{"valid kilometers", 1.0, Kilometers, 1.0 * 1000},
		{"valid nautical miles", 1.0, NauticalMiles, feetPerNauticalMile / feetPerMeter},
		{"valid meters", 1.0, Meters, 1.0},
		{"valid statute miles", 1.0, StatuteMiles, feetPerStatuteMile / feetPerMeter},
		{"valid parsec", 1.0, Parsec, metersPerParsec},
		{"invalid", -1.0, DistanceUnit{99}, 0.0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			d := NewDistance(test.measurement, test.unit)
			if d.M() != test.expected {
				// Check if the expected value is within 0.0001 of the actual value
				if d.M()-test.expected > 0.000001 {
					t.Errorf("expected %f, got %f", test.expected, d.M())
				}
			}
		})
	}
}

func TestDistanceStatuteMiles(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		measurement float64
		unit        DistanceUnit
		expected    float64
	}{
		{"valid feet", 1.0, Feet, 1.0 / feetPerStatuteMile},
		{"valid kilometers", 1.0, Kilometers, 1.0 / (feetPerStatuteMile / feetPerMeter / 1000)},
		{"valid nautical miles", 1.0, NauticalMiles, feetPerNauticalMile / feetPerStatuteMile},
		{"valid meters", 1.0, Meters, 1.0 / (feetPerMeter / 1000)},
		{"valid statute miles", 1.0, StatuteMiles, 1.0},
		{"valid parsec", 1.0, Parsec, metersPerParsec / (feetPerMeter / 1000)},
		{"invalid", -1.0, DistanceUnit{99}, 0.0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			d := NewDistance(test.measurement, test.unit)
			if d.SM() != test.expected {
				// Check if the expected value is within 0.0001 of the actual value
				if d.SM()-test.expected > 0.000001 {
					t.Errorf("expected %f, got %f", test.expected, d.SM())
				}
			}
		})
	}
}

func TestDistanceParsec(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		measurement float64
		unit        DistanceUnit
		expected    float64
	}{
		{"valid feet", 1.0, Feet, 1.0 / metersPerParsec / feetPerMeter},
		{"valid kilometers", 1.0, Kilometers, 1.0 / metersPerParsec * 1000},
		{"valid nautical miles", 1.0, NauticalMiles, feetPerNauticalMile / metersPerParsec},
		{"valid meters", 1.0, Meters, 1.0 / metersPerParsec},
		{"valid statute miles", 1.0, StatuteMiles, feetPerStatuteMile / metersPerParsec},
		{"valid parsec", 1.0, Parsec, 1.0},
		{"invalid", -1.0, DistanceUnit{99}, 0.0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			d := NewDistance(test.measurement, test.unit)
			if d.Parsec() != test.expected {
				// Check if the expected value is within 0.0001 of the actual value
				if d.Parsec()-test.expected > 0.000001 {
					t.Errorf("expected %f, got %f", test.expected, d.Parsec())
				}
			}
		})
	}
}

func TestDistanceAbs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		measurement float64
		unit        DistanceUnit
		expected    float64
	}{
		{"valid feet", 1.0, Feet, 1.0},
		{"valid kilometers", 1.0, Kilometers, 1.0},
		{"valid nautical miles", 1.0, NauticalMiles, 1.0},
		{"valid meters", 1.0, Meters, 1.0},
		{"valid statute miles", 1.0, StatuteMiles, 1.0},
		{"valid parsec", 1.0, Parsec, 1.0},
		{"invalid", -1.0, Feet, 1.0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			d := NewDistance(test.measurement, test.unit)
			if d.Abs().measurement != test.expected {
				t.Errorf("expected %f, got %f", test.expected, d.Abs().measurement)
			}
		})
	}
}

func TestDistanceAdd(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		d1       Distance
		d2       Distance
		expected Distance
	}{
		{
			name:     "adding feet",
			d1:       NewDistance(1.0, Feet),
			d2:       NewDistance(1.0, Feet),
			expected: NewDistance(2.0, Feet),
		},
		{
			name:     "adding kilometers",
			d1:       NewDistance(1.0, Kilometers),
			d2:       NewDistance(1.0, Kilometers),
			expected: NewDistance(2.0, Kilometers),
		},
		{
			name:     "adding nautical miles",
			d1:       NewDistance(1.0, NauticalMiles),
			d2:       NewDistance(1.0, NauticalMiles),
			expected: NewDistance(2.0, NauticalMiles),
		},
		{
			name:     "adding meters",
			d1:       NewDistance(1.0, Meters),
			d2:       NewDistance(1.0, Meters),
			expected: NewDistance(2.0, Meters),
		},
		{
			name:     "adding statute miles",
			d1:       NewDistance(1.0, StatuteMiles),
			d2:       NewDistance(1.0, StatuteMiles),
			expected: NewDistance(2.0, StatuteMiles),
		},
		{
			name:     "adding parsec",
			d1:       NewDistance(1.0, Parsec),
			d2:       NewDistance(1.0, Parsec),
			expected: NewDistance(2.0, Parsec),
		},
		{
			name:     "adding feet and kilometers",
			d1:       NewDistance(1.0, Feet),
			d2:       NewDistance(1.0, Kilometers),
			expected: NewDistance(1.0+1.0*1000*feetPerMeter, Feet),
		},
		{
			name:     "adding feet and nautical miles",
			d1:       NewDistance(1.0, Feet),
			d2:       NewDistance(1.0, NauticalMiles),
			expected: NewDistance(1.0+1.0*feetPerNauticalMile, Feet),
		},
		{
			name: "adding feet and a negative value",
			d1:   NewDistance(1.0, Feet),
			d2:   NewDistance(-1.0, Feet),
			expected: Distance{
				measurement: 0.0,
				unit:        DistanceUnit{distanceType: feet},
				valid:       true,
			},
		},
		{
			name: "adding feet and a larger magnitude negative value",
			d1:   NewDistance(1.0, Feet),
			d2:   NewDistance(-2.0, Feet),
			expected: Distance{
				measurement: -1.0,
				unit:        DistanceUnit{distanceType: feet},
				valid:       false,
			},
		},
		{
			name:     "adding invalid unit to feet",
			d1:       NewDistance(1.0, DistanceUnit{99}),
			d2:       NewDistance(1.0, Feet),
			expected: Distance{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			d := test.d1.Add(test.d2)
			if d != test.expected {
				t.Errorf("expected %v, got %v", test.expected, d)
			}
		})
	}
}

func TestDistanceSub(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		d1       Distance
		d2       Distance
		expected Distance
	}{
		{
			name:     "subtracting feet",
			d1:       NewDistance(1.0, Feet),
			d2:       NewDistance(1.0, Feet),
			expected: NewDistance(0.0, Feet),
		},
		{
			name:     "subtracting kilometers",
			d1:       NewDistance(1.0, Kilometers),
			d2:       NewDistance(1.0, Kilometers),
			expected: NewDistance(0.0, Kilometers),
		},
		{
			name:     "subtracting nautical miles",
			d1:       NewDistance(1.0, NauticalMiles),
			d2:       NewDistance(1.0, NauticalMiles),
			expected: NewDistance(0.0, NauticalMiles),
		},
		{
			name:     "subtracting meters",
			d1:       NewDistance(1.0, Meters),
			d2:       NewDistance(1.0, Meters),
			expected: NewDistance(0.0, Meters),
		},
		{
			name:     "subtracting statute miles",
			d1:       NewDistance(1.0, StatuteMiles),
			d2:       NewDistance(1.0, StatuteMiles),
			expected: NewDistance(0.0, StatuteMiles),
		},
		{
			name:     "subtracting parsec",
			d1:       NewDistance(1.0, Parsec),
			d2:       NewDistance(1.0, Parsec),
			expected: NewDistance(0.0, Parsec),
		},
		{
			name:     "subtract a foot from a kilometer",
			d1:       NewDistance(1.0, Kilometers),
			d2:       NewDistance(1.0, Feet),
			expected: NewDistance(1.0-1.0/feetPerMeter/1000, Kilometers),
		},
		{
			name: "subtract a nautical mile from a foot",
			d1:   NewDistance(1.0, Feet),
			d2:   NewDistance(1.0, NauticalMiles),
			expected: Distance{
				measurement: 1.0 - 1.0*feetPerNauticalMile,
				unit:        DistanceUnit{distanceType: feet},
				valid:       false,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			d := test.d1.Sub(test.d2)
			if d != test.expected {
				t.Errorf("expected %v, got %v", test.expected, d)
			}
		})
	}
}
