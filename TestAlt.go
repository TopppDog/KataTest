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

	parts := strings.Fields(input)
	if len(parts) != 3 {
		panic("Неверный формат выражения. Ожидалось: 'строка операция строка/число'")
	}
	str1 := parts[0]
	operator := parts[1]
	str2 := parts[2]
	if str1[0] == '"' && str1[len(str1)-1] == '"' {
		str1 = str1[1 : len(str1)-1]
	} else {
		panic("Первая строка должна быть в кавычках.")
	}
	if str2[0] == '"' && str2[len(str2)-1] == '"' {
		str2 = str2[1 : len(str2)-1]
	} else {
		if _, err := strconv.Atoi(str2); err != nil {
			panic("Вторая часть должна быть строкой в кавычках или числом.")
		}
	}

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
