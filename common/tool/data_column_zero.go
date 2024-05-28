package tool

import (
	"github.com/pkg/errors"
	"reflect"
)

func DataColumnZero(data any) error {
	typeOf := reflect.TypeOf(data)
	valueOf := reflect.ValueOf(data)
	if typeOf.Kind() != reflect.Pointer {
		return errors.New("must be pointer")
	}

	//member
	ele := typeOf.Elem()
	valueEle := valueOf.Elem()
	for i := 0; i < ele.NumField(); i++ {
		field := ele.Field(i)
		value := valueEle.Field(i)
		//field.Tag.Get("default")
		kind := field.Type.Kind()
		if kind == reflect.Int {
			//根据设置的tag进行值的设置
			value.Set(intZero())
		}
		if kind == reflect.Int32 {
			value.Set(int32Zero())
		}
		if kind == reflect.Int64 {
			value.Set(int64Zero())
		}
		if kind == reflect.String {
			value.Set(stringZero())
		}
		if kind == reflect.Float64 {
			value.Set(float64Zero())
		}
		if kind == reflect.Float32 {
			value.Set(float32Zero())
		}
	}

	return nil
}

func stringZero() reflect.Value {
	var i = ""
	return reflect.ValueOf(i)
}

func intZero() reflect.Value {
	var i = 0
	return reflect.ValueOf(i)
}

func int32Zero() reflect.Value {
	var i int32 = 0
	return reflect.ValueOf(i)
}
func int64Zero() reflect.Value {
	var i int64 = 0
	return reflect.ValueOf(i)
}

func float64Zero() reflect.Value {
	var i float64 = 0
	return reflect.ValueOf(i)
}
func float32Zero() reflect.Value {
	var i float32 = 0
	return reflect.ValueOf(i)
}
