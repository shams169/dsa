package main

import "fmt"

func BalancingSymbols(input string) bool {
	s := Stack[rune]{}
	// symbols := []rune(input)
	for _, symbol := range input {
		switch symbol {
		case '(', '[', '{':
			s.Push(symbol)
		case ')':
			if value, ok := s.Pop(); !ok || value != '(' {
				fmt.Println(string(symbol))
				return false
			}
		case ']':
			if value, ok := s.Pop(); !ok || value != '[' {
				fmt.Println(string(value))
				return false
			}
		case '}':
			if value, ok := s.Pop(); !ok || value != '{' {
				fmt.Println(string(value))
				return false
			}
		}
	}
	return s.IsEmpty()
}
