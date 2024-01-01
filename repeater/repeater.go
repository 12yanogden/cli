package repeater

import (
	"fmt"
	"time"

	"github.com/12yanogden/maps"
)

type Preset struct {
	Pattern   []string
	Width     int
	Frequency int
}

type Repeater struct {
	Index     int
	Pattern   []string
	Width     int
	Frequency int
}

var presets = map[string]Preset{
	"line_spinner": {
		Pattern:   []string{"|", "/", "-", `\`},
		Width:     1,
		Frequency: 100,
	},
	"wave": {
		Pattern:   []string{"▁", "▂", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃", "▁", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		Width:     6,
		Frequency: 100,
	},
	"braille": {
		Pattern:   []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"},
		Width:     1,
		Frequency: 100,
	},
}

func (repeater *Repeater) Init() {
	preset := presets["line_spinner"]

	repeater.InitCustom(
		preset.Pattern,
		preset.Width,
		preset.Frequency,
	)
}

func (repeater *Repeater) InitPreset(presetName string) {
	if !maps.HasKey(presets, presetName) {
		fmt.Println("preset '" + presetName + "' not found")
	}

	preset := presets[presetName]

	repeater.InitCustom(
		preset.Pattern,
		preset.Width,
		preset.Frequency,
	)
}

func (repeater *Repeater) InitCustom(pattern []string, width int, frequencyInMs int) {
	repeater.Index = 0
	repeater.Pattern = pattern
	repeater.Width = width
	repeater.Frequency = frequencyInMs
}

func (repeater *Repeater) Repeat() {
	fmt.Printf("\r%s", repeater.Pattern[(repeater.Index)%len(repeater.Pattern)])
	repeater.Index++

	remainingCount := repeater.Width - 1
	for i := 0; i < remainingCount; i++ {
		fmt.Printf("%s", repeater.Pattern[(i+repeater.Index)%len(repeater.Pattern)])
	}

	time.Sleep(time.Duration(repeater.Frequency) * time.Millisecond)
}

func (repeater *Repeater) End() {
	fmt.Println()
}
