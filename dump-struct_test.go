package dmpstruct

import (
	"reflect"
	"fmt"
	"testing"
	logrus "github.com/sirupsen/logrus"
)

type Employee struct {
	department string
	position   string
}

type OccupationInfo struct {
	Name string
	code uint
}

func init() {
	Log.Level = logrus.InfoLevel
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
			nil},
	}

	for _, c := range cases {
		got, err := Dump(c.in)
		/*		if err != c.err {
				t.Errorf("Dump(%q) == ",)
			}*/
		if err != c.err || !reflect.DeepEqual(got, c.want) {
			t.Errorf("Dump(\n%q\n) == (\n%q, \n%q\n), want \n%q\n", c.in, got, err, c.want)
		}
	}
}
