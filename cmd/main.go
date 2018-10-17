package main

import (
	logrus "github.com/sirupsen/logrus"
	"os"
	"evityuk/dmpstruct"
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
	
	dmpstruct.Init(os.Stdout, &logrus.TextFormatter{}, logrus.DebugLevel)

	dump, err := dmpstruct.Dump(&s)
	if err != nil {
		logrus.Println("Dump error: ", err)
	} else {
		logrus.WithFields(dump).Println("Dumped successfully: \n")
	}


}
