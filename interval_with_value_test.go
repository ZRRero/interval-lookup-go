package main

import (
	"fmt"
	"testing"
)

func TestGenerateIntervalWithValueValidationOk(t *testing.T) {
	init := 0
	end := 1
	value := "a"
	interval, err := GenerateIntervalWithValue(init, end, value)
	if err != nil {
		t.Fatalf("Failed with error: %v", err)
	}
	if interval.init != init || interval.end != end || interval.value != value {
		t.Fatalf("Interval (%v, %v, %v) is different from expected (%v, %v, %v", interval.init, interval.end, interval.value, init, end, value)
	}
}

func TestGenerateIntervalWithValueValidationNoOk(t *testing.T) {
	init := 0
	end := -1
	value := "a"
	_, err := GenerateIntervalWithValue(init, end, value)
	expected := fmt.Sprintf("The interval init %v is greater than the interval end %v", init, end)
	if err == nil || err.Error() != expected {
		t.Fatalf("Failed with error: %v", err)
	}
}
