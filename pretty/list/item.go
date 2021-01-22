package list

import (
	"bytes"
)

type Item struct {
	text     string
	children []*Item
}

func New(text string) *Item {
	return &Item{
		text:     text,
		children: make([]*Item, 0, 10),
	}
}

func (it *Item) Get(index int) *Item {
	if mod := len(it.children); mod > 0 {
		index = index % mod
		if index < 0 {
			index += mod
		}
	}
	return it.children[index]
}

func (it *Item) addItem(item *Item) {
	it.children = append(it.children, item)
}

func (it *Item) insertItem(item *Item, index int) bool {
	if index >= 0 && index <= len(it.children) {
		it.children = append(it.children, nil)
		for i := len(it.children); i > index; i-- {
			it.children[i] = it.children[i-1]
		}
		it.children[index] = item
		return true
	}
	return false
}

func (it *Item) Add(text string) {
	it.addItem(New(text))
}

func (it *Item) Insert(text string, index int) bool {
	return it.insertItem(New(text), index)
}

func (it *Item) SetText(text string) {
	it.text = text
}

func (it *Item) GetText() string {
	return it.text
}

func (it *Item) ToString(r Render) string {
	/*
		list := make([]showItem, 0)
		list = it.addToItemList(0, 1, 1, list)

		ret := bytes.NewBuffer([]byte{})
		for _, item := range list {
			_, _ = ret.WriteString(r.Indent(item.level))
			_, _ = ret.WriteString(r.Index(item.level, item.index, item.total))
			_, _ = ret.WriteString(item.text)
			_, _ = ret.WriteString(r.EndLine(item.level))
		}
		return ret.String()
	*/
	ret := bytes.NewBuffer([]byte{})
	it.writeToBuffer(0, 1, []bool{}, ret, r)
	return ret.String()
}
