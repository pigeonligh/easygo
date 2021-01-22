package table

import (
	"fmt"

	"github.com/pigeonligh/easygo/pretty/text"
)

type TableHeader struct {
	Text string

	Width int

	HeaderAlign text.Alignment
	TextAlign   text.Alignment
}

type Table struct {
	headers  []TableHeader
	maxWidth []int

	rows []rowType
}

func NewByHeaders(headers []TableHeader) *Table {
	hdr := make([]TableHeader, 0, len(headers))
	maxWidth := make([]int, 0, len(headers))
	for _, header := range headers {
		hdr = append(hdr, header)
		maxWidth = append(maxWidth, len(header.Text))
	}
	return &Table{
		headers:  hdr,
		maxWidth: maxWidth,
		rows:     make([]rowType, 0),
	}
}

func New(texts []string) *Table {
	hdr := make([]TableHeader, 0, len(texts))
	for _, t := range texts {
		hdr = append(hdr, TableHeader{
			Text:        t,
			Width:       0,
			HeaderAlign: text.AlignLeft,
			TextAlign:   text.AlignLeft,
		})
	}
	return NewByHeaders(hdr)
}

func (table *Table) AddRow(row Row) {
	table.rows = append(table.rows, row)
	rowSize := len(row)
	widthSize := len(table.maxWidth)
	for rowSize < widthSize {
		row = append(row, "")
		rowSize++
	}
	if rowSize > widthSize {
		row = row[0:widthSize]
	}
	for i := 0; i < widthSize; i++ {
		t := fmt.Sprint(row[i])
		if table.maxWidth[i] < len(t) {
			table.maxWidth[i] = len(t)
		}
	}
}

func (table *Table) AddRows(rows []Row) {
	for _, row := range rows {
		table.AddRow(row)
	}
}

func (table *Table) AddSeparator() {
	table.rows = append(table.rows, separator(0))
}

func (table *Table) Render() string {
	width := make([]int, 0, len(table.headers))
	headers := make([]interface{}, 0, len(table.headers))

	hdrAligns := make([]text.Alignment, 0, len(table.headers))
	aligns := make([]text.Alignment, 0, len(table.headers))

	for i, header := range table.headers {
		if header.Width > 0 {
			width = append(width, header.Width)
		} else {
			width = append(width, table.maxWidth[i])
		}
		headers = append(headers, header.Text)
		hdrAligns = append(hdrAligns, header.HeaderAlign)
		aligns = append(aligns, header.TextAlign)
	}

	ret := separator(0).Render(width, nil) + "\n"
	ret += Row(headers).Render(width, hdrAligns) + "\n"
	ret += separator(0).Render(width, nil) + "\n"
	for _, row := range table.rows {
		ret += row.Render(width, aligns) + "\n"
	}
	ret += separator(0).Render(width, nil) + "\n"
	return ret
}
