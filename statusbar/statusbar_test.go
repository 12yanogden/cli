package statusbar

import (
	"fmt"
	"testing"
	"time"
)

func TestStatusBar(t *testing.T) {
	var bar StatusBar
	active := true

	bar.Init("Gotta get stuff done")

	go doTask(&active)

	for active {
		fmt.Printf("\r%s", bar.Run())
		time.Sleep(100 * time.Millisecond)
	}

	active = true

	for active {
		fmt.Printf("\r%s", bar.Pass(&active))
		time.Sleep(30 * time.Millisecond)
	}

	fmt.Println()
}

func doTask(active *bool) {
	sleepSeconds(3)

	*active = false
}

func sleepSeconds(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}
