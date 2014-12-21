package slice

import "github.com/clipperhouse/typewriter"

var sortBy = &typewriter.Template{
	Name: "SortBy",
	Text: `
// SortBy returns a new ordered {{.SliceName}}, determined by a func defining ‘less’. See: http://clipperhouse.github.io/gen/#SortBy
func (rcv {{.SliceName}}) SortBy(less func({{.Type}}, {{.Type}}) bool) {{.SliceName}} {
	return rcv.sortBy(less)
}
`}

var isSortedBy = &typewriter.Template{
	Name: "IsSortedBy",
	Text: `
// IsSortedBy reports whether an instance of {{.SliceName}} is sorted, using the pass func to define ‘less’. See: http://clipperhouse.github.io/gen/#SortBy
func (rcv {{.SliceName}}) IsSortedBy(less func({{.Type}}, {{.Type}}) bool) bool {
	return rcv.isSortedBy(less)
}
`}

var sortByDesc = &typewriter.Template{
	Name: "SortByDesc",
	Text: `
// SortByDesc returns a new, descending-ordered {{.SliceName}}, determined by a func defining ‘less’. See: http://clipperhouse.github.io/gen/#SortBy
func (rcv {{.SliceName}}) SortByDesc(less func({{.Type}}, {{.Type}}) bool) {{.SliceName}} {
	greater := func(a, b {{.Type}}) bool {
		return less(b, a)
	}
	return rcv.sortBy(greater)
}
`}

var isSortedByDesc = &typewriter.Template{
	Name: "IsSortedByDesc",
	Text: `
// IsSortedDesc reports whether an instance of {{.SliceName}} is sorted in descending order, using the pass func to define ‘less’. See: http://clipperhouse.github.io/gen/#SortBy
func (rcv {{.SliceName}}) IsSortedByDesc(less func({{.Type}}, {{.Type}}) bool) bool {
	greater := func(a, b {{.Type}}) bool {
		return less(b, a)
	}
	return rcv.isSortedBy(greater)
}
`}
