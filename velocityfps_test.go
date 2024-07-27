package wx

import "testing"

func TestVelocityFps_ImplementsVelocity(t *testing.T) {
	var v Velocity = &VelocityFps{}
	_ = v
}
