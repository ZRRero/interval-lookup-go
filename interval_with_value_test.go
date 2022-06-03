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

func TestIntervalWithValue_IsOverlapped(t *testing.T) {
	intervalOne, _ := GenerateIntervalWithValue(0, 10, 0)
	intervalTwo, _ := GenerateIntervalWithValue(5, 15, 1)
	result := intervalOne.IsOverlapped(intervalTwo)
	if result != true {
		t.Fatalf("Intervals should be overlapped")
	}
}

func TestIntervalWithValue_IsNotOverlapped(t *testing.T) {
	intervalOne, _ := GenerateIntervalWithValue(0, 10, 0)
	intervalTwo, _ := GenerateIntervalWithValue(11, 20, 1)
	result := intervalOne.IsOverlapped(intervalTwo)
	if result != false {
		t.Fatalf("Intervals should not be overlapped")
	}
}

func TestIntervalWithValue_CompareEquals(t *testing.T) {
	intervalOne, _ := GenerateIntervalWithValue(0, 10, 0)
	intervalTwo, _ := GenerateIntervalWithValue(0, 10, 1)
	result := intervalOne.Compare(intervalTwo)
	if result != 0 {
		t.Fatalf("Compare result should be 0 and was %v", result)
	}
}

func TestIntervalWithValue_CompareGreater(t *testing.T) {
	intervalOne, _ := GenerateIntervalWithValue(0, 10, 0)
	intervalTwo, _ := GenerateIntervalWithValue(1, 10, 1)
	result := intervalOne.Compare(intervalTwo)
	if result != 1 {
		t.Fatalf("Compare result should be 1 and was %v", result)
	}
}

func TestIntervalWithValue_CompareLesser(t *testing.T) {
	intervalOne, _ := GenerateIntervalWithValue(0, 10, 0)
	intervalTwo, _ := GenerateIntervalWithValue(0, 9, 1)
	result := intervalOne.Compare(intervalTwo)
	if result != -1 {
		t.Fatalf("Compare result should be -1 and was %v", result)
	}
}
