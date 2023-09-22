package main

import (
	"fmt"
	"strconv"
)

func main() {
	for {
		var number1, number2 float64
		var operator string
		//first number
		number1 = inputNumber("first")
		//second number
		number2 = inputNumber("second")

		//operator
		operator = inputOperator()

		operate(number1, number2, operator)
	}
}

func inputNumber(idx string) float64 {
	var input string
	fmt.Printf("Enter the %v number: \n", idx)
	fmt.Scanln(&input)

	f, err := strconv.ParseFloat(input, 64)
	for err != nil {
		fmt.Println("Invalid number, please retry again: ")
		fmt.Scanln(&input)
		f, err = strconv.ParseFloat(input, 64)
	}
	return f
}

func inputOperator() string {
	var operator string
	isValid := false
	operators := []string{"+", "-", "*", "/"}

	for !isValid {
		fmt.Println("Enter the operator: ")
		fmt.Scanln(&operator)

		for _, v := range operators {
			if operator == v {
				isValid = true
			}
		}
	}
	return operator
}

func operate(n1, n2 float64, operator string) {
	switch operator {
	case "+":
		fmt.Printf("%v + %v = %v\n", n1, n2, n1+n2)
	case "-":
		fmt.Printf("%v - %v = %v\n", n1, n2, n1-n2)
	case "*":
		fmt.Printf("%v * %v = %v\n", n1, n2, n1*n2)
	case "/":
		if n2 == 0 {
			fmt.Println("The second number is invalid")
		}
		fmt.Printf("%v / %v = %v\n", n1, n2, n1/n2)
	}
}
