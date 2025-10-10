package main

import (
	"fmt"
	"os"
	"strconv"
)
func main(){
	if len(os.Args) <4 {
		fmt.Println("Usage: go run main.go <num1> <op> <num2>")
		return
	}
	a,_ := strconv.ParseFloat(os.Args[1],64)
	op := os.Args[2]
	b,_ := strconv.ParseFloat(os.Args[3],64)

	var result float64

	switch op {
	case "+":
		result = a+b
	case "-":
		result = a-b
	case "*":
		result = a*b
	case "/":
		if b == 0{
		fmt.Println("Error: division by zero")
		return
		}
		result = a+b
	default:
		fmt.Println("Invalid operator")	
		return
	}
	fmt.Println("Result: ",result)
}