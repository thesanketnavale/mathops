package mathops

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) (result int, divisionByZero bool) {
	if b == 0 {
		return 0, true
	}
	return a / b, false
}
