package wx

import "testing"

func TestVelocityMph_ImplementsVelocity(t *testing.T) {
	var v Velocity = &VelocityMph{}
	_ = v
}
