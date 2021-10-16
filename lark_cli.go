package lark_docs_md

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/chyroc/lark"
)

type formatOpt struct {
	cli          *lark.Lark
	dir          string
	staticPrefix string
}

func parchOpt(data interface{}, opt *formatOpt) {
	vv := reflect.ValueOf(data)

	if vv.Kind() == reflect.Ptr {
		vv = vv.Elem()
	}

	switch vv.Kind() {
	case reflect.Struct:
		// fmt.Println(vv.Type().Name())
		fieldOpt := vv.FieldByName("opt")
		// fmt.Println(vv.Type().Name(),fieldOpt)
		if !fieldOpt.IsValid() {
			for i := 0; i < vv.NumField(); i++ {
				fieldV := vv.Field(i)
				if !fieldV.CanInterface() {
					// fmt.Println(vv.Type().Name())
					panic(fieldV)
				}
				parchOpt(fieldV.Interface(), opt)
			}
		} else {
			setUnexportedField(fieldOpt, opt)
			// fieldOpt.Set(reflect.ValueOf(opt))
		}
	case reflect.Slice:
		for i := 0; i < vv.Len(); i++ {
			itemV := vv.Index(i)
			parchOpt(itemV.Interface(), opt)
		}
	case reflect.String,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Uintptr,
		reflect.Bool,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128:
		return
	case reflect.Invalid:
		return
	default:
		panic(fmt.Sprintf("不支持: %s", vv.Kind()))
	}
}

func setUnexportedField(field reflect.Value, value interface{}) {
	reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).
		Elem().
		Set(reflect.ValueOf(value))
}
