// Generated by: setup
// TypeWriter: slice
// Directive: +test on *PointerThing

package main

import (
	"errors"
	"math/rand"
)

// Sort implementation is a modification of http://golang.org/pkg/sort/#Sort
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found at http://golang.org/LICENSE.

// PointerThingSlice is a slice of type *PointerThing. Use it where you would use []*PointerThing.
type PointerThingSlice []*PointerThing

// Any verifies that one or more elements of PointerThingSlice return true for the passed func. See: http://clipperhouse.github.io/gen/#Any
func (rcv PointerThingSlice) Any(fn func(*PointerThing) bool) bool {
	for _, v := range rcv {
		if fn(v) {
			return true
		}
	}
	return false
}

// All verifies that all elements of PointerThingSlice return true for the passed func. See: http://clipperhouse.github.io/gen/#All
func (rcv PointerThingSlice) All(fn func(*PointerThing) bool) bool {
	for _, v := range rcv {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Count gives the number elements of PointerThingSlice that return true for the passed func. See: http://clipperhouse.github.io/gen/#Count
func (rcv PointerThingSlice) Count(fn func(*PointerThing) bool) (result int) {
	for _, v := range rcv {
		if fn(v) {
			result++
		}
	}
	return
}

// Distinct returns a new PointerThingSlice whose elements are unique. See: http://clipperhouse.github.io/gen/#Distinct
func (rcv PointerThingSlice) Distinct() (result PointerThingSlice) {
	appended := make(map[PointerThing]bool)
	for _, v := range rcv {
		if !appended[*v] {
			result = append(result, v)
			appended[*v] = true
		}
	}
	return result
}

// DistinctBy returns a new PointerThingSlice whose elements are unique, where equality is defined by a passed func. See: http://clipperhouse.github.io/gen/#DistinctBy
func (rcv PointerThingSlice) DistinctBy(equal func(*PointerThing, *PointerThing) bool) (result PointerThingSlice) {
Outer:
	for _, v := range rcv {
		for _, r := range result {
			if equal(v, r) {
				continue Outer
			}
		}
		result = append(result, v)
	}
	return result
}

// Each iterates over PointerThingSlice and executes the passed func against each element. See: http://clipperhouse.github.io/gen/#Each
func (rcv PointerThingSlice) Each(fn func(*PointerThing)) {
	for _, v := range rcv {
		fn(v)
	}
}

// First returns the first element that returns true for the passed func. Returns error if no elements return true. See: http://clipperhouse.github.io/gen/#First
func (rcv PointerThingSlice) First(fn func(*PointerThing) bool) (result *PointerThing, err error) {
	for _, v := range rcv {
		if fn(v) {
			result = v
			return
		}
	}
	err = errors.New("no PointerThingSlice elements return true for passed func")
	return
}

// MaxBy returns an element of PointerThingSlice containing the maximum value, when compared to other elements using a passed func defining ‘less’. In the case of multiple items being equally maximal, the last such element is returned. Returns error if no elements. See: http://clipperhouse.github.io/gen/#MaxBy
func (rcv PointerThingSlice) MaxBy(less func(*PointerThing, *PointerThing) bool) (result *PointerThing, err error) {
	l := len(rcv)
	if l == 0 {
		err = errors.New("cannot determine the MaxBy of an empty slice")
		return
	}
	m := 0
	for i := 1; i < l; i++ {
		if rcv[i] != rcv[m] && !less(rcv[i], rcv[m]) {
			m = i
		}
	}
	result = rcv[m]
	return
}

// MinBy returns an element of PointerThingSlice containing the minimum value, when compared to other elements using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such element is returned. Returns error if no elements. See: http://clipperhouse.github.io/gen/#MinBy
func (rcv PointerThingSlice) MinBy(less func(*PointerThing, *PointerThing) bool) (result *PointerThing, err error) {
	l := len(rcv)
	if l == 0 {
		err = errors.New("cannot determine the Min of an empty slice")
		return
	}
	m := 0
	for i := 1; i < l; i++ {
		if less(rcv[i], rcv[m]) {
			m = i
		}
	}
	result = rcv[m]
	return
}

// Single returns exactly one element of PointerThingSlice that returns true for the passed func. Returns error if no or multiple elements return true. See: http://clipperhouse.github.io/gen/#Single
func (rcv PointerThingSlice) Single(fn func(*PointerThing) bool) (result *PointerThing, err error) {
	var candidate *PointerThing
	found := false
	for _, v := range rcv {
		if fn(v) {
			if found {
				err = errors.New("multiple PointerThingSlice elements return true for passed func")
				return
			}
			candidate = v
			found = true
		}
	}
	if found {
		result = candidate
	} else {
		err = errors.New("no PointerThingSlice elements return true for passed func")
	}
	return
}

// Where returns a new PointerThingSlice whose elements return true for func. See: http://clipperhouse.github.io/gen/#Where
func (rcv PointerThingSlice) Where(fn func(*PointerThing) bool) (result PointerThingSlice) {
	for _, v := range rcv {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// SortBy returns a new ordered PointerThingSlice, determined by a func defining ‘less’. See: http://clipperhouse.github.io/gen/#SortBy
func (rcv PointerThingSlice) SortBy(less func(*PointerThing, *PointerThing) bool) PointerThingSlice {
	result := make(PointerThingSlice, len(rcv))
	copy(result, rcv)
	// Switch to heapsort if depth of 2*ceil(lg(n+1)) is reached.
	n := len(result)
	maxDepth := 0
	for i := n; i > 0; i >>= 1 {
		maxDepth++
	}
	maxDepth *= 2
	quickSortPointerThingSlice(result, less, 0, n, maxDepth)
	return result
}

// SortByDesc returns a new, descending-ordered PointerThingSlice, determined by a func defining ‘less’. See: http://clipperhouse.github.io/gen/#SortBy
func (rcv PointerThingSlice) SortByDesc(less func(*PointerThing, *PointerThing) bool) PointerThingSlice {
	greater := func(a, b *PointerThing) bool {
		return less(b, a)
	}
	return rcv.SortBy(greater)
}

// IsSortedBy reports whether an instance of PointerThingSlice is sorted, using the pass func to define ‘less’. See: http://clipperhouse.github.io/gen/#SortBy
func (rcv PointerThingSlice) IsSortedBy(less func(*PointerThing, *PointerThing) bool) bool {
	n := len(rcv)
	for i := n - 1; i > 0; i-- {
		if less(rcv[i], rcv[i-1]) {
			return false
		}
	}
	return true
}

// IsSortedDesc reports whether an instance of PointerThingSlice is sorted in descending order, using the pass func to define ‘less’. See: http://clipperhouse.github.io/gen/#SortBy
func (rcv PointerThingSlice) IsSortedByDesc(less func(*PointerThing, *PointerThing) bool) bool {
	greater := func(a, b *PointerThing) bool {
		return less(b, a)
	}
	return rcv.IsSortedBy(greater)
}

// AggregateOther iterates over PointerThingSlice, operating on each element while maintaining ‘state’. See: http://clipperhouse.github.io/gen/#Aggregate
func (rcv PointerThingSlice) AggregateOther(fn func(Other, *PointerThing) Other) (result Other) {
	for _, v := range rcv {
		result = fn(result, v)
	}
	return
}

// AverageOther sums Other over all elements and divides by len(PointerThingSlice). See: http://clipperhouse.github.io/gen/#Average
func (rcv PointerThingSlice) AverageOther(fn func(*PointerThing) Other) (result Other, err error) {
	l := len(rcv)
	if l == 0 {
		err = errors.New("cannot determine Average[Other] of zero-length PointerThingSlice")
		return
	}
	for _, v := range rcv {
		result += fn(v)
	}
	result = result / Other(l)
	return
}

// GroupByOther groups elements into a map keyed by Other. See: http://clipperhouse.github.io/gen/#GroupBy
func (rcv PointerThingSlice) GroupByOther(fn func(*PointerThing) Other) map[Other]PointerThingSlice {
	result := make(map[Other]PointerThingSlice)
	for _, v := range rcv {
		key := fn(v)
		result[key] = append(result[key], v)
	}
	return result
}

// MaxOther selects the largest value of Other in PointerThingSlice. Returns error on PointerThingSlice with no elements. See: http://clipperhouse.github.io/gen/#MaxCustom
func (rcv PointerThingSlice) MaxOther(fn func(*PointerThing) Other) (result Other, err error) {
	l := len(rcv)
	if l == 0 {
		err = errors.New("cannot determine Max of zero-length PointerThingSlice")
		return
	}
	result = fn(rcv[0])
	if l > 1 {
		for _, v := range rcv[1:] {
			f := fn(v)
			if f > result {
				result = f
			}
		}
	}
	return
}

// MinOther selects the least value of Other in PointerThingSlice. Returns error on PointerThingSlice with no elements. See: http://clipperhouse.github.io/gen/#MinCustom
func (rcv PointerThingSlice) MinOther(fn func(*PointerThing) Other) (result Other, err error) {
	l := len(rcv)
	if l == 0 {
		err = errors.New("cannot determine Min of zero-length PointerThingSlice")
		return
	}
	result = fn(rcv[0])
	if l > 1 {
		for _, v := range rcv[1:] {
			f := fn(v)
			if f < result {
				result = f
			}
		}
	}
	return
}

// SelectOther projects a slice of Other from PointerThingSlice, typically called a map in other frameworks. See: http://clipperhouse.github.io/gen/#Select
func (rcv PointerThingSlice) SelectOther(fn func(*PointerThing) Other) (result []Other) {
	for _, v := range rcv {
		result = append(result, fn(v))
	}
	return
}

// Shuffle returns a shuffled copy of PointerThingSlice, using a version of the Fisher-Yates shuffle. See: http://clipperhouse.github.io/gen/#Shuffle
func (rcv PointerThingSlice) Shuffle() PointerThingSlice {
	numItems := len(rcv)
	result := make(PointerThingSlice, numItems)
	copy(result, rcv)
	for i := 0; i < numItems; i++ {
		r := i + rand.Intn(numItems-i)
		result[r], result[i] = result[i], result[r]
	}
	return result
}

// SumOther sums *PointerThing over elements in PointerThingSlice. See: http://clipperhouse.github.io/gen/#Sum
func (rcv PointerThingSlice) SumOther(fn func(*PointerThing) Other) (result Other) {
	for _, v := range rcv {
		result += fn(v)
	}
	return
}

// Sort implementation based on http://golang.org/pkg/sort/#Sort, see top of this file

func swapPointerThingSlice(rcv PointerThingSlice, a, b int) {
	rcv[a], rcv[b] = rcv[b], rcv[a]
}

// Insertion sort
func insertionSortPointerThingSlice(rcv PointerThingSlice, less func(*PointerThing, *PointerThing) bool, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && less(rcv[j], rcv[j-1]); j-- {
			swapPointerThingSlice(rcv, j, j-1)
		}
	}
}

// siftDown implements the heap property on rcv[lo, hi).
// first is an offset into the array where the root of the heap lies.
func siftDownPointerThingSlice(rcv PointerThingSlice, less func(*PointerThing, *PointerThing) bool, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && less(rcv[first+child], rcv[first+child+1]) {
			child++
		}
		if !less(rcv[first+root], rcv[first+child]) {
			return
		}
		swapPointerThingSlice(rcv, first+root, first+child)
		root = child
	}
}

func heapSortPointerThingSlice(rcv PointerThingSlice, less func(*PointerThing, *PointerThing) bool, a, b int) {
	first := a
	lo := 0
	hi := b - a

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDownPointerThingSlice(rcv, less, i, hi, first)
	}

	// Pop elements, largest first, into end of rcv.
	for i := hi - 1; i >= 0; i-- {
		swapPointerThingSlice(rcv, first, first+i)
		siftDownPointerThingSlice(rcv, less, lo, i, first)
	}
}

// Quicksort, following Bentley and McIlroy,
// Engineering a Sort Function, SP&E November 1993.

// medianOfThree moves the median of the three values rcv[a], rcv[b], rcv[c] into rcv[a].
func medianOfThreePointerThingSlice(rcv PointerThingSlice, less func(*PointerThing, *PointerThing) bool, a, b, c int) {
	m0 := b
	m1 := a
	m2 := c
	// bubble sort on 3 elements
	if less(rcv[m1], rcv[m0]) {
		swapPointerThingSlice(rcv, m1, m0)
	}
	if less(rcv[m2], rcv[m1]) {
		swapPointerThingSlice(rcv, m2, m1)
	}
	if less(rcv[m1], rcv[m0]) {
		swapPointerThingSlice(rcv, m1, m0)
	}
	// now rcv[m0] <= rcv[m1] <= rcv[m2]
}

func swapRangePointerThingSlice(rcv PointerThingSlice, a, b, n int) {
	for i := 0; i < n; i++ {
		swapPointerThingSlice(rcv, a+i, b+i)
	}
}

func doPivotPointerThingSlice(rcv PointerThingSlice, less func(*PointerThing, *PointerThing) bool, lo, hi int) (midlo, midhi int) {
	m := lo + (hi-lo)/2 // Written like this to avoid integer overflow.
	if hi-lo > 40 {
		// Tukey's Ninther, median of three medians of three.
		s := (hi - lo) / 8
		medianOfThreePointerThingSlice(rcv, less, lo, lo+s, lo+2*s)
		medianOfThreePointerThingSlice(rcv, less, m, m-s, m+s)
		medianOfThreePointerThingSlice(rcv, less, hi-1, hi-1-s, hi-1-2*s)
	}
	medianOfThreePointerThingSlice(rcv, less, lo, m, hi-1)

	// Invariants are:
	//	rcv[lo] = pivot (set up by ChoosePivot)
	//	rcv[lo <= i < a] = pivot
	//	rcv[a <= i < b] < pivot
	//	rcv[b <= i < c] is unexamined
	//	rcv[c <= i < d] > pivot
	//	rcv[d <= i < hi] = pivot
	//
	// Once b meets c, can swap the "= pivot" sections
	// into the middle of the slice.
	pivot := lo
	a, b, c, d := lo+1, lo+1, hi, hi
	for {
		for b < c {
			if less(rcv[b], rcv[pivot]) { // rcv[b] < pivot
				b++
			} else if !less(rcv[pivot], rcv[b]) { // rcv[b] = pivot
				swapPointerThingSlice(rcv, a, b)
				a++
				b++
			} else {
				break
			}
		}
		for b < c {
			if less(rcv[pivot], rcv[c-1]) { // rcv[c-1] > pivot
				c--
			} else if !less(rcv[c-1], rcv[pivot]) { // rcv[c-1] = pivot
				swapPointerThingSlice(rcv, c-1, d-1)
				c--
				d--
			} else {
				break
			}
		}
		if b >= c {
			break
		}
		// rcv[b] > pivot; rcv[c-1] < pivot
		swapPointerThingSlice(rcv, b, c-1)
		b++
		c--
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	n := min(b-a, a-lo)
	swapRangePointerThingSlice(rcv, lo, b-n, n)

	n = min(hi-d, d-c)
	swapRangePointerThingSlice(rcv, c, hi-n, n)

	return lo + b - a, hi - (d - c)
}

func quickSortPointerThingSlice(rcv PointerThingSlice, less func(*PointerThing, *PointerThing) bool, a, b, maxDepth int) {
	for b-a > 7 {
		if maxDepth == 0 {
			heapSortPointerThingSlice(rcv, less, a, b)
			return
		}
		maxDepth--
		mlo, mhi := doPivotPointerThingSlice(rcv, less, a, b)
		// Avoiding recursion on the larger subproblem guarantees
		// a stack depth of at most lg(b-a).
		if mlo-a < b-mhi {
			quickSortPointerThingSlice(rcv, less, a, mlo, maxDepth)
			a = mhi // i.e., quickSortPointerThingSlice(rcv, mhi, b)
		} else {
			quickSortPointerThingSlice(rcv, less, mhi, b, maxDepth)
			b = mlo // i.e., quickSortPointerThingSlice(rcv, a, mlo)
		}
	}
	if b-a > 1 {
		insertionSortPointerThingSlice(rcv, less, a, b)
	}
}
