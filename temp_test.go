package wx

import "testing"

func TestTempUnitString(t *testing.T) {
	tt := []struct {
		name string
		unit TempUnit
		want string
	}{
		{"celsius", Celsius, "°C"},
		{"fahrenheit", Fahrenheit, "°F"},
		{"kelvin", Kelvin, "K"},
		{"rankine", Rankine, "R"},
		{"invalid", TempUnit{-1}, ""},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if tc.unit.String() != tc.want {
				t.Errorf("got %v, want %v", tc.unit.String(), tc.want)
			}
		})
	}
}

func TestNewTemp(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		unit        TempUnit
		valid       bool
	}{
		{"valid celsius", 0, Celsius, true},
		{"valid fahrenheit", 0, Fahrenheit, true},
		{"valid kelvin", 0, Kelvin, true},
		{"valid rankine", 0, Rankine, true},
		{"invalid celsius", -274, Celsius, false},
		{"invalid fahrenheit", -460, Fahrenheit, false},
		{"invalid kelvin", -1, Kelvin, false},
		{"invalid rankine", -1, Rankine, false},
		{"invalid unit", 0, TempUnit{-1}, false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			temp := NewTemp(tc.measurement, tc.unit)
			if temp.valid != tc.valid {
				t.Errorf("got %v, want %v", temp.valid, tc.valid)
			}
		})
	}
}

func TestTempString(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		unit        TempUnit
		want        string
	}{
		{"celsius", 0, Celsius, "0.0°C"},
		{"fahrenheit", 0, Fahrenheit, "0.0°F"},
		{"kelvin", 0, Kelvin, "0.0 K"},
		{"rankine", 0, Rankine, "0.0 R"},
		{"invalid", 0, TempUnit{-1}, "invalid temp unit"},
		{"invalid celsius", -274, Celsius, "invalid °C"},
		{"invalid fahrenheit", -460, Fahrenheit, "invalid °F"},
		{"invalid kelvin", -1, Kelvin, "invalid K"},
		{"invalid rankine", -1, Rankine, "invalid R"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			temp := NewTemp(tc.measurement, tc.unit)
			if temp.String() != tc.want {
				t.Errorf("got %v, want %v", temp.String(), tc.want)
			}
		})
	}
}

func TestTempValid(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		unit        TempUnit
		valid       bool
	}{
		{"valid celsius", freezingC, Celsius, true},
		{"valid fahrenheit", freezingF, Fahrenheit, true},
		{"valid kelvin", absoluteZeroK, Kelvin, true},
		{"valid rankine", absoluteZeroR, Rankine, true},
		{"invalid celsius", -274, Celsius, false},
		{"invalid fahrenheit", -460, Fahrenheit, false},
		{"invalid kelvin", -1, Kelvin, false},
		{"invalid rankine", -1, Rankine, false},
		{"invalid unit", 0, TempUnit{-1}, false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			temp := NewTemp(tc.measurement, tc.unit)
			if temp.Valid() != tc.valid {
				t.Errorf("got %v, want %v", temp.Valid(), tc.valid)
			}
		})
	}
}

func TestTempC(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		unit        TempUnit
		want        float64
	}{
		{"celsius", freezingC, Celsius, freezingC},
		{"fahrenheit", freezingF, Fahrenheit, freezingC},
		{"kelvin", 0, Kelvin, absoluteZeroC},
		{"rankine", 0, Rankine, absoluteZeroC},
		{"invalid", 0, TempUnit{-1}, 0},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			temp := NewTemp(tc.measurement, tc.unit)
			if temp.C() != tc.want {
				t.Errorf("got %v, want %v", temp.C(), tc.want)
			}
		})
	}
}

func TestTempToC(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		unit        TempUnit
		want        float64
	}{
		{"celsius", freezingC, Celsius, freezingC},
		{"fahrenheit", freezingF, Fahrenheit, freezingC},
		{"kelvin", absoluteZeroK, Kelvin, absoluteZeroC},
		{"rankine", absoluteZeroR, Rankine, absoluteZeroC},
		{"invalid", 0, TempUnit{-1}, 0},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			temp := NewTemp(tc.measurement, tc.unit)
			if temp.ToC().C() != tc.want {
				t.Errorf("got %v, want %v", temp.ToC().C(), tc.want)
			}
		})
	}
}

