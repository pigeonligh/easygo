package list

import "strings"

type ReduceRenderStruct struct {
	Sign string
}

func NewReduceRender(sign string) ReduceRenderStruct {
	return ReduceRenderStruct{
		Sign: sign,
	}
}

func (r ReduceRenderStruct) Indent(level int, ends []bool) string {
	if level == 0 {
		return ""
	}
	return strings.Repeat("\t", level-1)
}

func (r ReduceRenderStruct) Index(level, index int, isEnd bool) string {
	if level == 0 {
		return ""
	}
	return r.Sign + " "
}

func (r ReduceRenderStruct) EndLine(level int) string {
	return "\n"
}

func (r ReduceRenderStruct) EndOfItem(level int) string {
	return ""
}

var HyphenReduceRender = NewReduceRender("-")
var StarReduceRender = NewReduceRender("*")
var PlusReduceRender = NewReduceRender("+")
var CircleReduceRender = NewNormalRender("●")
