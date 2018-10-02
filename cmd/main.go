package main

import (
	"fmt"
	"yvitiuk/dmpstruct"
)

type Employee struct {
	department string
	position   string
}

type OccupationInfo struct {
	Name string
	code uint
}

type S struct {
	Name       string
	age        int8
	Occupation OccupationInfo
	Employee
}

func main() {
	s := S{"Dan", 50, OccupationInfo{"Boston", 3}, Employee{"Literature", "Writer"}}

	dump, err := dmpstruct.Dump(s)
	if err != nil {
		fmt.Println("Dump error: ", err)
	} else {
		fmt.Printf("Dumped: %q", dump)
	}

}
