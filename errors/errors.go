package merror

import (
	"errors"
	"fmt"
)

type Errors struct {
	errs []error
}

func New(v ...string) *Errors {
	errs := []error{}
	for _, s := range v {
		errs = append(errs, errors.New(s))
	}
	return &Errors{
		errs: errs,
	}
}

func Wrap(v ...error) *Errors {
	errs := []error{}
	errs = append(errs, v...)
	return &Errors{
		errs: errs,
	}
}

func Merge(a error, b error) *Errors {
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

func Append(e *Errors, v ...error) *Errors {
	if e == nil {
		return Wrap(v...)
	}
	errs := []error{}
	errs = append(errs, e.errs...)
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
