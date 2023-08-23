package ternary

// Op evaluates a ternary operation based on the condition `cond`.
//
// Returns `trueVal` if `cond` is true, otherwise returns `falseVal`.
//
// Example: result := Op[int](5 > 3, 10, 20) // result will be 10
func Op[T any](cond bool, trueVal T, falseVal T) T {
	if cond {
		return trueVal
	}
	return falseVal
}

// LazyOp is similar to Op, but evaluates the `trueVal` and `falseVal`
//
// lazily by accepting them as functions. This avoids unnecessary computations.
//
// Example: result := LazyOp[int](5 > 3, func() int { return 10 }, func() int { return 20 }) // result will be 10
func LazyOp[T any](cond bool, trueVal, falseVal func() T) T {
	if cond {
		return trueVal()
	}
	return falseVal()
}

// TernaryIf represents the conditional part of a ternary expression.
type TernaryIf[T any] struct {
	condition bool
}

// TernaryThen represents the true value of a ternary expression.
type TernaryThen[T any] struct {
	condition bool
	trueVal   T
}

// LazyTernaryThen represents the true value of a ternary expression with lazy evaluation.
type LazyTernaryThen[T any] struct {
	condition bool
	trueVal   func() T
}

// If begins a ternary expression by evaluating a condition.
//
// Example: If[int](5 > 3).Then(10).Else(20) // returns 10
func If[T any](cond bool) *TernaryIf[T] {
	return &TernaryIf[T]{condition: cond}
}

// Then specifies the true value of a ternary expression.
func (i *TernaryIf[T]) Then(trueVal T) *TernaryThen[T] {
	return &TernaryThen[T]{condition: i.condition, trueVal: trueVal}
}

// Else specifies the false value of a ternary expression and completes it.
//
// Example: If[int](5 < 3).Then(10).Else(20) // returns 20
func (t *TernaryThen[T]) Else(falseVal T) T {
	return Op[T](t.condition, t.trueVal, falseVal)
}

// LThen specifies the true value of a ternary expression with lazy evaluation.
func (i *TernaryIf[T]) LThen(trueVal func() T) *LazyTernaryThen[T] {
	return &LazyTernaryThen[T]{condition: i.condition, trueVal: trueVal}
}

// LElse specifies the false value of a ternary expression with lazy evaluation and completes it.
//
// Example: If[int](5 > 3).LThen(func() int { return 10 }).LElse(func() int { return 20 }) // returns 10
func (t *LazyTernaryThen[T]) LElse(falseVal func() T) T {
	return LazyOp[T](t.condition, t.trueVal, falseVal)
}
