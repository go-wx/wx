package wx

import (
	"github.com/go-wx/wx/internal/tests"
	"testing"
)

func TestVelocityKts_ImplementsVelocity(t *testing.T) {
	var v Velocity = &VelocityKts{}
	_ = v
}

func TestVelocityKts_Fps(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        float64
		valid       bool
	}{
		{
			name:        "valid",
			measurement: 1.0,
			want:        feetPerNauticalMile / secondsInHour,
			valid:       true,
		},
		{
			name:        "invalid",
			measurement: -1.0,
			want:        0.0,
			valid:       false,
		},
		{
			name:        "zero",
			measurement: 0.0,
			want:        0.0,
			valid:       true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityKts{}
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

func TestVelocityKts_Kts(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        float64
		valid       bool
	}{
		{
			name:        "valid",
			measurement: 1.0,
			want:        1.0,
			valid:       true,
		},
		{
			name:        "invalid",
			measurement: -1.0,
			want:        0.0,
			valid:       false,
		},
		{
			name:        "zero",
			measurement: 0.0,
			want:        0.0,
			valid:       true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityKts{}
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

func TestVelocityKts_Kph(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        float64
		valid       bool
	}{
		{
			name:        "valid",
			measurement: 1.0,
			want:        feetPerNauticalMile / feetPerMeter / 1000,
			valid:       true,
		},
		{
			name:        "invalid",
			measurement: -1.0,
			want:        0.0,
			valid:       false,
		},
		{
			name:        "zero",
			measurement: 0.0,
			want:        0.0,
			valid:       true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityKts{}
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

func TestVelocityKts_Mph(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        float64
		valid       bool
	}{
		{
			name:        "valid",
			measurement: 1.0,
			want:        feetPerNauticalMile / feetPerStatuteMile,
			valid:       true,
		},
		{
			name:        "invalid",
			measurement: -1.0,
			want:        0.0,
			valid:       false,
		},
		{
			name:        "zero",
			measurement: 0.0,
			want:        0.0,
			valid:       true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityKts{}
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

func TestVelocityKts_Mps(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        float64
		valid       bool
	}{
		{
			name:        "valid",
			measurement: 1.0,
			want:        feetPerNauticalMile / secondsInHour / feetPerMeter,
			valid:       true,
		},
		{
			name:        "invalid",
			measurement: -1.0,
			want:        0.0,
			valid:       false,
		},
		{
			name:        "zero",
			measurement: 0.0,
			want:        0.0,
			valid:       true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityKts{}
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

func TestVelocityKts_Set(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        float64
		valid       bool
	}{
		{
			name:        "valid",
			measurement: 1.0,
			want:        1.0,
			valid:       true,
		},
		{
			name:        "invalid",
			measurement: -1.0,
			want:        0.0,
			valid:       false,
		},
		{
			name:        "zero",
			measurement: 0.0,
			want:        0.0,
			valid:       true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityKts{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			if v.Kts() != tc.want {
				t.Errorf("got %v, want %v", v.Kts(), tc.want)
			}

			if v.Valid() != tc.valid {
				t.Errorf("got %v, want %v", v.Valid(), tc.valid)
			}
		})
	}
}

func TestVelocityKts_ToFps(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        float64
		valid       bool
	}{
		{
			name:        "valid",
			measurement: 1.0,
			want:        feetPerNauticalMile / secondsInHour,
			valid:       true,
		},
		{
			name:        "invalid",
			measurement: -1.0,
			want:        0.0,
			valid:       false,
		},
		{
			name:        "zero",
			measurement: 0.0,
			want:        0.0,
			valid:       true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityKts{}
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

func TestVelocityKts_ToKts(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        float64
		valid       bool
	}{
		{
			name:        "valid",
			measurement: 1.0,
			want:        1.0,
			valid:       true,
		},
		{
			name:        "invalid",
			measurement: -1.0,
			want:        0.0,
			valid:       false,
		},
		{
			name:        "zero",
			measurement: 0.0,
			want:        0.0,
			valid:       true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityKts{}
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

func TestVelocityKts_ToKph(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        float64
		valid       bool
	}{
		{
			name:        "valid",
			measurement: 1.0,
			want:        feetPerNauticalMile / feetPerMeter / 1000,
			valid:       true,
		},
		{
			name:        "invalid",
			measurement: -1.0,
			want:        0.0,
			valid:       false,
		},
		{
			name:        "zero",
			measurement: 0.0,
			want:        0.0,
			valid:       true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityKts{}
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

func TestVelocityKts_ToMph(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        float64
		valid       bool
	}{
		{
			name:        "valid",
			measurement: 1.0,
			want:        feetPerNauticalMile / feetPerStatuteMile,
			valid:       true,
		},
		{
			name:        "invalid",
			measurement: -1.0,
			want:        0.0,
			valid:       false,
		},
		{
			name:        "zero",
			measurement: 0.0,
			want:        0.0,
			valid:       true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityKts{}
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

func TestVelocityKts_ToMps(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        float64
		valid       bool
	}{
		{
			name:        "valid",
			measurement: 1.0,
			want:        feetPerNauticalMile / feetPerMeter / secondsInHour,
			valid:       true,
		},
		{
			name:        "invalid",
			measurement: -1.0,
			want:        0.0,
			valid:       false,
		},
		{
			name:        "zero",
			measurement: 0.0,
			want:        0.0,
			valid:       true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityKts{}
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

func TestVelocityKts_Units(t *testing.T) {
	v := VelocityKts{}
	if v.Units() != Knots {
		t.Errorf("got %v, want %v", v.Units(), Knots)
	}
}

func TestVelocityKts_Valid(t *testing.T) {
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
			v := VelocityKts{}
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
