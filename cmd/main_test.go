package main

import (
	"errors"
	dmpstruct "github.com/evityuk/dmpstruct/pkg/typestruct"
	"net/url"
	"reflect"
	"testing"
)

type NamedType int32
type Comparator func(int, int) int
type IntMap map[string]int
type SomeInterface interface {
	SomeMethod(string) (string, error)
}
type SomeInterfaceImpl struct{}

func (SomeInterfaceImpl) SomeMethod(arg1 string) (string, error) {
	return arg1, nil
}

var intPtrGetter = constPtr(10)
var rwChannel = make(chan string)
var interfaceImpl = SomeInterfaceImpl{}

func comparator(a int, b int) int {
	return b - a
}

var user = url.UserPassword("admin", "password")
var errInstance = errors.New("Some error")

func TestDumpCmdFlag(t *testing.T) {
	for _, c := range []struct {
		in, want interface{}
		err      error
	}{} {
		got, err := dmpstruct.Dump(c.in)

		if !reflect.DeepEqual(err, c.err) || !reflect.DeepEqual(got, c.want) {
			t.Errorf("Dump(\n%#v\n) == (\n%#v, \n%v\n), want (\n%#v, \n%v\n)", c.in, got, err, c.want, c.err)
		}
	}
}

func constPtr(ptr int) func() *int {
	return func() *int {
		return &ptr
	}
}
