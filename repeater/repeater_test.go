package repeater

import (
	"fmt"
	"testing"
	"time"
)

func TestRepeater(t *testing.T) {
	var repeater Repeater
	active := true

	repeater.InitPreset("wave")

	go doTask(&active)

	for active {
		fmt.Printf("\r%s", repeater.Repeat())
		time.Sleep(time.Duration(100) * time.Millisecond)
	}
}

func doTask(active *bool) {
	sleepSeconds(10)

	*active = false
}

func sleepSeconds(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}
