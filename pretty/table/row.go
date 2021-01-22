package table

import (
	"fmt"
	"strings"

	"github.com/pigeonligh/easygo/text"
)

type rowType interface {
	Render(width []int, align []text.Alignment) string
}

type Row []interface{}
type separator int

func (row Row) Render(width []int, align []text.Alignment) string {
	rowSize := len(row)
	if rowSize == 0 {
		return ""
	}

	ret := "|"

	for i := 0; i < rowSize; i++ {
		t := fmt.Sprint(row[i])
		if width[i] == 0 {
			ret += " " + t + " |"
		} else {
			ret += " " + text.Wrap(t, width[i], align[i]) + " |"
		}
	}
	return ret
}

func (separator) Render(width []int, align []text.Alignment) string {
	ret := "+"
	for _, w := range width {
		ret += strings.Repeat("-", w+2) + "+"
	}
	return ret
}
