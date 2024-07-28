package wx

import (
	"github.com/go-wx/wx/internal/tests"
	"testing"
)

func TestVelocityMph_ImplementsVelocity(t *testing.T) {
	var v Velocity = &VelocityMph{}
	_ = v
}

func TestVelocityMph_Fps(t *testing.T) {
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
			want:        1 / secondsInHour * feetPerStatuteMile,
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
			v := VelocityMph{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			if !tests.CloseEnough(v.Fps(), tc.want, 1e-6) {
				t.Errorf("got %v, want %v", v.Fps(), tc.want)
			}

			if v.Valid() != tc.valid {
				t.Errorf("got %v, want %v", v.Valid(), tc.valid)
			}
		})
	}
}

func TestVelocityMph_Kts(t *testing.T) {
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
			name:        "100 mph",
			measurement: 100,
			want:        100 * feetPerStatuteMile / feetPerNauticalMile,
			valid:       true,
		},
		{
			name:        "-100 mph",
			measurement: -100,
			want:        0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityMph{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			if !tests.CloseEnough(v.Kts(), tc.want, 1e-6) {
				t.Errorf("got %v, want %v", v.Kts(), tc.want)
			}

			if v.Valid() != tc.valid {
				t.Errorf("got %v, want %v", v.Valid(), tc.valid)
			}
		})
	}
}

func TestVelocityMph_Kph(t *testing.T) {
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
			name:        "100 mph",
			measurement: 100,
			want:        100 * feetPerStatuteMile / feetPerMeter / 1000,
			valid:       true,
		},
		{
			name:        "-100 mph",
			measurement: -100,
			want:        0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityMph{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			if !tests.CloseEnough(v.Kph(), tc.want, 1e-6) {
				t.Errorf("got %v, want %v", v.Kph(), tc.want)
			}

			if v.Valid() != tc.valid {
				t.Errorf("got %v, want %v", v.Valid(), tc.valid)
			}
		})
	}
}

func TestVelocityMph_Mph(t *testing.T) {
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
			name:        "100 mph",
			measurement: 100,
			want:        100,
			valid:       true,
		},
		{
			name:        "-100 mph",
			measurement: -100,
			want:        0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityMph{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			if !tests.CloseEnough(v.Mph(), tc.want, 1e-6) {
				t.Errorf("got %v, want %v", v.Mph(), tc.want)
			}

			if v.Valid() != tc.valid {
				t.Errorf("got %v, want %v", v.Valid(), tc.valid)
			}
		})
	}
}

func TestVelocityMph_Mps(t *testing.T) {
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
			name:        "100 mph",
			measurement: 100,
			want:        100 * feetPerStatuteMile / feetPerMeter / secondsInHour,
			valid:       true,
		},
		{
			name:        "-100 mph",
			measurement: -100,
			want:        0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityMph{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			if !tests.CloseEnough(v.Mps(), tc.want, 1e-6) {
				t.Errorf("got %v, want %v", v.Mps(), tc.want)
			}

			if v.Valid() != tc.valid {
				t.Errorf("got %v, want %v", v.Valid(), tc.valid)
			}
		})
	}
}

func TestVelocityMph_Set(t *testing.T) {
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
			want:        1,
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
			v := VelocityMph{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			if !tests.CloseEnough(v.Mph(), tc.want, 1e-6) {
				t.Errorf("got %v, want %v", v.Mph(), tc.want)
			}

			if v.Valid() != tc.valid {
				t.Errorf("got %v, want %v", v.Valid(), tc.valid)
			}
		})
	}
}

func TestVelocityMph_ToFps(t *testing.T) {
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
			want:        1 / secondsInHour * feetPerStatuteMile,
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
			v := VelocityMph{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			fps := v.ToFps()
			if !tests.CloseEnough(fps.Fps(), tc.want, 1e-6) {
				t.Errorf("got %v, want %v", fps.Fps(), tc.want)
			}

			if fps.Valid() != tc.valid {
				t.Errorf("got %v, want %v", fps.Valid(), tc.valid)
			}
		})
	}
}

func TestVelocityMph_ToKts(t *testing.T) {
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
			name:        "100 mph",
			measurement: 100,
			want:        100 * feetPerStatuteMile / feetPerNauticalMile,
			valid:       true,
		},
		{
			name:        "-100 mph",
			measurement: -100,
			want:        0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityMph{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			kts := v.ToKts()
			if !tests.CloseEnough(kts.Kts(), tc.want, 1e-6) {
				t.Errorf("got %v, want %v", kts.Kts(), tc.want)
			}

			if kts.Valid() != tc.valid {
				t.Errorf("got %v, want %v", kts.Valid(), tc.valid)
			}
		})
	}
}

func TestVelocityMph_ToKph(t *testing.T) {
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
			name:        "100 mph",
			measurement: 100,
			want:        100 * feetPerStatuteMile / feetPerMeter / 1000,
			valid:       true,
		},
		{
			name:        "-100 mph",
			measurement: -100,
			want:        0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityMph{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			kph := v.ToKph()
			if !tests.CloseEnough(kph.Kph(), tc.want, 1e-6) {
				t.Errorf("got %v, want %v", kph.Kph(), tc.want)
			}

			if kph.Valid() != tc.valid {
				t.Errorf("got %v, want %v", kph.Valid(), tc.valid)
			}
		})
	}
}

func TestVelocityMph_ToMph(t *testing.T) {
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
			want:        1,
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
			v := VelocityMph{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			mph := v.ToMph()
			if !tests.CloseEnough(mph.Mph(), tc.want, 1e-6) {
				t.Errorf("got %v, want %v", mph.Mph(), tc.want)
			}

			if mph.Valid() != tc.valid {
				t.Errorf("got %v, want %v", mph.Valid(), tc.valid)
			}
		})
	}
}

func TestVelocityMph_ToMps(t *testing.T) {
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
			name:        "100 mph",
			measurement: 100,
			want:        100 * feetPerStatuteMile / feetPerMeter / secondsInHour,
			valid:       true,
		},
		{
			name:        "-100 mph",
			measurement: -100,
			want:        0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityMph{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			mps := v.ToMps()
			if !tests.CloseEnough(mps.Mps(), tc.want, 1e-6) {
				t.Errorf("got %v, want %v", mps.Mps(), tc.want)
			}

			if mps.Valid() != tc.valid {
				t.Errorf("got %v, want %v", mps.Valid(), tc.valid)
			}
		})
	}
}

func TestVelocityMph_Units(t *testing.T) {
	v := VelocityMph{}
	units := v.Units()
	if units != MilesPerHour {
		t.Errorf("got %v, want mph", units)
	}
}

func TestVelocityMph_Valid(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		valid       bool
	}{
		{
			name:        "zero",
			measurement: 0,
			valid:       true,
		},
		{
			name:        "positive",
			measurement: 1,
			valid:       true,
		},
		{
			name:        "negative",
			measurement: -1,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityMph{}
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
