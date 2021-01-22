package text

import "strings"

type Alignment int

const (
	AlignLeft Alignment = iota
	AlignRight
	AlignCenter
)

func Wrap(t string, width int, align Alignment) string {
	if width < 5 && len(t) > width {
		width = 5
	}
	if len(t) > width {
		t = t[0:width-3] + "..."
	}
	tot := width - len(t)
	right := 0
	switch align {
	case AlignLeft:
		right = tot
	case AlignCenter:
		right = tot / 2
	case AlignRight:
		right = 0
	}
	left := tot - right
	return strings.Repeat(" ", left) + t + strings.Repeat(" ", right)
}
