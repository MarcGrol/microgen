// Generated by: gen
// TypeWriter: slice
// Directive: +gen on Cyclist

package gambler

// Sort implementation is a modification of http://golang.org/pkg/sort/#Sort
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found at http://golang.org/LICENSE.

// CyclistSlice is a slice of type Cyclist. Use it where you would use []Cyclist.
type CyclistSlice []Cyclist

// SortBy returns a new ordered CyclistSlice, determined by a func defining ‘less’. See: http://clipperhouse.github.io/gen/#SortBy
func (rcv CyclistSlice) SortBy(less func(Cyclist, Cyclist) bool) CyclistSlice {
	result := make(CyclistSlice, len(rcv))
	copy(result, rcv)
	// Switch to heapsort if depth of 2*ceil(lg(n+1)) is reached.
	n := len(result)
	maxDepth := 0
	for i := n; i > 0; i >>= 1 {
		maxDepth++
	}
	maxDepth *= 2
	quickSortCyclistSlice(result, less, 0, n, maxDepth)
	return result
}

// Where returns a new CyclistSlice whose elements return true for func. See: http://clipperhouse.github.io/gen/#Where
func (rcv CyclistSlice) Where(fn func(Cyclist) bool) (result CyclistSlice) {
	for _, v := range rcv {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// SelectString projects a slice of string from CyclistSlice, typically called a map in other frameworks. See: http://clipperhouse.github.io/gen/#Select
func (rcv CyclistSlice) SelectString(fn func(Cyclist) string) (result []string) {
	for _, v := range rcv {
		result = append(result, fn(v))
	}
	return
}

// GroupByString groups elements into a map keyed by string. See: http://clipperhouse.github.io/gen/#GroupBy
func (rcv CyclistSlice) GroupByString(fn func(Cyclist) string) map[string]CyclistSlice {
	result := make(map[string]CyclistSlice)
	for _, v := range rcv {
		key := fn(v)
		result[key] = append(result[key], v)
	}
	return result
}

// Sort implementation based on http://golang.org/pkg/sort/#Sort, see top of this file

func swapCyclistSlice(rcv CyclistSlice, a, b int) {
	rcv[a], rcv[b] = rcv[b], rcv[a]
}

// Insertion sort
func insertionSortCyclistSlice(rcv CyclistSlice, less func(Cyclist, Cyclist) bool, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && less(rcv[j], rcv[j-1]); j-- {
			swapCyclistSlice(rcv, j, j-1)
		}
	}
}

// siftDown implements the heap property on rcv[lo, hi).
// first is an offset into the array where the root of the heap lies.
func siftDownCyclistSlice(rcv CyclistSlice, less func(Cyclist, Cyclist) bool, lo, hi, first int) {
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
		swapCyclistSlice(rcv, first+root, first+child)
		root = child
	}
}

func heapSortCyclistSlice(rcv CyclistSlice, less func(Cyclist, Cyclist) bool, a, b int) {
	first := a
	lo := 0
	hi := b - a

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDownCyclistSlice(rcv, less, i, hi, first)
	}

	// Pop elements, largest first, into end of rcv.
	for i := hi - 1; i >= 0; i-- {
		swapCyclistSlice(rcv, first, first+i)
		siftDownCyclistSlice(rcv, less, lo, i, first)
	}
}

// Quicksort, following Bentley and McIlroy,
// Engineering a Sort Function, SP&E November 1993.

// medianOfThree moves the median of the three values rcv[a], rcv[b], rcv[c] into rcv[a].
func medianOfThreeCyclistSlice(rcv CyclistSlice, less func(Cyclist, Cyclist) bool, a, b, c int) {
	m0 := b
	m1 := a
	m2 := c
	// bubble sort on 3 elements
	if less(rcv[m1], rcv[m0]) {
		swapCyclistSlice(rcv, m1, m0)
	}
	if less(rcv[m2], rcv[m1]) {
		swapCyclistSlice(rcv, m2, m1)
	}
	if less(rcv[m1], rcv[m0]) {
		swapCyclistSlice(rcv, m1, m0)
	}
	// now rcv[m0] <= rcv[m1] <= rcv[m2]
}

func swapRangeCyclistSlice(rcv CyclistSlice, a, b, n int) {
	for i := 0; i < n; i++ {
		swapCyclistSlice(rcv, a+i, b+i)
	}
}

func doPivotCyclistSlice(rcv CyclistSlice, less func(Cyclist, Cyclist) bool, lo, hi int) (midlo, midhi int) {
	m := lo + (hi-lo)/2 // Written like this to avoid integer overflow.
	if hi-lo > 40 {
		// Tukey's Ninther, median of three medians of three.
		s := (hi - lo) / 8
		medianOfThreeCyclistSlice(rcv, less, lo, lo+s, lo+2*s)
		medianOfThreeCyclistSlice(rcv, less, m, m-s, m+s)
		medianOfThreeCyclistSlice(rcv, less, hi-1, hi-1-s, hi-1-2*s)
	}
	medianOfThreeCyclistSlice(rcv, less, lo, m, hi-1)

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
				swapCyclistSlice(rcv, a, b)
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
				swapCyclistSlice(rcv, c-1, d-1)
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
		swapCyclistSlice(rcv, b, c-1)
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
	swapRangeCyclistSlice(rcv, lo, b-n, n)

	n = min(hi-d, d-c)
	swapRangeCyclistSlice(rcv, c, hi-n, n)

	return lo + b - a, hi - (d - c)
}

func quickSortCyclistSlice(rcv CyclistSlice, less func(Cyclist, Cyclist) bool, a, b, maxDepth int) {
	for b-a > 7 {
		if maxDepth == 0 {
			heapSortCyclistSlice(rcv, less, a, b)
			return
		}
		maxDepth--
		mlo, mhi := doPivotCyclistSlice(rcv, less, a, b)
		// Avoiding recursion on the larger subproblem guarantees
		// a stack depth of at most lg(b-a).
		if mlo-a < b-mhi {
			quickSortCyclistSlice(rcv, less, a, mlo, maxDepth)
			a = mhi // i.e., quickSortCyclistSlice(rcv, mhi, b)
		} else {
			quickSortCyclistSlice(rcv, less, mhi, b, maxDepth)
			b = mlo // i.e., quickSortCyclistSlice(rcv, a, mlo)
		}
	}
	if b-a > 1 {
		insertionSortCyclistSlice(rcv, less, a, b)
	}
}
