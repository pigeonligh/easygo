package errors

import (
	"errors"
	"fmt"
)

type Errors struct {
	errs []error
}

func New(v ...string) error {
	errs := []error{}
	for _, s := range v {
		errs = append(errs, errors.New(s))
	}
	return &Errors{
		errs: errs,
	}
}

func Wrap(v ...error) error {
	errs := []error{}
	errs = append(errs, v...)
	return &Errors{
		errs: errs,
	}
}

func Merge(a error, b error) error {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	errsa, oka := a.(*Errors)
	errsb, okb := b.(*Errors)

	errs := []error{}

	if oka {
		errs = append(errs, errsa.errs...)
	} else {
		errs = append(errs, a)
	}
	if okb {
		errs = append(errs, errsb.errs...)
	} else {
		errs = append(errs, b)
	}
	return &Errors{
		errs: errs,
	}
}

func Append(e error, v ...error) error {
	if e == nil {
		return Wrap(v...)
	}
	ex, ok := e.(*Errors)
	if !ok {
		ex = Wrap(e).(*Errors)
	}
	errs := []error{}
	errs = append(errs, ex.errs...)
	errs = append(errs, v...)
	return &Errors{
		errs: errs,
	}
}

func (e *Errors) Error() string {
	ret := ""
	for _, err := range e.errs {
		if ret != "" {
			ret += ", "
		}
		ret += err.Error()
	}
	return fmt.Sprintf("[%s]", ret)
}
