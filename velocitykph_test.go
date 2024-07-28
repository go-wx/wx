package wx

import (
	"github.com/go-wx/wx/internal/tests"
	"testing"
)

func TestVelocityKph_ImplementsVelocity(t *testing.T) {
	var v Velocity = &VelocityKph{}
	_ = v
}

func TestVelocityKph_Fps(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        float64
		valid       bool
	}{
		{"zero", 0, 0, true},
		{"positive", 1, 1 / secondsInHour * feetPerMeter * 1000, true},
		{"negative", -1, 0, false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityKph{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			if !tests.CloseEnough(v.Fps(), tc.want, 1e-6) {
				t.Errorf("got %v, want %v", v.Fps(), tc.want)
			}
		})
	}
}

func TestVelocityKph_Kts(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        float64
		valid       bool
	}{
		{
			name:        "zero",
			measurement: 0,
			want:        0,
			valid:       true,
		},
		{
			name:        "100 kph",
			measurement: 100,
			want:        100 * 1000 * feetPerMeter / feetPerNauticalMile,
			valid:       true,
		},
		{
			name:        "negative",
			measurement: -1,
			want:        0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityKph{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			if !tests.CloseEnough(v.Kts(), tc.want, 1e-6) {
				t.Errorf("got %v, want %v", v.Kts(), tc.want)
			}
		})
	}
}

func TestVelocityKph_Kph(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        float64
		valid       bool
	}{
		{"zero", 0, 0, true},
		{"positive", 1, 1, true},
		{"negative", -1, 0, false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityKph{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			if !tests.CloseEnough(v.Kph(), tc.want, 1e-6) {
				t.Errorf("got %v, want %v", v.Kph(), tc.want)
			}
		})
	}
}

func TestVelocityKph_Mph(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        float64
		valid       bool
	}{
		{
			name:        "zero",
			measurement: 0,
			want:        0,
			valid:       true,
		},
		{
			name:        "positive",
			measurement: 1,
			want:        1 * 1000 * feetPerMeter / feetPerStatuteMile,
			valid:       true,
		},
		{
			name:        "negative",
			measurement: -1,
			want:        0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityKph{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			if !tests.CloseEnough(v.Mph(), tc.want, 1e-6) {
				t.Errorf("got %v, want %v", v.Mph(), tc.want)
			}
		})
	}
}

func TestVelocityKph_Mps(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        float64
		valid       bool
	}{
		{
			name:        "zero",
			measurement: 0,
			want:        0,
			valid:       true,
		},
		{
			name:        "100 kph",
			measurement: 100,
			want:        100 / secondsInHour * 1000,
			valid:       true,
		},
		{
			name:        "-100 kph",
			measurement: -100,
			want:        0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityKph{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			if !tests.CloseEnough(v.Mps(), tc.want, 1e-6) {
				t.Errorf("got %v, want %v", v.Mps(), tc.want)
			}
		})
	}
}

func TestVelocityKph_Set(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        float64
		valid       bool
	}{
		{"zero", 0, 0, true},
		{"positive", 1, 1, true},
		{"negative", -1, 0, false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityKph{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			if v.valid != tc.valid {
				t.Errorf("got %v, want %v", v.valid, tc.valid)
			}

			if v.measurement != tc.want {
				t.Errorf("got %v, want %v", v.measurement, tc.want)
			}
		})
	}
}

func TestVelocityKph_ToFps(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        float64
		valid       bool
	}{
		{
			name:        "zero",
			measurement: 0,
			want:        0,
			valid:       true,
		},
		{
			name:        "100 kph",
			measurement: 100,
			want:        100 / secondsInHour * 1000 * feetPerMeter,
			valid:       true,
		},
		{
			name:        "-100 kph",
			measurement: -100,
			want:        0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityKph{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			fps := v.ToFps()
			if !tests.CloseEnough(fps.Fps(), tc.want, 1e-6) {
				t.Errorf("got %v, want %v", fps.Fps(), tc.want)
			}

			if fps.valid != tc.valid {
				t.Errorf("got %v, want %v", fps.valid, tc.valid)
			}
		})
	}
}

func TestVelocityKph_ToKts(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        float64
		valid       bool
	}{
		{
			name:        "zero",
			measurement: 0,
			want:        0,
			valid:       true,
		},
		{
			name:        "100 kph",
			measurement: 100,
			want:        100 * 1000 * feetPerMeter / feetPerNauticalMile,
			valid:       true,
		},
		{
			name:        "-100 kph",
			measurement: -100,
			want:        0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityKph{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			kts := v.ToKts()
			if !tests.CloseEnough(kts.Kts(), tc.want, 1e-6) {
				t.Errorf("got %v, want %v", kts.Kts(), tc.want)
			}

			if kts.valid != tc.valid {
				t.Errorf("got %v, want %v", kts.valid, tc.valid)
			}
		})
	}
}

func TestVelocityKph_ToKph(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        float64
		valid       bool
	}{
		{"zero", 0, 0, true},
		{"positive", 1, 1, true},
		{"negative", -1, 0, false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityKph{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			kph := v.ToKph()
			if !tests.CloseEnough(kph.Kph(), tc.want, 1e-6) {
				t.Errorf("got %v, want %v", kph.Kph(), tc.want)
			}

			if kph.valid != tc.valid {
				t.Errorf("got %v, want %v", kph.valid, tc.valid)
			}
		})
	}
}

func TestVelocityKph_ToMph(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        float64
		valid       bool
	}{
		{
			name:        "zero",
			measurement: 0,
			want:        0,
			valid:       true,
		},
		{
			name:        "100 kph",
			measurement: 100,
			want:        100 * 1000 * feetPerMeter / feetPerStatuteMile,
			valid:       true,
		},
		{
			name:        "-100 kph",
			measurement: -100,
			want:        0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityKph{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			mph := v.ToMph()
			if !tests.CloseEnough(mph.Mph(), tc.want, 1e-6) {
				t.Errorf("got %v, want %v", mph.Mph(), tc.want)
			}

			if mph.valid != tc.valid {
				t.Errorf("got %v, want %v", mph.valid, tc.valid)
			}
		})
	}
}

func TestVelocityKph_ToMps(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        float64
		valid       bool
	}{
		{
			name:        "zero",
			measurement: 0,
			want:        0,
			valid:       true,
		},
		{
			name:        "100 kph",
			measurement: 100,
			want:        100 / secondsInHour * 1000,
			valid:       true,
		},
		{
			name:        "-100 kph",
			measurement: -100,
			want:        0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityKph{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			mps := v.ToMps()
			if !tests.CloseEnough(mps.Mps(), tc.want, 1e-6) {
				t.Errorf("got %v, want %v", mps.Mps(), tc.want)
			}

			if mps.valid != tc.valid {
				t.Errorf("got %v, want %v", mps.valid, tc.valid)
			}
		})
	}
}

func TestVelocityKph_Units(t *testing.T) {
	v := VelocityKph{}
	if v.Units() != KilometersPerHour {
		t.Errorf("got %v, want %v", v.Units(), KilometersPerHour)
	}
}

func TestVelocityKph_Valid(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		valid       bool
	}{
		{"zero", 0, true},
		{"positive", 1, true},
		{"negative", -1, false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityKph{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			if v.Valid() != tc.valid {
				t.Errorf("got %v, want %v", v.Valid(), tc.valid)
			}
		})
	}
}
