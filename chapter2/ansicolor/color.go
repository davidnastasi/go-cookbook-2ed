package ansicolor

import "fmt"

type Color int

const (
	Black Color = iota - 1
	None
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)
type ColorText struct {
	TextColor Color
	Text string
}

func (r *ColorText) String() string {
	if r.TextColor == None {
		return r.Text
	}
	value := 30
	if r.TextColor != Black {
		value += int(r.TextColor)
	}
	return fmt.Sprintf("\033[0;%dm%s\033[0m", value, r.Text)
}
