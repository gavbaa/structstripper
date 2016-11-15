package structstripper_test

import (
	"github.com/gavbaa/structstripper"
	"testing"
)

type Point1 struct {
	X int `json:"bob"`
	Y int `bson:"t1" json:",omitempty"`
}

func TestSimpleNameRemoveAll(t *testing.T) {
	sc := structstripper.StripConfig{
	}
	newType := structstripper.Strip(sc, Point1{})
	if newType.NumField() != 0 {
		t.Error("NumField should be 0.")
	}
}

func TestSimpleNameRemoveSome(t *testing.T) {
	sc := structstripper.StripConfig{
		Selectors: []structstripper.FieldSelector{
			structstripper.FieldBySimpleNameSelector{Name: "X"},
		},
	}
	newType := structstripper.Strip(sc, Point1{})
	if newType.NumField() != 1 {
		t.Error("NumField should be 1.")
	}
}

func TestSimpleNameRemoveNone(t *testing.T) {
	sc := structstripper.StripConfig{
		Selectors: []structstripper.FieldSelector{
			structstripper.FieldBySimpleNameSelector{Name: "X"},
			structstripper.FieldBySimpleNameSelector{Name: "Y"},
		},
	}
	newType := structstripper.Strip(sc, Point1{})
	if newType.NumField() != 2 {
		t.Error("NumField should be 2.")
	}
}

func TestTagPreserve(t *testing.T) {
	sc := structstripper.StripConfig{
		Selectors: []structstripper.FieldSelector{
			structstripper.FieldBySimpleNameSelector{Name: "X"},
			structstripper.FieldBySimpleNameSelector{Name: "Y"},
		},
	}
	newType := structstripper.Strip(sc, Point1{})
	if newType.NumField() != 2 {
		t.Error("NumField should be 2.")
	}
	if newType.Field(0).Tag != `json:"bob"` {
		t.Error("Field 0's tag does not match.")
	}
	if newType.Field(1).Tag != `bson:"t1" json:",omitempty"` {
		t.Error("Field 1's tag does not match.")
	}
}

func TestSimpleTagFirstKeyMatches(t *testing.T) {
	sc := structstripper.StripConfig{
		Selectors: []structstripper.FieldSelector{
			structstripper.FieldBySimpleTagSelector{Tag: "json", Value: "bob"},
			structstripper.FieldBySimpleTagSelector{Tag: "bson", Value: "t1"},
		},
	}
	newType := structstripper.Strip(sc, Point1{})
	if newType.NumField() != 2 {
		t.Error("NumField should be 2.")
	}
}

func TestSimpleTagLastKeyMatches(t *testing.T) {
	sc := structstripper.StripConfig{
		Selectors: []structstripper.FieldSelector{
			structstripper.FieldBySimpleTagSelector{Tag: "json", Value: ",omitempty"},
		},
	}
	newType := structstripper.Strip(sc, Point1{})
	if newType.NumField() != 1 {
		t.Error("NumField should be 1.")
	}
}

func TestSimpleTagMisses(t *testing.T) {
	sc := structstripper.StripConfig{
		Selectors: []structstripper.FieldSelector{
			structstripper.FieldBySimpleTagSelector{Tag: "json", Value: "nope"},
			structstripper.FieldBySimpleTagSelector{Tag: "bson", Value: "also_nope"},
		},
	}
	newType := structstripper.Strip(sc, Point1{})
	if newType.NumField() != 0 {
		t.Error("NumField should be 0.")
	}
}

func TestSimpleNameConfigHelper(t *testing.T) {
	sc := structstripper.NewSimpleNameConfig([]string{"X", "Y"})
	newType := structstripper.Strip(sc, Point1{})
	if newType.NumField() != 2 {
		t.Error("NumField should be 2.")
	}
}