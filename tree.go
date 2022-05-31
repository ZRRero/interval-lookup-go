package main

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

type Tree[T constraints.Ordered] struct {
	root *leaf[T]
}

func (receiver *Tree[T]) SearchMultipleIntervals(value T, includeInit, includeEnd bool) ([]IntervalWithValue[T], error) {
	if &receiver.root == nil {
		return nil, errorRootIsNil{}
	}
	return receiver.root.SearchMultipleIntervals(make([]IntervalWithValue[T], 0), value, includeInit, includeEnd)
}

func (receiver *Tree[T]) SearchFirstInterval(value T, includeInit, includeEnd bool) (IntervalWithValue[T], error) {
	if &receiver.root == nil {
		return IntervalWithValue[T]{}, errorRootIsNil{}
	}
	return receiver.root.SearchFirstInterval(value, includeInit, includeEnd)
}

func GenerateTree[T constraints.Ordered](intervals IntervalsWithValue[T], searchOverlapping bool) (Tree[T], error) {
	slices.SortFunc(intervals, func(i, j IntervalWithValue[T]) bool {
		return i.Compare(j) <= 0
	})
	if searchOverlapping && intervals.detectOverlapping() {
		return Tree[T]{}, errorIntervalsOverlapped{}
	}
	root := GenerateLeaf(intervals, 0, len(intervals))
	return Tree[T]{root: root}, nil
}
