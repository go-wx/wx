package wx

import (
	"github.com/go-wx/wx/internal/tests"
	"testing"
)

func TestVelocityFps_ImplementsVelocity(t *testing.T) {
	var v Velocity = &VelocityFps{}
	_ = v
}

func TestVelocityFps_Fps(t *testing.T) {
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
			v := VelocityFps{}
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

func TestVelocityFps_Kts(t *testing.T) {
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
			name:        "100 fps",
			measurement: 100,
			want:        100 * secondsInHour / feetPerNauticalMile,
			valid:       true,
		},
		{
			name:        "-100 fps",
			measurement: -100,
			want:        0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityFps{}
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

func TestVelocityFps_Kph(t *testing.T) {
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
			name:        "100 fps",
			measurement: 100,
			want:        100 * secondsInHour / feetPerMeter / 1000,
			valid:       true,
		},
		{
			name:        "-100 fps",
			measurement: -100,
			want:        0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityFps{}
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

func TestVelocityFps_Mph(t *testing.T) {
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
			name:        "100 fps",
			measurement: 100,
			want:        100 * secondsInHour / feetPerStatuteMile,
			valid:       true,
		},
		{
			name:        "-100 fps",
			measurement: -100,
			want:        0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityFps{}
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

func TestVelocityFps_Mps(t *testing.T) {
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
			name:        "100 fps",
			measurement: 100,
			want:        100 / feetPerMeter,
			valid:       true,
		},
		{
			name:        "-100 fps",
			measurement: -100,
			want:        0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityFps{}
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

func TestVelocityFps_Set(t *testing.T) {
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
			v := VelocityFps{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			if !tests.CloseEnough(v.measurement, tc.want, 1e-6) {
				t.Errorf("got %v, want %v", v.measurement, tc.want)
			}
		})
	}
}

func TestVelocityFps_ToFps(t *testing.T) {
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
			v := VelocityFps{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			v2 := v.ToFps()
			if !tests.CloseEnough(v2.Fps(), tc.want, 1e-6) {
				t.Errorf("got %v, want %v", v2.Fps(), tc.want)
			}
		})
	}
}

func TestVelocityFps_ToKts(t *testing.T) {
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
			name:        "100 fps",
			measurement: 100,
			want:        100 * secondsInHour / feetPerNauticalMile,
			valid:       true,
		},
		{
			name:        "-100 fps",
			measurement: -100,
			want:        0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityFps{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			v2 := v.ToKts()
			if !tests.CloseEnough(v2.Kts(), tc.want, 1e-6) {
				t.Errorf("got %v, want %v", v2.Kts(), tc.want)
			}
		})
	}
}

func TestVelocityFps_ToKph(t *testing.T) {
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
			name:        "100 fps",
			measurement: 100,
			want:        100 * secondsInHour / feetPerMeter / 1000,
			valid:       true,
		},
		{
			name:        "-100 fps",
			measurement: -100,
			want:        0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityFps{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			v2 := v.ToKph()
			if !tests.CloseEnough(v2.Kph(), tc.want, 1e-6) {
				t.Errorf("got %v, want %v", v2.Kph(), tc.want)
			}
		})
	}
}

func TestVelocityFps_ToMph(t *testing.T) {
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
			name:        "100 fps",
			measurement: 100,
			want:        100 * secondsInHour / feetPerStatuteMile,
			valid:       true,
		},
		{
			name:        "-100 fps",
			measurement: -100,
			want:        0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityFps{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			v2 := v.ToMph()
			if !tests.CloseEnough(v2.Mph(), tc.want, 1e-6) {
				t.Errorf("got %v, want %v", v2.Mph(), tc.want)
			}
		})
	}
}

func TestVelocityFps_ToMps(t *testing.T) {
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
			name:        "100 fps",
			measurement: 100,
			want:        100 / feetPerMeter,
			valid:       true,
		},
		{
			name:        "-100 fps",
			measurement: -100,
			want:        0,
			valid:       false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityFps{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			v2 := v.ToMps()
			if !tests.CloseEnough(v2.Mps(), tc.want, 1e-6) {
				t.Errorf("got %v, want %v", v2.Mps(), tc.want)
			}
		})
	}
}

func TestVelocityFps_Units(t *testing.T) {
	v := VelocityFps{}
	if got := v.Units(); got != FeetPerSecond {
		t.Errorf("got %v, want %v", got, FeetPerSecond)
	}
}

func TestVelocityFps_Valid(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        bool
	}{
		{"zero", 0, true},
		{"positive", 1, true},
		{"negative", -1, false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityFps{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.want {
					t.Errorf("unexpected error: %v", err)
				}
			}

			if got := v.Valid(); got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
