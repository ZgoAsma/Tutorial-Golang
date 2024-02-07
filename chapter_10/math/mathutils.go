package math

func Multiply(x, y int) int {
	return x * y
}

func Add(x, y int) int {
	return x + y
}

func SelfMath(mathFunc func(int, int) int) func(int) int {
	return func(x int) int {
		return mathFunc(x, x)
	}
}
func Mul(x, y int) int {
	return x * y
}

// aggregate applies the given math function to the first 3 inputs
func Aggregate(a, b, c int, arithmetic func(int, int) int) int {
	return arithmetic(arithmetic(a, b), c)
}
