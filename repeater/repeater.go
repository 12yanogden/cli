package repeater

import (
	"github.com/12yanogden/errors"
	"github.com/12yanogden/maps"
)

type Preset struct {
	Pattern []string
	Width   int
}

type Repeater struct {
	Index   int
	Pattern []string
	Width   int
}

var presets = map[string]Preset{
	"line_spinner": {
		Pattern: []string{"|", "/", "-", `\`},
		Width:   1,
	},
	"wave": {
		Pattern: []string{"▁", "▂", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃", "▁", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		Width:   8,
	},
	"braille": {
		Pattern: []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"},
		Width:   1,
	},
}

func (repeater *Repeater) Init() {
	preset := presets["line_spinner"]

	repeater.InitCustom(
		preset.Pattern,
		preset.Width,
	)
}

func (repeater *Repeater) InitPreset(presetName string) {
	if !maps.HasKey(presets, presetName) {
		errors.Scream("preset '" + presetName + "' not found")
	}

	preset := presets[presetName]

	repeater.InitCustom(
		preset.Pattern,
		preset.Width,
	)
}

func (repeater *Repeater) InitCustom(pattern []string, width int) {
	repeater.Index = 0
	repeater.Pattern = pattern
	repeater.Width = width
}

func (repeater *Repeater) Repeat() string {
	out := repeater.Pattern[(repeater.Index)%len(repeater.Pattern)]
	repeater.Index++

	remainingCount := repeater.Width - 1
	for i := 0; i < remainingCount; i++ {
		out += repeater.Pattern[(i+repeater.Index)%len(repeater.Pattern)]
	}

	return out
}
