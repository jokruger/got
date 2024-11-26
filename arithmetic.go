package got

// Add returns the sum of two numbers.
func Add[T Numeric](a, b T) T {
	return a + b
}

// Sub returns the difference of two numbers.
func Sub[T Numeric](a, b T) T {
	return a - b
}

// Mul returns the product of two numbers.
func Mul[T Numeric](a, b T) T {
	return a * b
}

// Div returns the quotient of two numbers.
func Div[T Numeric](a, b T) T {
	return a / b
}
