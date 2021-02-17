package main

import "testing"

func TestShuffle(t *testing.T) {
	original := ThingSlice{
		{"First", 60},
		{"Second", -20},
		{"Third", 100},
	}

	shuffled := original.Shuffle()

	notFound := func(t Thing) bool {
		for x := 0; x < len(shuffled); x++ {
			if shuffled[x] == t {
				return false
			}
		}
		return true
	}

	if original.Any(notFound) {
		t.Error("The shuffled slice is missing elements")
	}
}

func TestPointerShuffle(t *testing.T) {
	original := PointerThingSlice{
		{"First", 60},
		{"Second", -20},
		{"Third", 100},
	}

	shuffled := original.Shuffle()

	notFound := func(t *PointerThing) bool {
		for x := 0; x < len(shuffled); x++ {
			if shuffled[x] == t {
				return false
			}
		}
		return true
	}

	if original.Any(notFound) {
		t.Error("The shuffled slice is missing elements")
	}
}