func TestTempF(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		unit        TempUnit
		want        float64
	}{
		{"celsius", freezingC, Celsius, freezingF},
		{"fahrenheit", freezingF, Fahrenheit, freezingF},
		{"kelvin", absoluteZeroK, Kelvin, absoluteZeroF},
		{"rankine", absoluteZeroR, Rankine, absoluteZeroF},
		{"invalid", 0, TempUnit{-1}, 0},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			temp := NewTemp(tc.measurement, tc.unit)
			if temp.F() != tc.want {
				t.Errorf("got %v, want %v", temp.F(), tc.want)
			}
		})
	}
}

func TestTempToF(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		unit        TempUnit
		want        float64
	}{
		{"celsius", freezingC, Celsius, freezingF},
		{"fahrenheit", freezingF, Fahrenheit, freezingF},
		{"kelvin", absoluteZeroK, Kelvin, absoluteZeroF},
		{"rankine", absoluteZeroR, Rankine, absoluteZeroF},
		{"invalid", 0, TempUnit{-1}, 0},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			temp := NewTemp(tc.measurement, tc.unit)
			if temp.ToF().F() != tc.want {
				t.Errorf("got %v, want %v", temp.ToF().F(), tc.want)
			}
		})
	}
}

func TestTempK(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		unit        TempUnit
		want        float64
	}{
		{"celsius", absoluteZeroC, Celsius, absoluteZeroK},
		{"fahrenheit", absoluteZeroF, Fahrenheit, absoluteZeroK},
		{"kelvin", absoluteZeroK, Kelvin, absoluteZeroK},
		{"rankine", absoluteZeroR, Rankine, absoluteZeroK},
		{"invalid", 0, TempUnit{-1}, 0},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			temp := NewTemp(tc.measurement, tc.unit)
			if temp.K() != tc.want {
				t.Errorf("got %v, want %v", temp.K(), tc.want)
			}
		})
	}
}

func TestTempToK(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		unit        TempUnit
		want        float64
	}{
		{"celsius", absoluteZeroC, Celsius, absoluteZeroK},
		{"fahrenheit", absoluteZeroF, Fahrenheit, absoluteZeroK},
		{"kelvin", absoluteZeroK, Kelvin, absoluteZeroK},
		{"rankine", absoluteZeroR, Rankine, absoluteZeroK},
		{"invalid", 0, TempUnit{-1}, 0},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			temp := NewTemp(tc.measurement, tc.unit)
			if temp.ToK().K() != tc.want {
				t.Errorf("got %v, want %v", temp.ToK().K(), tc.want)
			}
		})
	}
}

func TestTempR(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		unit        TempUnit
		want        float64
	}{
		{"celsius", absoluteZeroC, Celsius, absoluteZeroR},
		{"fahrenheit", absoluteZeroF, Fahrenheit, absoluteZeroR},
		{"kelvin", absoluteZeroK, Kelvin, absoluteZeroR},
		{"rankine", absoluteZeroR, Rankine, absoluteZeroR},
		{"invalid", 0, TempUnit{-1}, 0},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			temp := NewTemp(tc.measurement, tc.unit)
			if temp.R() != tc.want {
				t.Errorf("got %v, want %v", temp.R(), tc.want)
			}
		})
	}
}

func TestTempToR(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		unit        TempUnit
		want        float64
	}{
		{"celsius", absoluteZeroC, Celsius, absoluteZeroR},
		{"fahrenheit", absoluteZeroF, Fahrenheit, absoluteZeroR},
		{"kelvin", absoluteZeroK, Kelvin, absoluteZeroR},
		{"rankine", absoluteZeroR, Rankine, absoluteZeroR},
		{"invalid", 0, TempUnit{-1}, 0},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			temp := NewTemp(tc.measurement, tc.unit)
			if temp.ToR().R() != tc.want {
				t.Errorf("got %v, want %v", temp.ToR().R(), tc.want)
			}
		})
	}
}
