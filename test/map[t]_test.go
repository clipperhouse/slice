package main

import (
	"reflect"
	"testing"
)

func TestMapThing(t *testing.T) {
	things := ThingSlice{
		{"First", 60},
		{"Second", -20},
		{"Third", 100},
	}

	mapped := things.MapOther(func(t Thing) Other {
		return t.Number
	})

	expected := []Other{60, -20, 100}

	if !reflect.DeepEqual(mapped, expected) {
		t.Error("The slice is not mapped")
	}
}
