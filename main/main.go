package main

import (
	"fmt"
	"math"
	"strconv"
)

var history = []string{}

func main() {
	for {
		var options string
		fmt.Printf("Check history press c, otherwise press any button")
		fmt.Scanln(&options)
		if options == "c" {
			for i := range history {
				fmt.Println(history[i])
			}
			continue
		}
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
	//SR = square root
	operators := []string{"+", "-", "*", "/", "**", "sr"}

	for !isValid {
		fmt.Printf("Enter the operator(%v): \n", operators)
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
	var output string
	switch operator {
	case "+":
		output = fmt.Sprintf("%v + %v = %v\n", n1, n2, n1+n2)
		history = append(history, output)
		fmt.Println(output)
	case "-":
		output = fmt.Sprintf("%v - %v = %v\n", n1, n2, n1-n2)
		history = append(history, output)
		fmt.Println(output)
	case "*":
		output = fmt.Sprintf("%v * %v = %v\n", n1, n2, n1*n2)
		history = append(history, output)
		fmt.Println(output)
	case "/":
		if n2 == 0 {
			fmt.Println("The second number is invalid")
			return
		}
		output = fmt.Sprintf("%v / %v = %v\n", n1, n2, n1/n2)
		history = append(history, output)
		fmt.Println(output)
	case "**":
		output = fmt.Sprintf("%v ** %v = %v\n", n1, n2, math.Pow(n1, n2))
		history = append(history, output)
		fmt.Println(output)
	case "sr":
		output = fmt.Sprintf("The square root of %v = %v. The square root of %v = %v\n", n1, math.Sqrt(n1), n2, math.Sqrt(n2))
		history = append(history, output)
		fmt.Println(output)
	}
}
