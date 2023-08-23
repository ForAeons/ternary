# Package `ternary`

[![Go Reference](https://pkg.go.dev/badge/github.com/ForAeons/ternary.svg)](https://pkg.go.dev/github.com/ForAeons/ternary)

The `ternary` package provides utilities to write ternary operations in Go. It includes both strict and lazy evaluation of ternary expressions, allowing for more concise and functional programming.

## Functions

1. **`ternary.Op[T any](cond bool, trueVal, falseVal T) T`**:  
   Evaluates a ternary operation based on the condition `cond`. Returns `trueVal` if `cond` is `true`, otherwise returns `falseVal`. Type inference is utilized here.

   ```go
   result := ternary.Op[int](5 > 3, 10, 20) // result will be 10
   result = ternary.Op(5 > 3, 10, 20) // type inference will infer the type of result to be int
   ```

2. **`ternary.LazyOp[T any](cond bool, trueVal, falseVal func() T) T`**:  
   Similar to `Op`, but evaluates the `trueVal` and `falseVal` lazily by accepting them as functions. This avoids unnecessary computations. Can be useful for nil checks or expensive operations.

   ```go
   var err error
   result1 := ternary.LazyOp[string](err != nil, func() int { return "not nil" }, func() int { return "nil" }) // result will be "nil"
   result2 := ternary.LazyOp(5 > 3, func() int { return 10 }, func() int { return 20 }) // type inference will infer the type of result to be int
   ```

   You can also use it to avoid nil pointer dereferences:

   ```go
    type Foo struct {
        Bar string
    }

    var foo *Foo

    result := ternary.LazyOp[string](foo != nil, func() string { return foo.Bar }, func() string { return "nil" }) // result will be "nil"
   ```

## Types

1. **`ternary.TernaryIf[T any]`**:  
   Represents the conditional part of a ternary expression.

2. **`ternary.TernaryThen[T any]`**:  
   Represents the true value of a ternary expression.

3. **`ternary.LazyTernaryThen[T any]`**:  
   Represents the true value of a ternary expression with lazy evaluation.

## Methods

1. **`ternary.If[T any](cond bool) *TernaryIf[T]`**:  
   Begins a ternary expression by evaluating a condition.

   ```go
   ternary.If[int](5 > 3).Then(10).Else(20) // returns 10
   ```

2. **`Then(trueVal T) *TernaryThen[T]`**:  
   Specifies the true value of a ternary expression.

3. **`Else(falseVal T) T`**:  
   Specifies the false value of a ternary expression and completes it.

   ```go
   ternary.If[int](5 < 3).Then(10).Else(20) // returns 20
   ```

4. **`LThen(trueVal func() T) *LazyTernaryThen[T]`**:  
   Specifies the true value of a ternary expression with lazy evaluation.

5. **`LElse(falseVal func() T) T`**:  
   Specifies the false value of a ternary expression with lazy evaluation and completes it.

   ```go
   ternary.If[int](5 > 3).LThen(func() int { return 10 }).LElse(func() int { return 20 }) // returns 10
   ```

## Nesting Ternary Methods

For longer ternary statements, the ternary methods can be nested:

```go
var bar string = ternary.If[string](false).
  Then("foo").
  Else(ternary.If[string](true).
    Then("bar").
    Else("baz"),
  )
```

## Summary

The `ternary` package provides a convenient way to write and evaluate ternary expressions in Go. By supporting both strict and lazy evaluation and allowing for nesting, it enables developers to write more concise and expressive code. Type inference in the `Op` method further enhances readability.
