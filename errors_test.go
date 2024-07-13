package wx

import "testing"

func TestNewWxErr(t *testing.T) {
	err := "test error"
	context := "test context"
	expected := "test context: test error"
	actual := NewWxErr(err, context).Error()
	if expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestError_Error(t *testing.T) {
	err := "test error"
	context := "test context"
	expected := "test context: test error"
	actual := Error{err, context}.Error()
	if expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}
