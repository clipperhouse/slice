package slice

import "github.com/clipperhouse/typewriter"

var distinct = &typewriter.Template{
	Name: "Distinct",
	Text: `
// Distinct returns a new {{.SliceName}} whose elements are unique. See: http://clipperhouse.github.io/gen/#Distinct
func (rcv {{.SliceName}}) Distinct() (result {{.SliceName}}) {
	appended := make(map[{{.Type.Name}}]bool)
	for _, v := range rcv {
		{{ if .Type.Pointer -}}
			if !appended[*v] {
				result = append(result, v)
				appended[*v] = true
			}
		{{ else -}}
			if !appended[v] {
				result = append(result, v)
				appended[v] = true
			}
		{{ end -}}
	}
	return result
}
`,
	TypeConstraint: typewriter.Constraint{Comparable: true},
}
