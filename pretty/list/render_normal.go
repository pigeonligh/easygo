package list

import "strings"

type NormalRenderStruct struct {
	Sign string
}

func NewNormalRender(sign string) NormalRenderStruct {
	return NormalRenderStruct{
		Sign: sign,
	}
}

func (r NormalRenderStruct) Indent(level int, ends []bool) string {
	return strings.Repeat("\t", level)
}

func (r NormalRenderStruct) Index(level, index int, isEnd bool) string {
	return r.Sign + " "
}

func (r NormalRenderStruct) EndLine(level int) string {
	return "\n"
}

func (r NormalRenderStruct) EndOfItem(level int) string {
	return ""
}

var HyphenRender = NewNormalRender("-")
var StarRender = NewNormalRender("*")
var PlusRender = NewNormalRender("+")
var CircleRender = NewNormalRender("●")
