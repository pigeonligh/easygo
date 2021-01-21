// +build vendor

package main

import (
	"fmt"

	"github.com/pigeonligh/easygo/collections/counter"
	"github.com/pigeonligh/easygo/collections/meter"
	log "github.com/pigeonligh/easygo/elog"
	"github.com/pigeonligh/easygo/errors"
	"github.com/pigeonligh/easygo/pretty/list"
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
	fmt.Println(l.ToString(list.HyphenReduceRender))
	fmt.Println(l.ToString(list.TreeRenderWithoutHeader))
	fmt.Println(l.ToString(list.NumberRender))
}

func main() {
	logInit()

	testLog()
	testErrors()

	testMeter()
	testCounter()

	testList()
}
