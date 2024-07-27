package wx

import "testing"

func TestVelocityKph_ImplementsVelocity(t *testing.T) {
	var v Velocity = &VelocityKph{}
	_ = v
}
