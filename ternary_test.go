package ternary

import (
	"testing"
)

func TestOp(t *testing.T) {
	result := Op[int](true, 10, 20)
	if result != 10 {
		t.Errorf("Expected 10, got %d", result)
	}

	result = Op[int](false, 10, 20)
	if result != 20 {
		t.Errorf("Expected 20, got %d", result)
	}
}

func TestLazyOp(t *testing.T) {
	calledTrue := false
	calledFalse := false
	trueVal := func() int { calledTrue = true; return 10 }
	falseVal := func() int { calledFalse = true; return 20 }

	result := LazyOp[int](true, trueVal, falseVal)
	if result != 10 || !calledTrue || calledFalse {
		t.Errorf("Expected 10, got %d; trueVal called: %v, falseVal called: %v", result, calledTrue, calledFalse)
	}

	calledTrue = false
	calledFalse = false
	result = LazyOp[int](false, trueVal, falseVal)
	if result != 20 || calledTrue || !calledFalse {
		t.Errorf("Expected 20, got %d; trueVal called: %v, falseVal called: %v", result, calledTrue, calledFalse)
	}
}

func TestTernaryIfThenElse(t *testing.T) {
	result := If[int](true).Then(10).Else(20)
	if result != 10 {
		t.Errorf("Expected 10, got %d", result)
	}

	result = If[int](false).Then(10).Else(20)
	if result != 20 {
		t.Errorf("Expected 20, got %d", result)
	}
}

func TestLazyTernaryIfLThenLElse(t *testing.T) {
	calledTrue := false
	calledFalse := false
	trueVal := func() int { calledTrue = true; return 10 }
	falseVal := func() int { calledFalse = true; return 20 }

	result := If[int](true).LThen(trueVal).LElse(falseVal)
	if result != 10 || !calledTrue || calledFalse {
		t.Errorf("Expected 10, got %d; trueVal called: %v, falseVal called: %v", result, calledTrue, calledFalse)
	}

	calledTrue = false
	calledFalse = false
	result = If[int](false).LThen(trueVal).LElse(falseVal)
	if result != 20 || calledTrue || !calledFalse {
		t.Errorf("Expected 20, got %d; trueVal called: %v, falseVal called: %v", result, calledTrue, calledFalse)
	}
}

func TestNestedTernary(t *testing.T) {
	result := If[string](false).Then("foo").Else(If[string](true).Then("bar").Else("baz"))
	if result != "bar" {
		t.Errorf("Expected 'bar', got '%s'", result)
	}
}
