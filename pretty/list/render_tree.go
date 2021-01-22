package list

type treeRenderStruct struct {
	linkHeader bool
}

func (r treeRenderStruct) Indent(level int, ends []bool) string {
	if level == 0 {
		return ""
	}
	ret := ""
	for i := 0; i < level-1; i++ {
		if ends[i] {
			ret += "   "
		} else {
			ret += "│  "
		}
	}
	return ret
}

func (r treeRenderStruct) Index(level, index int, isEnd bool) string {
	if level == 0 {
		return ""
	}
	if !r.linkHeader && level == 1 && index == 1 {
		if isEnd {
			return " ─ "
		} else {
			return "╭─ "
		}
	}
	var sign string
	if isEnd {
		sign = "╰"
	} else {
		sign = "├"
	}
	return sign + "─ "
}

func (r treeRenderStruct) EndLine(level int) string {
	return "\n"
}

func (r treeRenderStruct) EndOfItem(level int) string {
	return ""
}

var TreeRender = treeRenderStruct{linkHeader: true}
var TreeRenderWithoutHeader = treeRenderStruct{linkHeader: false}
