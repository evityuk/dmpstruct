package dmpstruct

import (
	"errors"
	"fmt"
	"os"

	logrus "github.com/sirupsen/logrus"

	"net/url"
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

func TestDumpPositive(t *testing.T) {
	namedTypeValue := NamedType(3)
	structPtr := &(struct {
		FieldA string
		FieldB *OccupationInfo
		FieldC *NamedType
	}{
		FieldA: "someString", FieldB: &OccupationInfo{"Office", 6}, FieldC: &namedTypeValue,
	})

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
			"age":  fmt.Sprintf(FormatUnexportedString, "age", "int8"),
			"Name": "Dan",
			"Occupation": map[string]interface{}{
				"Name": "Boston",
				"code": fmt.Sprintf(FormatUnexportedString, "code", "uint"),
			},
			"Employee": map[string]interface{}{
				"department": fmt.Sprintf(FormatUnexportedString, "department", "string"),
				"position":   fmt.Sprintf(FormatUnexportedString, "position", "string"),
			},
		},
			nil},
		{struct {
			Public          string
			ZeroValueString string
			ZeroValueInt    int
			ZeroValueFloat  float32
			ZeroValueBool   bool

			IntPtr *int

			SomeNamedType NamedType
			NamedType

			Channel chan string

			//SomeFunc Comparator
		}{
			Public:        "Field",
			SomeNamedType: NamedType(16),
			IntPtr:        intPtrGetter(),
			NamedType:     NamedType(32),
			Channel:       rwChannel,
			//	SomeFunc:        comparator,
		}, map[string]interface{}{
			"ZeroValueBool": false,
			"SomeNamedType": NamedType(16),
			"NamedType":     NamedType(32),

			"Public":          "Field",
			"ZeroValueString": "",
			"ZeroValueInt":    0,
			"ZeroValueFloat":  float32(0.0),
			"IntPtr":          intPtrGetter(),

			"Channel": rwChannel,

			//"SomeFunc": reflect.ValueOf(comparator).Interface(),
		},
			nil},
		{struct {
			ErrField       error
			InterfaceField interface {
				SomeMethod(string) (string, error)
			}

			SomeStruct url.URL

			IntArr   [5]int
			IntSlice []int

			ObjArr   [2]interface{}
			ObjSlice []interface{}
		}{
			errInstance,
			interfaceImpl,
			url.URL{
				Scheme:   "https",
				User:     user,
				Host:     "github.com",
				Path:     "evityuk/dmpstruct",
				RawQuery: "x=1&y=2",
			},
			[5]int{1, 2, 4, 8, 16},
			[]int{1, 2, 3, 5, 8, 13},
			[2]interface{}{0.0, errInstance},
			[]interface{}{"string", 2, 3.14, true, nil, map[string]interface{}{"2": 4, "3": "8", "4th": 16.0}},
		}, map[string]interface{}{
			"ErrField":       errInstance,
			"InterfaceField": interfaceImpl,
			"SomeStruct": map[string]interface{}{
				"Fragment":   "",
				"Scheme":     "https",
				"Opaque":     "",
				"ForceQuery": false,
				"User": map[string]interface{}{
					"username":    fmt.Sprintf(FormatUnexportedString, "username", "string"),
					"password":    fmt.Sprintf(FormatUnexportedString, "password", "string"),
					"passwordSet": fmt.Sprintf(FormatUnexportedString, "passwordSet", "bool"),
				},
				"Host":     "github.com",
				"Path":     "evityuk/dmpstruct",
				"RawPath":  "",
				"RawQuery": "x=1&y=2",
			},
			"IntArr":   [5]int{1, 2, 4, 8, 16},
			"IntSlice": []int{1, 2, 3, 5, 8, 13},
			"ObjArr":   [2]interface{}{0.0, errInstance},
			"ObjSlice": []interface{}{"string", 2, 3.14, true, nil, map[string]interface{}{"2": 4, "3": "8", "4th": 16.0}},
		},
			nil},
		{
			structPtr, map[string]interface{}{
				"FieldA": "someString",
				"FieldB": map[string]interface{}{
					"Name": "Office",
					"code": fmt.Sprintf(FormatUnexportedString, "code", "uint"),
				},
				"FieldC": &namedTypeValue,
			}, nil,
		},
	}

	for _, c := range cases {
		got, err := Dump(c.in)

		if !reflect.DeepEqual(err, c.err) || !reflect.DeepEqual(got, c.want) {
			t.Errorf("Dump(\n%#v\n) == (\n%#v, \n%v\n), want (\n%#v, \n%v\n)", c.in, got, err, c.want, c.err)
		}
	}
}

func TestDumpNegative(t *testing.T) {
	cases := []struct {
		in, want interface{}
		err      error
	}{
		{
			nil, map[string]interface{}(nil), errors.New("structObject isn't struct or pointer to struct"),
		},
		{
			1, map[string]interface{}(nil), errors.New("structObject isn't struct or pointer to struct"),
		},
		{
			3.14, map[string]interface{}(nil), errors.New("structObject isn't struct or pointer to struct"),
		},
		{
			"string", map[string]interface{}(nil), errors.New("structObject isn't struct or pointer to struct"),
		},
		{
			[1]int{1}, map[string]interface{}(nil), errors.New("structObject isn't struct or pointer to struct"),
		},
		{
			&[]string{"string1", "string2", ""}, map[string]interface{}(nil), errors.New("structObject isn't struct or pointer to struct"),
		},
	}

	for _, c := range cases {
		got, err := Dump(c.in)
		if !reflect.DeepEqual(err, c.err) || !reflect.DeepEqual(got, c.want) {
			t.Errorf("Dump(\n%#v\n) == (\n%#v, \n%v\n), want (\n%#v, \n%v\n)", c.in, got, err, c.want, c.err)
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
