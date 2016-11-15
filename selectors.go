package structstripper

import "reflect"

type FieldBySimpleNameSelector struct {
	Name string
}

func (f FieldBySimpleNameSelector) IncludeField(stripConfig StripConfig, someStruct interface{}) (bool, []reflect.StructField) {
	sElem := reflect.TypeOf(someStruct)
	for fieldIdx := 0; fieldIdx < sElem.NumField(); fieldIdx += 1 {
		sField := sElem.Field(fieldIdx)
		if sField.Name == f.Name {
			return true, []reflect.StructField{sField}
		}
	}
	return false, nil
}

type FieldBySimpleTagSelector struct {
	Tag string
	Value string
}

func (f FieldBySimpleTagSelector) IncludeField(stripConfig StripConfig, someStruct interface{}) (bool, []reflect.StructField) {
	sElem := reflect.TypeOf(someStruct)
	for fieldIdx := 0; fieldIdx < sElem.NumField(); fieldIdx += 1 {
		sField := sElem.Field(fieldIdx)
		if value, ok := sField.Tag.Lookup(f.Tag); ok {
			if value == f.Value {
				return true, []reflect.StructField{sField}
			}
		}
	}
	return false, nil
}

