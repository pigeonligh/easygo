package list

import (
	"fmt"
	"strings"
)

type numberRenderStruct struct {
}

func (r numberRenderStruct) Indent(level int, ends []bool) string {
	return strings.Repeat("\t", level)
}

func (r numberRenderStruct) Index(level, index int, isEnd bool) string {
	return fmt.Sprintf("%d. ", index)
}

func (r numberRenderStruct) EndLine(level int) string {
	return "\n"
}

func (r numberRenderStruct) EndOfItem(level int) string {
	return ""
}

var NumberRender = numberRenderStruct{}
