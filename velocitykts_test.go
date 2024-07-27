package wx

import "testing"

func TestVelocityKts_ImplementsVelocity(t *testing.T) {
	var v Velocity = &VelocityKts{}
	_ = v
}
