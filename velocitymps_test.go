package wx

import "testing"

func TestVelocityMps_ImplementVelocity(t *testing.T) {
	var v Velocity = &VelocityMps{}
	_ = v
}
