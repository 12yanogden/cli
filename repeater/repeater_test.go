package repeater

import (
	"testing"
)

func TestWheel(t *testing.T) {
	var repeater Repeater

	repeater.InitPreset("braille")

	for i := 0; i <= 100; i++ {
		repeater.Play()
	}

	repeater.End()
}
