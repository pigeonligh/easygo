package flags

import (
	"fmt"
	"net"
	"reflect"
	"time"

	"github.com/spf13/pflag"
)

func getRefValueOfPointer(
	ptr interface{},
) (refValue reflect.Value, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	refAddr := reflect.ValueOf(ptr)
	refValue = refAddr.Elem()
	err = nil
	return
}

func ObjectVar(fs *pflag.FlagSet, ptr interface{}, prefix string) error {
	ref, err := getRefValueOfPointer(ptr)
	if err != nil {
		return err
	}

	refType := ref.Type()
	fieldCount := ref.NumField()
	for i := 0; i < fieldCount; i++ {
		fieldType := refType.Field(i)
		fieldValue := ref.Field(i)
		name := prefix + fieldType.Tag.Get("name")
		help := fieldType.Tag.Get("help")
		shorthand := fieldType.Tag.Get("shorthand")
		if len(prefix) > 0 {
			shorthand = ""
		}
		fieldprefix := prefix + fieldType.Tag.Get("prefix")

		obj := fieldValue.Addr().Interface()
		switch ptr := obj.(type) {
		case *bool:
			fs.BoolVarP(ptr, name, shorthand, *ptr, help)

		case *int:
			fs.IntVarP(ptr, name, shorthand, *ptr, help)
		case *[]int:
			fs.IntSliceVarP(ptr, name, shorthand, *ptr, help)
		case *int32:
			fs.Int32VarP(ptr, name, shorthand, *ptr, help)
		case *[]int32:
			fs.Int32SliceVarP(ptr, name, shorthand, *ptr, help)
		case *int64:
			fs.Int64VarP(ptr, name, shorthand, *ptr, help)
		case *[]int64:
			fs.Int64SliceVarP(ptr, name, shorthand, *ptr, help)

		case *uint:
			fs.UintVarP(ptr, name, shorthand, *ptr, help)
		case *[]uint:
			fs.UintSliceVarP(ptr, name, shorthand, *ptr, help)
		case *uint32:
			fs.Uint32VarP(ptr, name, shorthand, *ptr, help)
		case *uint64:
			fs.Uint64VarP(ptr, name, shorthand, *ptr, help)

		case *float32:
			fs.Float32VarP(ptr, name, shorthand, *ptr, help)
		case *[]float32:
			fs.Float32SliceVarP(ptr, name, shorthand, *ptr, help)
		case *[]float64:
			fs.Float64SliceVarP(ptr, name, shorthand, *ptr, help)

		case *time.Duration:
			fs.DurationVarP(ptr, name, shorthand, *ptr, help)
		case *[]time.Duration:
			fs.DurationSliceVarP(ptr, name, shorthand, *ptr, help)

		case *string:
			fs.StringVarP(ptr, name, shorthand, *ptr, help)
		case *[]string:
			fs.StringSliceVarP(ptr, name, shorthand, *ptr, help)

		case *net.IP:
			fs.IPVarP(ptr, name, shorthand, *ptr, help)
		case *[]net.IP:
			fs.IPSliceVarP(ptr, name, shorthand, *ptr, help)
		case *net.IPMask:
			fs.IPMaskVarP(ptr, name, shorthand, *ptr, help)
		case *net.IPNet:
			fs.IPNetVarP(ptr, name, shorthand, *ptr, help)

		default:
			return ObjectVar(fs, ptr, fieldprefix)
		}
	}
	return nil
}
