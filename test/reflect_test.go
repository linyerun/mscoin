package test

import (
	"reflect"
	"testing"
)

func TestIsValid(t *testing.T) {
	type A struct {
	}

	println(reflect.ValueOf(&A{}).IsValid())
	println(reflect.ValueOf(A{}).IsValid())
}
