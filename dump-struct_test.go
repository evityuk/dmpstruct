package dmpstruct

import (
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
			"age":  "Field 'age' of type 'int8' is unexported",
			"Name": "Dan",
			"Occupation": map[string]interface{}{
				"Name": "Boston",
				"code": "Field 'code' of type 'uint' is unexported",
			},
			"Employee": map[string]interface{}{
				"department": "Field 'department' of type 'string' is unexported",
				"position":   "Field 'position' of type 'string' is unexported",
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
