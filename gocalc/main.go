package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) == 2 && os.Args[1] == "history" {
		data, err := os.ReadFile("history.txt")
		if err != nil {
			fmt.Println("No history found")
			return
		}
		fmt.Println(string(data))
		return
	}

	if len(os.Args) < 4 {
		fmt.Println("Usage: go run main.go <num1> <op> <num2>")
		return
	}
	a, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Invalid number", os.Args[1])
		return
	}
	op := os.Args[2]
	b, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Println("Invalid number", os.Args[3])
		return
	}
	var result float64

	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			fmt.Println("Error: division by zero")
			return
		}
		result = a / b
	default:
		fmt.Println("Invalid operator")
		return
	}
	fmt.Println("Result: ", result)

	f, err := os.OpenFile("history.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Could not write history: ", err)
		return
	}
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("%.2f %s %.2f = %.2f\n", a, op, b, result))
	if err != nil {
		fmt.Printf("Could not write history: %v\n", err)
		return
	}

}
