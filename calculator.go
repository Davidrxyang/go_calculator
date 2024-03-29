package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Addition function
func add(x, y float64) float64 {
	return x + y
}

// Subtraction function
func subtract(x, y float64) float64 {
	return x - y
}

// Multiplication function
func multiply(x, y float64) float64 {
	return x * y
}

// Division function
func divide(x, y float64) float64 {
	if y == 0 {
		fmt.Println("Error: Division by zero")
		return 0
	}
	return x / y
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Simple Calculator")
	fmt.Println("-----------------")
	fmt.Println("Enter the operation (+, -, *, /):")

	operation, _ := reader.ReadString('\n')
	operation = strings.TrimSpace(operation)

	fmt.Println("Enter the first number:")
	firstInput, _ := reader.ReadString('\n')
	firstInput = strings.TrimSpace(firstInput)
	firstNumber, err := strconv.ParseFloat(firstInput, 64)
	if err != nil {
		fmt.Println("Error reading the first number:", err)
		return
	}

	fmt.Println("Enter the second number:")
	secondInput, _ := reader.ReadString('\n')
	secondInput = strings.TrimSpace(secondInput)
	secondNumber, err := strconv.ParseFloat(secondInput, 64)
	if err != nil {
		fmt.Println("Error reading the second number:", err)
		return
	}

	switch operation {
	case "+":
		fmt.Println("Result:", add(firstNumber, secondNumber))
	case "-":
		fmt.Println("Result:", subtract(firstNumber, secondNumber))
	case "*":
		fmt.Println("Result:", multiply(firstNumber, secondNumber))
	case "/":
		fmt.Println("Result:", divide(firstNumber, secondNumber))
	default:
		fmt.Println("Invalid operation")
	}
}
