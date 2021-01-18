// +build vendor

package main

import (
	log "github.com/pigeonligh/easygo/elog"
	"github.com/pigeonligh/easygo/errors"
)

func getError() error {
	return errors.New("hello")
}

func main() {
	log.Default()

	var errs error
	for i := 0; i < 5; i++ {
		err := getError()
		errs = errors.Merge(errs, err)
	}

	l := log.With(map[string]string{
		"fruit": "apple",
	})

	l.Info(errs)
	log.Info(errs)
}
