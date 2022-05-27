package main

import "golang.org/x/exp/constraints"

type leaf[T constraints.Ordered] struct {
	interval IntervalWithValue[T]
	left     *leaf[T]
	right    *leaf[T]
}

func (receiver *leaf[T]) SearchMultipleIntervals(foundIntervals []IntervalWithValue[T], value T, includeInit, includeEnd bool) ([]IntervalWithValue[T], error) {
	evaluation := receiver.interval.Evaluate(value, includeInit, includeEnd)
	if foundIntervals == nil {
		foundIntervals = make([]IntervalWithValue[T], 0)
	}
	switch evaluation {
	case "Under":
		if receiver.left != nil {
			foundIntervals, _ = receiver.left.SearchMultipleIntervals(foundIntervals, value, includeInit, includeEnd)
		}
	case "Upper":
		if receiver.right != nil {
			foundIntervals, _ = receiver.right.SearchMultipleIntervals(foundIntervals, value, includeInit, includeEnd)
		}
	case "Contained":
		foundIntervals = append(foundIntervals, receiver.interval)
		if receiver.left != nil {
			foundIntervals, _ = receiver.left.SearchMultipleIntervals(foundIntervals, value, includeInit, includeEnd)
		}
		if receiver.right != nil {
			foundIntervals, _ = receiver.right.SearchMultipleIntervals(foundIntervals, value, includeInit, includeEnd)
		}
	default:
		return foundIntervals, nil
	}
	return foundIntervals, nil
}

func (receiver *leaf[T]) SearchFirstInterval(value T, includeInit, includeEnd bool) (IntervalWithValue[T], error) {
	evaluation := receiver.interval.Evaluate(value, includeInit, includeEnd)
	switch evaluation {
	case "Under":
		if receiver.left != nil {
			return receiver.left.SearchFirstInterval(value, includeInit, includeEnd)
		}
	case "Upper":
		if receiver.right != nil {
			return receiver.right.SearchFirstInterval(value, includeInit, includeEnd)
		}
	case "Contained":
		return receiver.interval, nil
	default:
		return IntervalWithValue[T]{}, errorEvaluationValueOutOfBounds{value: value, evaluationResult: evaluation}
	}
	return IntervalWithValue[T]{}, errorEvaluationValueOutOfBounds{value: value, evaluationResult: evaluation}
}

func GenerateLeaf[T constraints.Ordered](sortedIntervals []IntervalWithValue[T], min, max int) *leaf[T] {
	if min == max {
		return nil
	}
	half := (min + max) / 2
	interval := sortedIntervals[half]
	left := GenerateLeaf(sortedIntervals, min, half)
	right := GenerateLeaf(sortedIntervals, half+1, max)
	return &leaf[T]{interval, left, right}
}
