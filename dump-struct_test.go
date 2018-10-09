package dmpstruct

import (
	"errors"
	"fmt"
	"os"
	logrus "github.com/sirupsen/logrus"
//	"net/url"
	"reflect"
	"testing"
)

type Employee struct {
	department string
	position   string
}

type OccupationInfo struct {
	Name string
	code uint
}

type NamedType int64
type Comparator func(int, int) int
type IntMap map[string]int
type SomeInterface interface {
	SomeMethod(string) (string, error)
}
type SomeInterfaceImpl struct{}

func (this SomeInterfaceImpl) SomeMethod(arg1 string) (string, error) {
	return arg1, nil
}

var intPtrGetter = constPtr(10)
var readableChannel = make(chan string)
var interfaceImpl = SomeInterfaceImpl{}

func comparator(a int, b int) int {
	return b - a
}

func TestDump(t *testing.T) {
	cases := []struct {
		in, want interface{}
		err      error
	}{
		{struct {
			Name       string
			age        int8
			Occupation OccupationInfo
			Employee
		}{
			"Dan", 50, OccupationInfo{"Boston", 33}, Employee{"Copywriting", "Editor in chief"},
		}, map[string]interface{}{
			"age":  fmt.Sprintf(FORMAT_UNEXPORTED_STRING, "age", "int8"),
			"Name": "Dan",
			"Occupation": map[string]interface{}{
				"Name": "Boston",
				"code": fmt.Sprintf(FORMAT_UNEXPORTED_STRING, "code", "uint"),
			},
			"Employee": map[string]interface{}{
				"department": fmt.Sprintf(FORMAT_UNEXPORTED_STRING, "department", "string"),
				"position":   fmt.Sprintf(FORMAT_UNEXPORTED_STRING, "position", "string"),
			},
		},
			nil},/*
		{struct {
			Public          string
			ZeroValueString string
			ZeroValueInt    int
			ZeroValueFloat  float32
			ZeroValueBool   bool

		//	IntPtr *int

			SomeNamedType NamedType
			NamedType

	//		ReadableChannel chan<- string
			//	WriteableChannel chan-> string

//			SomeFunc Comparator
		}{
			Public: "Field", SomeNamedType: 16, 
			//IntPtr: intPtrGetter(), 
			NamedType: 32,
		//	ReadableChannel: readableChannel,
	//		SomeFunc:        comparator,
		}, map[string]interface{}{
			"Public":          "Field",
			"ZeroValueString": "",
			"ZeroValueInt":    0,
			"ZeroValueFloat":  float32(0.0),
			"ZeroValueBool": false,

			"SomeNamedType":   16,
			"NamedType":       32,
	//		"IntPtr":        intPtrGetter(),

			//"ReadableChannel": readableChannel,

		//	"SomeFunc": reflect.ValueOf(comparator).Interface(),
		},
			nil},*/
		/*{struct {
			ErrField       error
			InterfaceField interface {
				SomeMethod(string) (string, error)
			}

			SomeStruct url.URL

			IntArr   [5]int
			IntSlice []int

			ObjArray []interface{}
		}{
			errors.New("Some error"),
			interfaceImpl,
			url.URL{
				Scheme:   "https",
				User:     url.UserPassword("admin", "password"),
				Host:     "github.com",
				Path:     "evityuk/dmpstruct",
				RawQuery: "x=1&y=2",
			},
			[5]int{1, 2, 4, 8, 16},
			[]int{1,2,3,5,8,13},
			[]interface{}{"string", 2, 3.14, true, nil},
		}, map[string]interface{}{
			"ErrField":       errors.New("Some error"),
			"InterfaceField": interfaceImpl,
			"SomeStruct": url.URL{
				Scheme:   "https",
				User:     url.UserPassword("admin", "password"),
				Host:     "github.com",
				Path:     "evityuk/dmpstruct",
				RawQuery: "x=1&y=2",
			},
			"IntArr":   [5]int{1, 2, 4, 8, 16},
			"IntSlice": []int{1,2,3,5,8,13},
			"ObjArray": []interface{}{"string", 2, 3.14, true, nil},
		}, nil},*/
		{
			nil, map[string]interface{}{}, errors.New("structObject isn't struct"), //nil implicitly converted to empty map[string]interface{}
		},
	}

	for _, c := range cases {
		got, err := Dump(c.in)
		/*		if err != c.err {
				t.Errorf("Dump(%q) == ",)
			}*/
	//		fmt.Println("Got: ", got, reflect.TypeOf(got))
		//	fmt.Println(reflect.DeepEqual(err, c.err), got, c.want)
		if !reflect.DeepEqual(err, c.err) /*(err != nil && err.Error() != c.err.Error())*/ || !reflect.DeepEqual(got, c.want) {
			t.Errorf("Dump(\n%q\n) == (\n%q, \n%q\n), want (\n%q, \n%q\n)", c.in, got, err, c.want, c.err)
		}
	}
}

func init() {
	Log.Level = logrus.InfoLevel
	Log.Out = os.Stdout
}

func constPtr(ptr int) func() *int {
	return func() *int {
		return &ptr
	}
}
