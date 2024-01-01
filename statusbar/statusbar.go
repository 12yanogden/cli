package statusbar

import "github.com/12yanogden/cli/repeater"

type StatusBar struct {
	Active bool
	Msg    string
}

func (bar StatusBar) Start(msg string) {
	var repeater repeater.Repeater

	bar.Active = true
	bar.Msg = msg

	repeater.InitPreset("wave")

	for bar.Active {
		repeater.Repeat()
	}

	repeater.End()
}

func (bar StatusBar) Pass() {
	bar.Active = false
}

func (bar StatusBar) Fail() {
	bar.Active = false
}
