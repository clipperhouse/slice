package slicewriter

import (
	"io"
	"regexp"
	"strings"

	"github.com/clipperhouse/typewriter"
)

func init() {
	err := typewriter.Register(NewSliceWriter())
	if err != nil {
		panic(err)
	}
}

type cache struct {
	values []typewriter.TagValue
}

func SliceName(typ typewriter.Type) string {
	return typ.Name + "Slice"
}

type SliceWriter struct{}

func NewSliceWriter() *SliceWriter {
	return &SliceWriter{}
}

func (sw *SliceWriter) Name() string {
	return "slice"
}

func (sw *SliceWriter) Imports(typ typewriter.Type) (result []typewriter.ImportSpec) {
	// typewriter uses golang.org/x/tools/imports, depend on that
	return
}

func (sw *SliceWriter) Write(w io.Writer, typ typewriter.Type) error {
	tag, found, err := typ.Tags.ByName("slice")

	if err != nil {
		return err
	}

	if !found {
		return nil
	}

	if includeSortSupport(tag.Values) {
		s := `// Sort implementation is a modification of http://golang.org/pkg/sort/#Sort
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found at http://golang.org/LICENSE.

`
		w.Write([]byte(s))
	}

	tmpl, err := templates.Get(typ, typewriter.TagValue{Name: "slice"})

	if err != nil {
		return err
	}

	m := model{
		Type:      typ,
		SliceName: SliceName(typ),
	}

	if err := tmpl.Execute(w, m); err != nil {
		return err
	}

	for _, v := range tag.Values {
		var tp typewriter.Type

		if len(v.TypeParameters) > 0 {
			tp = v.TypeParameters[0]
		}

		m := model{
			Type:          typ,
			SliceName:     SliceName(typ),
			TypeParameter: tp,
			TagValue:      v,
		}

		tmpl, err := templates.Get(typ, v) // already validated above

		if err != nil {
			return err
		}

		if err := tmpl.Execute(w, m); err != nil {
			return err
		}
	}

	if includeSortInterface(tag.Values) {
		tmpl, err := templates.Get(typ, typewriter.TagValue{Name: "sortInterface"})

		if err != nil {
			return err
		}

		if err := tmpl.Execute(w, m); err != nil {
			return err
		}
	}

	if includeSortSupport(tag.Values) {
		tmpl, err := templates.Get(typ, typewriter.TagValue{Name: "sortImplementation"})

		if err != nil {
			return err
		}

		if err := tmpl.Execute(w, m); err != nil {
			return err
		}
	}

	return nil
}

func includeSortSupport(values []typewriter.TagValue) bool {
	for _, v := range values {
		if strings.HasPrefix(v.Name, "SortBy") {
			return true
		}
	}
	return false
}

func includeSortInterface(values []typewriter.TagValue) bool {
	reg := regexp.MustCompile(`^Sort(Desc)?$`)
	for _, v := range values {
		if reg.MatchString(v.Name) {
			return true
		}
	}
	return false
}
