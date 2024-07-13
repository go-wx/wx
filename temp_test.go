package wx

import "testing"

func TestTempUnit_String(t *testing.T) {
	tests := []struct {
		name string
		u    TempUnit
		want string
	}{
		{"Celsius", Celsius, "C"},
		{"Fahrenheit", Fahrenheit, "F"},
		{"Kelvin", Kelvin, "K"},
		{"Rankine", Rankine, "R"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.String(); got != tt.want {
				t.Errorf("TempUnit.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
