package dmpstruct

import (
	"errors"
	"fmt"
	"reflect"
)

func Dump(structObject interface{}) (map[string]interface{}, error) {
//	fmt.Println("Dumping...", structObject)
	if reflect.TypeOf(structObject).Kind() != reflect.Struct {
		return nil, errors.New("structObject isn't struct")
	}

	valOf := reflect.ValueOf(structObject)
	sType := valOf.Type()
	nFields := valOf.NumField()

	results := make(map[string]interface{})

	for i := 0; i < nFields; i++ {
		field := valOf.Field(i)
		fName := sType.Field(i).Name
//		fmt.Println("kind:", field.Kind().String(), "name:", fName, "field:", field)
		if field.CanInterface() {
			if field.Kind() == reflect.Struct {
				if dump, err := Dump(field.Interface()); err != nil {
					results[fName] = fmt.Sprintf("Error dumping field '%s' of type '%s': '%s'", sType.Field(i).Name, field.Kind().String(), err)
				} else {
					results[fName] = dump
				}
			} else {
				results[fName] = field.Interface()
			}
		} else {
			//			results[field.Kind().String()] = nil
			results[fName] = fmt.Sprintf("Field '%s' of type '%s' is unexported", sType.Field(i).Name, field.Kind().String())
		}
	}

	return results, nil
}
