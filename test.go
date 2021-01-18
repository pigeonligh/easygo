// +build vendor

package main

import (
	"github.com/pigeonligh/easygo/collections/meter"
	log "github.com/pigeonligh/easygo/elog"
	"github.com/pigeonligh/easygo/errors"
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

func main() {
	logInit()

	testLog()
	testErrors()
	testMeter()
}
