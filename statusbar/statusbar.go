package statusbar

import (
	"fmt"
	"strings"

	"github.com/12yanogden/cli/colors"
	"github.com/12yanogden/cli/reel"
	"github.com/12yanogden/cli/repeater"
	"github.com/12yanogden/errors"
	"github.com/12yanogden/intslices"
	"github.com/12yanogden/str"
)

type StatusBar struct {
	Msg    string
	Wave   *(repeater.Repeater)
	Reveal *(reel.Reel)
}

func (bar *StatusBar) Init(msg string) {
	var wave repeater.Repeater
	var reveal reel.Reel
	revealFrames := []string{
		"[        ]",
		"[       ]",
		"[      ]",
		"[     ]",
		"[    ]",
		"[   ]",
		"[  ]",
		"[ ]",
		"[]",
	}

	wave.InitPreset("wave")

	reveal.Init(revealFrames, 10)

	bar.Msg = msg
	bar.Wave = &wave
	bar.Reveal = &reveal
}

func (bar *StatusBar) Run() string {
	if bar.Wave == nil {
		errors.Scream("must init the status bar before running")
	}

	return fmt.Sprintf("[%s] %s", bar.Wave.Repeat(), bar.Msg)
}

func (bar *StatusBar) Pass(active *bool) string {
	return revealStatus(bar, "PASSED", "GREEN", active)
}

func (bar *StatusBar) Fail(active *bool) string {
	return revealStatus(bar, "FAILED", "RED", active)
}

func revealStatus(bar *StatusBar, status string, colorKey string, active *bool) string {
	if !hasStatusFrames(bar, status) {
		(*bar).Reveal.AddFrames(calcStatusFrames(status, colorKey))
	}

	if bar.Reveal.Index == (len(bar.Reveal.Frames) - 1) {
		*active = false
	}

	return fmt.Sprintf("%s %s", (*bar).Reveal.Play(), bar.Msg)
}

func hasStatusFrames(bar *StatusBar, status string) bool {
	return strings.Contains(bar.Reveal.Frames[len(bar.Reveal.Frames)-1], status)
}

func calcStatusFrames(status string, colorKey string) []string {
	status = str.Center(status, len(status)+2)
	frames := []string{}

	for i := range intslices.Seq(0, len(status)) {
		frames = append(frames, "["+colors.COLORS[colorKey]+status[:i]+colors.COLORS["RESET"]+"]")
	}

	return frames
}
