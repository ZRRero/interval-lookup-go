package main

import (
	"golang.org/x/exp/constraints"
)

type IntervalWithValue[T constraints.Ordered] struct {
	init  T
	end   T
	value any
}

func GenerateIntervalWithValue[T constraints.Ordered](init, end T, value any) (IntervalWithValue[T], error) {
	interval := IntervalWithValue[T]{init: init, end: end, value: value}
	if err := interval.validate(); err != nil {
		return IntervalWithValue[T]{}, err
	}
	return interval, nil
}

func (receiver *IntervalWithValue[T]) validate() error {
	if receiver.init > receiver.end {
		return errorInvalidInterval[T]{receiver.init, receiver.end}
	}
	return nil
}

func (receiver *IntervalWithValue[T]) Compare(other IntervalWithValue[T]) int {
	if receiver.init == other.init && receiver.end == other.end {
		return 0
	}
	if receiver.init > other.init {
		return 1
	}
	if receiver.end < other.end {
		return -1
	}
	return 0
}

func (receiver *IntervalWithValue[T]) IsOverlapped(other IntervalWithValue[T]) bool {
	return !(receiver.init >= other.end || receiver.end <= other.init)
}

func (receiver *IntervalWithValue[T]) Evaluate(toEvaluate T, includeInit, includeEnd bool) string {
	if includeInit {
		if toEvaluate < receiver.init {
			return "Under"
		}
	} else {
		if toEvaluate <= receiver.init {
			return "Under"
		}
	}
	if includeEnd {
		if toEvaluate > receiver.end {
			return "Upper"
		}
	} else {
		if toEvaluate >= receiver.end {
			return "Upper"
		}
	}
	return "Contained"
}
