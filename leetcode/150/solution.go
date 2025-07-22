package _150

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	var stack []int

	for _, token := range tokens {
		if val, err := strconv.Atoi(token); err != nil {
			elements := stack[len(stack)-2:]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, elements[0]+elements[1])
			case "-":
				stack = append(stack, elements[0]-elements[1])
			case "*":
				stack = append(stack, elements[0]*elements[1])
			case "/":
				stack = append(stack, elements[0]/elements[1])
			}
		} else {
			stack = append(stack, val)
		}
	}

	return stack[0]
}
