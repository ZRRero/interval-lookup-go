package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type errorInvalidInterval[T constraints.Ordered] struct {
	init T
	end  T
}

func (e errorInvalidInterval[T]) Error() string {
	return fmt.Sprintf("The interval init %v is greater than the interval end %v", e.init, e.end)
}

type errorEvaluationValueOutOfBounds struct {
	value            any
	evaluationResult string
}

func (e errorEvaluationValueOutOfBounds) Error() string {
	return fmt.Sprintf("The value %v generated the evaluation result %v, which is not recognized", e.value, e.evaluationResult)
}

type errorRootIsNil struct{}

func (e errorRootIsNil) Error() string {
	return fmt.Sprintf("The value to search is nil")
}

type errorIntervalsOverlapped struct{}

func (e errorIntervalsOverlapped) Error() string {
	return fmt.Sprintf("There are overlapped intervals")
}
