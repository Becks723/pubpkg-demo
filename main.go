package pubpkg

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Printf("Enter number1: ")
	fmt.Scan(&x)
	fmt.Printf("Enter number2: ")
	fmt.Scan(&y)

	var op rune
	fmt.Printf("Enter an operation (+-*/): ")
	fmt.Scanf("%c", &op)

	var result int
	switch op {
	case '+':
		result = Add(x, y)
	case '-':
		result = Minus(x, y)
	case '*':
		result = Multiply(x, y)
	case '/':
		result = Divide(x, y)
	default:
		fmt.Printf("Unknown operation: %c", op)
	}

	fmt.Printf("Result is %d!", result)
	fmt.Println()
}

func Add(x, y int) int {
	return x + y
}

func Minus(x, y int) int {
	return x - y
}

func Multiply(x, y int) int {
	return x * y
}

func Divide(x, y int) int {
	return x / y
}
