package basicutil

import (
	"github.com/jokruger/got"
)

// Add returns the sum of two numbers.
func Add[T got.Numeric](a, b T) T {
	return a + b
}

// Sub returns the difference of two numbers.
func Sub[T got.Numeric](a, b T) T {
	return a - b
}

// Mul returns the product of two numbers.
func Mul[T got.Numeric](a, b T) T {
	return a * b
}

// Div returns the quotient of two numbers.
func Div[T got.Numeric](a, b T) T {
	return a / b
}
