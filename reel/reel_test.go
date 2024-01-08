package reel

import (
	"fmt"
	"testing"
	"time"

	"github.com/12yanogden/intslices"
)

func TestReel(t *testing.T) {
	var reel Reel
	frames := []string{
		"[        ]",
		"[       ]",
		"[      ]",
		"[     ]",
		"[    ]",
		"[   ]",
		"[  ]",
		"[ ]",
		"[]",
		"[ ]",
		"[ P]",
		"[ PA]",
		"[ PAS]",
		"[ PASS]",
		"[ PASSE]",
		"[ PASSED]",
		"[ PASSED ]",
	}

	reel.Init(frames, 10)

	for range intslices.Seq(0, len(frames)-1) {
		fmt.Printf("\r%s", reel.Play())
		time.Sleep(time.Duration(30) * time.Millisecond)
	}

	fmt.Printf("\n")
}
