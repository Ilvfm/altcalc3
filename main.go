package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Привет,давай смотреть,что я умею")
		return
	}

	expression := os.Args[1]
	result := evaluateExpression(expression)

	fmt.Println(result)
}

func evaluateExpression(expression string) string {
	parts := strings.Split(expression, " ")
	if len(parts) != 3 {
		panic("Ошибка")
	}

	left := strings.Trim(parts[0], "\"")
	operator := parts[1]
	right := strings.Trim(parts[2], "\"")
	fmt.Scan(&left, &operator, &right)

	switch operator {
	case "+":
		return left + right
	case "-":
		return strings.ReplaceAll(left, right, "")
	case "*":
		n := parseInt(right)
		return strings.Repeat(left, n)
	case "/":
		n := parseInt(right)
		return left[:len(left)/n]
	default:
		panic("Ошибка")
	}
}

func parseInt(s string) int {
	var result int
	var input string
	fmt.Scanf("%v", &input)
	_, err := fmt.Sscanf(s, "%d", &result)
	if err != nil {
		panic("Ошибка")
	}
	return result
}
