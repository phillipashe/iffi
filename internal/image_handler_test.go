package image_handler

import "testing"

func TestHandleImage(t *testing.T) {
	expected := "hello world"
	if actual := HandleImage(); actual != expected {
		t.Errorf("This isn't right, %q", actual)
	}
}
