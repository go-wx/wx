package wx

import (
	"github.com/go-wx/wx/internal/tests"
	"testing"
)

func TestVelocityMps_ImplementVelocity(t *testing.T) {
	var v Velocity = &VelocityMps{}
	_ = v
}

func TestVelocityMps_Fps(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        float64
		valid       bool
	}{
		{
			name:        "positive",
			measurement: 1.0,
			want:        feetPerMeter,
			valid:       true,
		},
		{
			name:        "negative",
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
			v := VelocityMps{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			got := v.Fps()
			if !tests.CloseEnough(got, tc.want, 1e-6) {
				t.Errorf("got %v, want %v", got, tc.want)
			}

			if v.Valid() != tc.valid {
				t.Errorf("got %v, want %v", v.Valid(), tc.valid)
			}
		})
	}
}

func TestVelocityMps_Kts(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        float64
		valid       bool
	}{
		{
			name:        "positive",
			measurement: 1.0,
			want:        feetPerMeter * secondsInHour / feetPerNauticalMile,
			valid:       true,
		},
		{
			name:        "negative",
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
			v := VelocityMps{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			got := v.Kts()
			if !tests.CloseEnough(got, tc.want, 1e-6) {
				t.Errorf("got %v, want %v", got, tc.want)
			}

			if v.Valid() != tc.valid {
				t.Errorf("got %v, want %v", v.Valid(), tc.valid)
			}
		})
	}
}

func TestVelocityMps_Kph(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        float64
		valid       bool
	}{
		{
			name:        "positive",
			measurement: 1.0,
			want:        secondsInHour / 1000,
			valid:       true,
		},
		{
			name:        "negative",
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
			v := VelocityMps{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			got := v.Kph()
			if !tests.CloseEnough(got, tc.want, 1e-6) {
				t.Errorf("got %v, want %v", got, tc.want)
			}

			if v.Valid() != tc.valid {
				t.Errorf("got %v, want %v", v.Valid(), tc.valid)
			}
		})
	}
}

func TestVelocityMps_Mph(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        float64
		valid       bool
	}{
		{
			name:        "positive",
			measurement: 1.0,
			want:        feetPerMeter / feetPerStatuteMile * secondsInHour,
			valid:       true,
		},
		{
			name:        "negative",
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
			v := VelocityMps{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			got := v.Mph()
			if !tests.CloseEnough(got, tc.want, 1e-6) {
				t.Errorf("got %v, want %v", got, tc.want)
			}

			if v.Valid() != tc.valid {
				t.Errorf("got %v, want %v", v.Valid(), tc.valid)
			}
		})
	}
}

func TestVelocityMps_Mps(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        float64
		valid       bool
	}{
		{
			name:        "positive",
			measurement: 1.0,
			want:        1.0,
			valid:       true,
		},
		{
			name:        "negative",
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
			v := VelocityMps{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			got := v.Mps()
			if !tests.CloseEnough(got, tc.want, 1e-6) {
				t.Errorf("got %v, want %v", got, tc.want)
			}

			if v.Valid() != tc.valid {
				t.Errorf("got %v, want %v", v.Valid(), tc.valid)
			}
		})
	}
}

func TestVelocityMps_Set(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        float64
		valid       bool
	}{
		{
			name:        "positive",
			measurement: 1.0,
			want:        1.0,
			valid:       true,
		},
		{
			name:        "negative",
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
			v := VelocityMps{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			if v.measurement != tc.want {
				t.Errorf("got %v, want %v", v.measurement, tc.want)
			}

			if v.valid != tc.valid {
				t.Errorf("got %v, want %v", v.valid, tc.valid)
			}
		})
	}
}

func TestVelocityMps_ToFps(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        float64
		valid       bool
	}{
		{
			name:        "positive",
			measurement: 1.0,
			want:        feetPerMeter,
			valid:       true,
		},
		{
			name:        "negative",
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
			v := VelocityMps{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			vf := v.ToFps()
			if vf.Fps() != tc.want {
				t.Errorf("got %v, want %v", vf.Fps(), tc.want)
			}

			if vf.Valid() != tc.valid {
				t.Errorf("got %v, want %v", vf.Valid(), tc.valid)
			}
		})
	}
}

func TestVelocityMps_ToKts(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        float64
		valid       bool
	}{
		{
			name:        "positive",
			measurement: 1.0,
			want:        feetPerMeter * secondsInHour / feetPerNauticalMile,
			valid:       true,
		},
		{
			name:        "negative",
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
			v := VelocityMps{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			vk := v.ToKts()
			if vk.Kts() != tc.want {
				t.Errorf("got %v, want %v", vk.Kts(), tc.want)
			}

			if vk.Valid() != tc.valid {
				t.Errorf("got %v, want %v", vk.Valid(), tc.valid)
			}
		})
	}
}

func TestVelocityMps_ToKph(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        float64
		valid       bool
	}{
		{
			name:        "positive",
			measurement: 1.0,
			want:        secondsInHour / 1000,
			valid:       true,
		},
		{
			name:        "negative",
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
			v := VelocityMps{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			vk := v.ToKph()
			if vk.Kph() != tc.want {
				t.Errorf("got %v, want %v", vk.Kph(), tc.want)
			}

			if vk.Valid() != tc.valid {
				t.Errorf("got %v, want %v", vk.Valid(), tc.valid)
			}
		})
	}
}

func TestVelocityMps_ToMph(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        float64
		valid       bool
	}{
		{
			name:        "positive",
			measurement: 1.0,
			want:        feetPerMeter / feetPerStatuteMile * secondsInHour,
			valid:       true,
		},
		{
			name:        "negative",
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
			v := VelocityMps{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			vm := v.ToMph()
			if vm.Mph() != tc.want {
				t.Errorf("got %v, want %v", vm.Mph(), tc.want)
			}

			if vm.Valid() != tc.valid {
				t.Errorf("got %v, want %v", vm.Valid(), tc.valid)
			}
		})
	}
}

func TestVelocityMps_ToMps(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        float64
		valid       bool
	}{
		{
			name:        "positive",
			measurement: 1.0,
			want:        1.0,
			valid:       true,
		},
		{
			name:        "negative",
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
			v := VelocityMps{}
			if err := v.Set(tc.measurement); err != nil {
				if tc.valid {
					t.Errorf("unexpected error: %v", err)
				}
			}

			vm := v.ToMps()
			if vm.Mps() != tc.want {
				t.Errorf("got %v, want %v", vm.Mps(), tc.want)
			}

			if vm.Valid() != tc.valid {
				t.Errorf("got %v, want %v", vm.Valid(), tc.valid)
			}
		})
	}
}

func TestVelocityMps_Units(t *testing.T) {
	v := VelocityMps{}
	if v.Units() != MetersPerSecond {
		t.Errorf("got %v, want %v", v.Units(), MetersPerSecond)
	}
}

func TestVelocityMps_Valid(t *testing.T) {
	tt := []struct {
		name        string
		measurement float64
		want        bool
	}{
		{
			name:        "positive",
			measurement: 1.0,
			want:        true,
		},
		{
			name:        "negative",
			measurement: -1.0,
			want:        false,
		},
		{
			name:        "zero",
			measurement: 0.0,
			want:        true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v := VelocityMps{}
			_ = v.Set(tc.measurement)

			if v.Valid() != tc.want {
				t.Errorf("got %v, want %v", v.Valid(), tc.want)
			}
		})
	}
}
