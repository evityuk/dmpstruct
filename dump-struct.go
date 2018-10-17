//Package dmpstruct is a GoLang try by writing utility package and cmd for dumping structs to map[string]interface{}
package dmpstruct

import (
	"errors"
	"io"

	logrus "github.com/sirupsen/logrus"

	"fmt"
	"os"
	"reflect"
)

const (
	//FormatErrString is a fmt.Printf like dump err format string
	FormatErrString        = "Error dumping field %q of type %q: %q"
	//FormatUnexportedString is fmt.Printf like unexported struct field replacement format string
	FormatUnexportedString = "Field %q of type %q is unexported"
)

//Log is global logger instance
var Log = logrus.New()

func init() {
	Log.Out = os.Stderr
	Log.Formatter = &logrus.TextFormatter{}
	Log.Level = logrus.FatalLevel
}

//Init func initializes logging output, formatter and level globally per package
func Init(debugWriter io.Writer, formatter logrus.Formatter, level logrus.Level) {
	Log.Out = debugWriter
	Log.Formatter = formatter
	Log.Level = level
}

//Dump func returns converted recursively to map[string]interface{} struct object passed to it
//Each struct field represents key in returned map, the value is left as is.
//It returns error if passed object isn't a struct of pointer to struct
func Dump(structObject interface{}) (map[string]interface{}, error) {
	Log.Debugln("Dumping object", structObject)

	if structObject == nil || (reflect.ValueOf(structObject).Kind() == reflect.Ptr && reflect.ValueOf(structObject).Elem().Kind() != reflect.Struct) || (reflect.TypeOf(structObject).Kind() != reflect.Struct && reflect.TypeOf(structObject).Kind() != reflect.Ptr) {
		return nil, errors.New("structObject isn't struct or pointer to struct")
	}

	valOf := reflect.ValueOf(structObject)
	if reflect.TypeOf(structObject).Kind() == reflect.Ptr {
		valOf = valOf.Elem()
	}
	sType := valOf.Type()
	nFields := valOf.NumField()

	results := make(map[string]interface{})

	for i := 0; i < nFields; i++ {
		field := valOf.Field(i)
		fName := sType.Field(i).Name

		Log.Debugf("Dumping field: kind: %v; name: %v; value: %v; typeInfo: %#v; valueInfo: %#v\n", field.Kind().String(), fName, field, sType.Field(i), field)
		if field.CanInterface() {
			var isStructPointer = field.Kind() == reflect.Ptr && field.Elem().Kind() == reflect.Struct

			if field.Kind() == reflect.Struct || isStructPointer {
				var fieldStructObject = field.Interface()
				if isStructPointer {
					fieldStructObject = field.Elem().Interface()
				}

				if dump, err := Dump(fieldStructObject); err != nil {
					results[fName] = fmt.Sprintf(FormatErrString, sType.Field(i).Name, field.Kind().String(), err)
				} else {
					results[fName] = dump
				}
			} else {
				results[fName] = field.Interface()
			}
		} else {
			//			results[field.Kind().String()] = nil
			results[fName] = fmt.Sprintf(FormatUnexportedString, sType.Field(i).Name, field.Kind().String())
		}
	}

	return results, nil
}
