// +build vendor

package main

import (
	"fmt"
	"time"

	"github.com/pigeonligh/easygo/collections/counter"
	"github.com/pigeonligh/easygo/collections/meter"
	log "github.com/pigeonligh/easygo/elog"
	"github.com/pigeonligh/easygo/errors"
	"github.com/pigeonligh/easygo/pretty/list"
	"github.com/pigeonligh/easygo/pretty/table"
	"github.com/pigeonligh/easygo/pretty/text"
)

func logInit() {
	log.Default() // or log.Debug()
}

func getError() error {
	return errors.New("hello")
}

func testLog() {
	l := log.With(map[string]string{
		"fruit": "apple",
	})

	l.Info("Hello world")
}

func testErrors() {
	var errs error
	for i := 0; i < 5; i++ {
		err := getError()
		errs = errors.Merge(errs, err)
	}

	log.Info(errs)
}

func testMeter() {
	m := meter.New()
	n := meter.Make(1, 7, 4)
	m.AddMeter(n)
	m.Add(8)
	m.Add(5)
	log.Info(m.Sum(), m.Average(), m.Max(), m.Min())

	log.Info(meter.Sum(1, 2, 3, 4, 5))
}

func testCounter() {
	c := counter.New()
	c.Pushes("apple", 5)
	c.Pushes("banana", 3)
	d := counter.New()
	d.Pushes("apple", 2)
	d.Pushes("banana", 5)

	log.Info(c.Add(d))
	log.Info(c.Sub(d))
	log.Info(c.Max(d))
	log.Info(c.Min(d))
}

func testList() {
	l := list.New("Hello world")
	l.Add("hello 1")
	l.Get(-1).Add("world 1")
	l.Get(-1).Get(-1).Add("!!!")
	l.Get(-1).Add("world 2")
	l.Get(-1).Get(-1).Add("!!!")
	l.Add("hello 2")
	l.Get(-1).Add("world 3")
	l.Get(-1).Add("world 4")

	fmt.Println(l.ToString(list.CircleRender))
	// fmt.Println(l.ToString(list.HyphenReduceRender))
	// fmt.Println(l.ToString(list.TreeRenderWithoutHeader))
	// fmt.Println(l.ToString(list.NumberRender))
}

func testTable() {
	t := table.NewByHeaders([]table.TableHeader{
		{Text: "#", HeaderAlign: text.AlignCenter, TextAlign: text.AlignCenter},
		{Text: "ask", HeaderAlign: text.AlignCenter, TextAlign: text.AlignLeft},
		{Text: "answer", HeaderAlign: text.AlignCenter, TextAlign: text.AlignLeft},
	})
	t.AddRow(table.Row{1, "Hello, nice to meet you.", "Nice to meet you, too."})
	t.AddRow(table.Row{2, "How are you?", "I'm fine, thank you."})
	t.AddRow(table.Row{3, "What day is it today?", fmt.Sprintf("It's %s.", time.Now().Format("Monday"))})

	fmt.Println(t.Render())
}

func testText() {
	color := text.Colors{text.FontBold, text.FgGreen, text.BgBlue}

	fmt.Println(color.Setup("hello world!"))
}

func main() {
	logInit()

	testLog()
	testErrors()

	testMeter()
	testCounter()

	testList()
	testTable()
	testText()
}
