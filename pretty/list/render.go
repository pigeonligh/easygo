package list

import (
	"bytes"
)

type Render interface {
	Indent(level int, ends []bool) string
	Index(level, index int, isEnd bool) string
	EndLine(level int) string
	EndOfItem(level int) string
}

func (it *Item) writeToBuffer(nowLevel, index int, ends []bool, writer *bytes.Buffer, r Render) {
	var end bool = true
	if nowLevel > 0 {
		end = ends[len(ends)-1]
	}

	_, _ = writer.WriteString(r.Indent(nowLevel, ends))
	_, _ = writer.WriteString(r.Index(nowLevel, index, end))
	_, _ = writer.WriteString(it.text)
	_, _ = writer.WriteString(r.EndLine(nowLevel))

	total := len(it.children)
	newEnds := make([]bool, len(ends))
	copy(newEnds, ends)
	newEnds = append(newEnds, false)
	for i, item := range it.children {
		newEnds[nowLevel] = i+1 == total
		item.writeToBuffer(nowLevel+1, i+1, newEnds, writer, r)
	}

	_, _ = writer.WriteString(r.EndOfItem(nowLevel))
}
