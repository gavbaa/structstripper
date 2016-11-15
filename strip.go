package structstripper

import "reflect"

type StripConfig struct {
	Selectors []FieldSelector
}

type FieldSelector interface {
	IncludeField(stripConfig StripConfig, someStruct interface{}) (bool, []reflect.StructField)
}

func Strip(config StripConfig, someStruct interface{}) reflect.Type {
	newFields := []reflect.StructField{}

	for _, selector := range config.Selectors {
		res, addFields := selector.IncludeField(config, someStruct)
		if res {
			newFields = append(newFields, addFields...)
		}
	}

	return reflect.StructOf(newFields)
}

func NewSimpleNameConfig(names []string) StripConfig {
	fs := make([]FieldSelector, len(names), len(names))
	for i, name := range names {
		fs[i] = FieldBySimpleNameSelector{Name: name}
	}
	return StripConfig{
		Selectors: fs,
	}

}