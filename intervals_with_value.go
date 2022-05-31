package main

import "golang.org/x/exp/constraints"

type IntervalsWithValue[T constraints.Ordered] []IntervalWithValue[T]

func (intervals IntervalsWithValue[T]) detectOverlapping() bool {
	compare := intervals[0]
	for index, value := range intervals {
		if index != 0 && compare.IsOverlapped(&value) {
			return true
		}
		compare = value
	}
	return false
}
