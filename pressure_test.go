package wx

import "testing"

func TestPressureUnit_String(t *testing.T) {
	tests := []struct {
		name string
		p    PressureUnit
		want string
	}{
		{"mb", HPa, "hPa"},
		{"inHg", InHg, "inHg"},
		{"psi", Psi, "psi"},
		{"kPa", KPa, "kPa"},
		{"mb", Mb, "mb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("PressureUnit.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
