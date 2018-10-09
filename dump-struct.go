//Go try by writing utility package and cmd for dumping structs to map[string]interface{}
package dmpstruct

import (
	"errors"
	"io"

	logrus "github.com/sirupsen/logrus"

	"fmt"
	"reflect"
)

const (
	FORMAT_ERR_STRING = "Error dumping field %q of type %q: %q"
	FORMAT_UNEXPORTED_STRING = "Field %q of type %q is unexported"
)

var Log = logrus.New()

func Init(debugWriter io.Writer, formatter logrus.Formatter, level logrus.Level) {
	Log.Out = debugWriter
	Log.Formatter = formatter
	Log.Level = level
}

func Dump(structObject interface{}) (map[string]interface{}, error) {
	Log.Debugln("Dumping...", structObject)
	if structObject == nil || reflect.TypeOf(structObject).Kind() != reflect.Struct {
//		fmt.Println("error!")
		return make(map[string]interface{})/*nil*/, errors.New("structObject isn't struct")
	}

	valOf := reflect.ValueOf(structObject)
	sType := valOf.Type()
	nFields := valOf.NumField()

	results := make(map[string]interface{})

	for i := 0; i < nFields; i++ {
		field := valOf.Field(i)
		fName := sType.Field(i).Name
		Log.Debugln("kind:", field.Kind().String(), "name:", fName, "field:", field)
		if field.CanInterface() {
			if field.Kind() == reflect.Struct {
				if dump, err := Dump(field.Interface()); err != nil {
					results[fName] = fmt.Sprintf(FORMAT_ERR_STRING, sType.Field(i).Name, field.Kind().String(), err)
				} else {
					results[fName] = dump
				}
			} else {
				results[fName] = field.Interface()
			}
		} else {
			//			results[field.Kind().String()] = nil
			results[fName] = fmt.Sprintf(FORMAT_UNEXPORTED_STRING, sType.Field(i).Name, field.Kind().String())
		}
	}

	return results, nil
}
