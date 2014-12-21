package slice

import "github.com/clipperhouse/typewriter"

var sort = &typewriter.Template{
	Name: "Sort",
	Text: `
// Sort returns a new ordered {{.SliceName}}. See: http://clipperhouse.github.io/gen/#Sort
func (rcv {{.SliceName}}) Sort() {{.SliceName}} {
	less := func(a, b {{.Type}}) bool {
		return a < b
	}
	return rcv.sortBy(less)
}
`,
	TypeConstraint: typewriter.Constraint{Ordered: true},
}

var isSorted = &typewriter.Template{
	Name: "IsSorted",
	Text: `
// IsSorted reports whether {{.SliceName}} is sorted. See: http://clipperhouse.github.io/gen/#Sort
func (rcv {{.SliceName}}) IsSorted() bool {
	less := func(a, b {{.Type}}) bool {
		return a < b
	}
	return rcv.isSortedBy(less)
}
`,
	TypeConstraint: typewriter.Constraint{Ordered: true},
}

var sortDesc = &typewriter.Template{
	Name: "SortDesc",
	Text: `
// SortDesc returns a new reverse-ordered {{.SliceName}}. See: http://clipperhouse.github.io/gen/#Sort
func (rcv {{.SliceName}}) SortDesc() {{.SliceName}} {
	greater := func(a, b {{.Type}}) bool {
		return b < a
	}
	return rcv.sortBy(greater)
}
`,
	TypeConstraint: typewriter.Constraint{Ordered: true},
}

var isSortedDesc = &typewriter.Template{
	Name: "IsSortedDesc",
	Text: `
// IsSortedDesc reports whether {{.SliceName}} is reverse-sorted. See: http://clipperhouse.github.io/gen/#Sort
func (rcv {{.SliceName}}) IsSortedDesc() bool {
	greater := func(a, b {{.Type}}) bool {
		return b < a
	}
	return rcv.isSortedBy(greater)
}
`,
	TypeConstraint: typewriter.Constraint{Ordered: true},
}
