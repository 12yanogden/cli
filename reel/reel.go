package reel

import "github.com/12yanogden/cli/colors"

type Reel struct {
	Index  int
	Frames []string
	Width  int
}

func (r *Reel) Init(frames []string, width int) {
	r.Index = 0
	r.Frames = frames
	r.Width = width

	r.padFrames()
}

func (r *Reel) Play() string {
	out := r.Frames[r.Index]

	r.Index = (r.Index + 1) % len(r.Frames)

	return out
}

func (r *Reel) padFrames() {
	for i := range (*r).Frames {
		colors.Left(&(*r).Frames[i], &(*r).Width)
	}
}

func (r *Reel) AddFrames(frames []string) {
	(*r).Frames = append((*r).Frames, frames...)

	r.padFrames()
}
