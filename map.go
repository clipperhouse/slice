package slice

import "github.com/clipperhouse/typewriter"

var mapT = &typewriter.Template{
	Name: "Map",
	Text: `
// Map{{.TypeParameter.LongName}} produces a new slice of values by mapping each value in list through a transformation function.
func (rcv {{.SliceName}}) Map{{.TypeParameter.LongName}}(fn func({{.Type}}) {{.TypeParameter}}) []{{.TypeParameter}} {
	numItems := len(rcv)
	result := make([]{{.TypeParameter}}, numItems)
	for i := 0; i < numItems; i++ {
		result[i] = fn(rcv[i])
	}
	return result
}
`,
	TypeParameterConstraints: []typewriter.Constraint{
		// exactly one type parameter is required, but no constraints on that type
		{},
	},
}
