package text

import (
	"fmt"
	"strings"
)

type Color int
type Colors []Color

const (
	FontReset Color = iota
	FontBold
	FontFaint
	FontItalic
	FontUnderline
	FontBlinkSlow
	FontBlinkRapid
	FontReverseVideo
	FontConcealed
	FontCrossedOut
)

const (
	FgBlack Color = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

const (
	FgHiBlack Color = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

const (
	BgBlack Color = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

const (
	BgHiBlack Color = iota + 100
	BgHiRed
	BgHiGreen
	BgHiYellow
	BgHiBlue
	BgHiMagenta
	BgHiCyan
	BgHiWhite
)

func (c Colors) Setup(s string) string {
	colors := []string{}
	for _, v := range c {
		colors = append(colors, fmt.Sprint(v))
	}
	if len(colors) == 0 {
		colors = append(colors, "0")
	}
	pre := "\033[" + strings.Join(colors, ";") + "m"
	suf := "\033[0m"
	return pre + s + suf
}
