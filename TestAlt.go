package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Введите выражение:")
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		panic("Ошибка при чтении ввода.")
	}

	input := scanner.Text()
	result := calculate(input)
	fmt.Println(result)
}

func calculate(input string) string {
	input = strings.TrimSpace(input)
	if input[0] != '"' {
		panic("Первая строка должна быть в кавычках.")
	}
	endQuoteIndex := strings.Index(input[1:], "\"") + 1
	if endQuoteIndex == 0 {
		panic("Неверный формат выражения. Ожидалось: 'строка операция строка/число'")
	}
	str1 := input[1:endQuoteIndex]
	remaining := input[endQuoteIndex+1:]
	remaining = strings.TrimSpace(remaining)
	if len(remaining) < 1 {
		panic("Неверный формат выражения. Ожидалось: 'строка операция строка/число'")
	}
	operator := string(remaining[0])
	remaining = remaining[1:]
	remaining = strings.TrimSpace(remaining)
	if remaining[0] != '"' {
		if _, err := strconv.Atoi(remaining); err != nil {
			panic("Вторая часть должна быть строкой в кавычках или числом.")
		}
		return formatResult(handleOperation(str1, operator, remaining))
	} else {
		endQuoteIndex = strings.Index(remaining[1:], "\"") + 1
		if endQuoteIndex == 0 {
			panic("Неверный формат выражения. Ожидалось: 'строка операция строка/число'")
		}
		str2 := remaining[1:endQuoteIndex]
		return formatResult(handleOperation(str1, operator, str2))
	}
}

func handleOperation(str1, operator, str2 string) string {
	switch operator {
	case "+":
		return handleAddition(str1, str2)
	case "-":
		return handleSubtraction(str1, str2)
	case "*":
		return handleMultiplication(str1, str2)
	case "/":
		return handleDivision(str1, str2)
	default:
		panic("Неподдерживаемая операция.")
	}
}

func handleAddition(str1, str2 string) string {
	result := str1 + str2
	return truncate(result)
}

func handleSubtraction(str1, str2 string) string {
	result := strings.ReplaceAll(str1, str2, "")
	return truncate(result)
}

func handleMultiplication(str1, str2 string) string {
	n, _ := strconv.Atoi(str2)
	if n < 1 || n > 10 {
		panic("Число должно быть от 1 до 10.")
	}
	result := strings.Repeat(str1, n)
	return truncate(result)
}

func handleDivision(str1, str2 string) string {
	n, _ := strconv.Atoi(str2)
	if n < 1 || n > 10 {
		panic("Число должно быть от 1 до 10.")
	}
	if len(str1) < n {
		return ""
	}
	result := str1[:len(str1)/n]
	return truncate(result)
}

func truncate(s string) string {
	if len(s) > 40 {
		return s[:40] + "..."
	}
	return s
}

func formatResult(s string) string {
	return fmt.Sprintf("\"%s\"", s)
}
