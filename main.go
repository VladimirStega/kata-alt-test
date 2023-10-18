package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Введите выражение")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		expression := strings.Split(text, " ")
		if len(expression) != 3 {
			fmt.Println("Невырный формат выражения")
			continue
		}

		a := expression[0]
		operator := expression[1]
		b := expression[2]

		result, err := calculate(a, operator, b)
		if err != nil {
			fmt.Println("Ошибка:", err)
		}

		if a[0] != '"' || a[len(a)-1] != '"' {
			fmt.Println("Первый операнд должен быть в кавычках")
			continue
		}

		if operator == "-" && (b[0] != '"' || b[len(b)-1] != '"') {
			fmt.Println("Второй операнд должен быть в кавычках")
			continue
		}

		if len(a) > 12 || len(b) > 12 {
			fmt.Println("Операнды должены быть не больше 10 символов")
			continue
		}

		if len(result) > 40 {
			result = "\"" + strings.ReplaceAll(result[:40], "\"", "") + "..." + "\""
		}

		if err == nil {
			fmt.Println("Результат:", "\""+strings.ReplaceAll(result, "\"", "")+"\"")
		}
	}

}

func calculate(a, operator, b string) (string, error) {
	a = strings.ReplaceAll(a, "\"", "")
	b = strings.ReplaceAll(b, "\"", "")
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		if strings.Contains(a, b) {
			return strings.Replace(a, b, "", -1), nil
		}
		return a, nil

	case "*":
		num, err := strconv.Atoi(b)
		if err != nil || num == 0 {
			return "", fmt.Errorf("Умножение на ноль или неверный формат числа")
		}
		if err != nil || num > 10 {
			return "", fmt.Errorf("Нельзя умножать на число больше 10")
		}
		return strings.Repeat(a, num), nil

	case "/":
		num, err := strconv.Atoi(b)
		if err != nil || num == 0 {
			return "", fmt.Errorf("Деление на ноль или неверный формат числа")
		}
		if err != nil || len(a) < num {
			return "", fmt.Errorf("Деление невозможно, строка слишком короткая")
		}
		if err != nil || len(a)%2 != 0 {
			return a[:len(a)/num+1], nil
		}
		if err != nil || num > 10 {
			return "", fmt.Errorf("Нельзя делить на число больше 10")
		}
		if err != nil || len(a) == num {
			return a[:1], nil
		}

		return a[:len(a)/num], nil

	default:
		return "", fmt.Errorf("Неверная операция")
	}
}
