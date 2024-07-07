package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Calculate")
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	oper := scanner.Text()
	components := strings.Split(oper, " ")
	if len(components) != 3 {
		panic("Колилество операндов не удволетворяет заданию")
		return
	}
	var op1 string = components[0]
	var op2 string = components[2]
	var n string = components[1]

	operand1, err1 := strconv.Atoi(op1)
	operand2, err := strconv.Atoi(op2)
	if err1 == nil && err == nil {
		if operand1 == 0 || operand2 == 0 || operand1 > 10 || operand2 > 10 {
			panic("Не входит в диапазон от 1 до 10")
		}
		result := calculate(operand1, operand2, n)
		fmt.Println(result)
		os.Exit(0)
	}

	if err1 == nil && err == nil {
		result := calculate(operand1, operand2, n)
		fmt.Println(result)
		os.Exit(0)
	}
	if err1 != nil && err != nil || err1 == nil && err == nil {
		operand1 = convertRomanToArabic(op1, n)
		operand2 = convertRomanToArabic(op2, n)
		if operand1 == 0 || operand2 == 0 {
			panic("Символы не входят в римскую систему")

		}
		result := calculate(operand1, operand2, n)
		rim := convert(result)
		fmt.Println(rim)
	} else {
		panic("Используются одновременно разные системы счисления")

	}

}

func convertRomanToArabic(comp, n string) int {
	if n == "-" {
		panic("В римской системе нет отрицательных чисел")

	}
	if comp == "I" {
		return 1
	}
	if comp == "II" {
		return 2
	}
	if comp == "III" {
		return 3
	}
	if comp == "IV" {
		return 4
	}
	if comp == "V" {
		return 5
	}
	if comp == "VI" {
		return 6
	}
	if comp == "VII" {
		return 7
	}
	if comp == "VIII" {
		return 8
	}
	if comp == "IX" {
		return 9
	}
	if comp == "X" {
		return 10
	}
	if comp == "L" {
		return 50
	}
	if comp == "C" {
		return 100
	}
	return 0
}

func calculate(operand1 int, operand2 int, n string) int {
	switch n {
	case "+":
		return operand1 + operand2
	case "-":
		return operand1 - operand2
	case "*":
		return operand1 * operand2
	case "/":
		return operand1 / operand2
	default:
		return 0
	}
}
func convert(result int) string {
	var rim strings.Builder
	a := result
	var romanNumbers []string = []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "L", "C"}

	if a > 0 {

		if a/100 >= 1 {
			rim.WriteString(romanNumbers[11])
			a -= 100
		}
		if a/50 >= 1 {
			rim.WriteString(romanNumbers[10])
			a -= 50
		}
		if a/10 >= 1 {
			count := a / 10
			for i := 0; i < count; i++ {
				rim.WriteString(romanNumbers[9])
				a -= 10
			}
		}
		if a > 0 {
			rim.WriteString(romanNumbers[a-1])
			a -= a
		}

	}
	return rim.String()

}
