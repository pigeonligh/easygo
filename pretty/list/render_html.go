package list

import (
	"fmt"
	"strings"
)

type HTMLRenderStruct struct {
	classPrefix string
}

func NewHTMLRender(classPrefix string) HTMLRenderStruct {
	return HTMLRenderStruct{
		classPrefix: classPrefix,
	}
}

func (r HTMLRenderStruct) Indent(level int, ends []bool) string {
	ret := strings.Repeat("\t", level)
	ret += fmt.Sprintf("<ul class=\"%s-%d\">\n", r.classPrefix, level)
	ret += strings.Repeat("\t", level+1)
	return ret
}

func (r HTMLRenderStruct) Index(level, index int, isEnd bool) string {
	return "<li>"
}

func (r HTMLRenderStruct) EndLine(level int) string {
	return "</li>\n"
}

func (r HTMLRenderStruct) EndOfItem(level int) string {
	return strings.Repeat("\t", level) + "</ul>\n"
}
