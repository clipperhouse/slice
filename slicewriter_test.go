package slicewriter

import (
	"bytes"
	"fmt"
	"go/parser"
	"go/token"
	"testing"

	"github.com/clipperhouse/typewriter"
)

var pkg *typewriter.Package

func init() {
	pkg = typewriter.NewPackage("dummy", "SomePackage")

	t1, err := pkg.Eval("int")

	if err != nil {
		panic(err)
	}

	t2, err := pkg.Eval("*rune")

	if err != nil {
		panic(err)
	}

	t1.Tags = typewriter.TagSlice{
		typewriter.Tag{
			Name: "slice",
			Values: []typewriter.TagValue{
				{"GroupBy", []typewriter.Type{t2}},
				{"Where", nil},
			},
		},
	}

	pkg.Types = append(pkg.Types, t1)
}

func TestWrite(t *testing.T) {
	for _, typ := range pkg.Types {
		var b bytes.Buffer

		sw := NewSliceWriter()

		b.WriteString(fmt.Sprintf("package %s\n\n", pkg.Name()))
		sw.Write(&b, typ)

		src := b.String()

		fset := token.NewFileSet()
		if _, err := parser.ParseFile(fset, "testwrite.go", src, 0); err != nil {
			t.Error(err)
		}
	}
}
